package policy

default allow = false

user := oauth2("github", "2d8d2c5e9878098d0657")

allow {
  user.login = "NickCao"
}
