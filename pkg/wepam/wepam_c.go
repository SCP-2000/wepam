package main

/*
#include <security/pam_modules.h>
#include <security/pam_ext.h>
#include <string.h>

char *get_user(pam_handle_t *pamh) {
  const char *user;
  if (pam_get_user(pamh, &user, NULL) != PAM_SUCCESS || user == NULL)
        return NULL;
  return strdup(user);
}

char *get_item(pam_handle_t *pamh, int item_type) {
  const void *item;
  if (pam_get_item(pamh, item_type, &item) != PAM_SUCCESS || item == NULL)
    return NULL;
  return strdup((const char *) item);
}

int prompt(pam_handle_t *pamh, int style, const char *fmt) {
  return pam_prompt(pamh, style, NULL, "%s", fmt);
}
*/
import "C"
