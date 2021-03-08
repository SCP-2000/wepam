package main

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_modules.h>
typedef const char ** cchar;
char *get_user(pam_handle_t *pamh);
char *get_item(pam_handle_t *pamh, int item_type);
int prompt(pam_handle_t *pamh, const char *fmt);
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

func GoStringSlice(argc C.int, argv **C.char) []string {
	length := int(argc)
	orig := (*[1 << 30]*C.char)(unsafe.Pointer(argv))
	strings := make([]string, length)
	for i := 0; i < length; i++ {
		strings[i] = C.GoString(orig[i])
	}
	return strings
}

/* Authentication API's */
//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv C.cchar) C.int {
	user := C.get_user(pamh)
	if user == nil {
		return C.PAM_USER_UNKNOWN
	}

	items_t := map[string]C.int{
		"service": C.PAM_SERVICE,
		"user":    C.PAM_USER,
		"ruser":   C.PAM_RUSER,
		"rhost":   C.PAM_RHOST,
	}
	items := make(map[string]string)
	for k, v := range items_t {
		item := C.get_item(pamh, v)
		if item == nil {
			items[k] = ""
		} else {
			items[k] = C.GoString(item)
		}
	}
	err := Auth(GoStringSlice(argc, argv), items, func(s string) error {
		if C.prompt(pamh, C.CString(s)) != C.PAM_SUCCESS {
			return fmt.Errorf("failed to prompt user")
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return C.PAM_AUTH_ERR
	}
	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv C.cchar) C.int {
	return C.PAM_IGNORE
}

func main() {}
