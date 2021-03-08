package policy

default allow = false

user := oauth2("github", "111109bc0316c05e23aa")

allow {
  user.login = "NickCao"
}
