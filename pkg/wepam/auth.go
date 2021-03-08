package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/SCP-2000/wepam/pkg/adapter"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
	"time"
)

func Auth(args []string, pam_items map[string]string, callback func(string) error) error {
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	p := flags.String("p", "/etc/wepam/policy.rego", "path to policy file")
	err := flags.Parse(args)
	if err != nil {
		return err
	}

	r := rego.New(
		rego.Load([]string{*p}, nil),
		rego.Input(map[string]interface{}{
			"pam": pam_items,
		}),
		// oauth2 device authorization flow function, usage: oauth2("<provider>", "<client_id>")
		rego.Function2(
			&rego.Function{
				Name:    "oauth2",
				Decl:    types.NewFunction(types.Args(types.S, types.S), types.A),
				Memoize: true,
			}, func(ctx rego.BuiltinContext, p1 *ast.Term, p2 *ast.Term) (*ast.Term, error) {
				var provider, client_id string
				if err := ast.As(p1.Value, &provider); err != nil {
					return nil, err
				}
				if err := ast.As(p2.Value, &client_id); err != nil {
					return nil, err
				}

				config, err := adapter.New(provider, client_id)
				if err != nil {
					return nil, err
				}

				// TODO: better UX
				data, err := config.Auth(ctx.Context, func(url string, code string) error {
					return callback(fmt.Sprintf("please go to %s and input %s", url, code))
				})
				if err != nil {
					return nil, err
				}

				value, err := ast.ValueFromReader(bytes.NewBuffer(data))
				if err != nil {
					return nil, err
				}
				return ast.NewTerm(value), nil
			}),
		// query for authentication decision
		rego.Query("data.policy.allow"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute) // TODO: user configurable timeout
	defer cancel()

	res, err := r.Eval(ctx)
	if err != nil {
		return err
	}

	if len(res) != 1 || len(res[0].Expressions) != 1 {
		return fmt.Errorf("malformed policy evaluation result")
	}

	allow, ok := res[0].Expressions[0].Value.(bool)
	if !ok {
		return fmt.Errorf("malformed policy evaluation result")
	}

	if allow {
		return nil
	}
	return fmt.Errorf("authencation failed")
}
