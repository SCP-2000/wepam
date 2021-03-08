module github.com/SCP-2000/wepam

go 1.16

require (
	github.com/google/go-github/v33 v33.0.0
	github.com/open-policy-agent/opa v0.26.0
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93
)

replace golang.org/x/oauth2 => github.com/SCP-2000/oauth2 v0.0.0-20210308124225-a3914ffd5793
