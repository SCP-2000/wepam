module github.com/SCP-2000/wepam

go 1.16

require (
	github.com/google/go-github/v33 v33.0.0
	github.com/open-policy-agent/opa v0.26.0
	github.com/spf13/pflag v1.0.1
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93
	google.golang.org/appengine v1.6.7 // indirect
)

replace golang.org/x/oauth2 => github.com/SCP-2000/oauth2 v0.0.0-20210307144543-857e62bb3dae
