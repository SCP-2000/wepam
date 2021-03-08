package policy

default allow = false

user := oauth2("github", "111109bc0316c05e23aa")
user2 := oauth2("github", "584467084a6d0612ce0d")

allow {
  user.login = "NickCao"
  user2.login = "NickCao"
}
