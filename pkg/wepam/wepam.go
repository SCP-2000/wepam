package main

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_modules.h>
#include <stdlib.h>
typedef const char ** cchar;
char *get_user(pam_handle_t *pamh);
char *get_item(pam_handle_t *pamh, int item_type);
int prompt(pam_handle_t *pamh, int style, const char *fmt);
*/
import "C"
import (
	"fmt"
	"github.com/SCP-2000/wepam/pkg/oauth2"
	"runtime"
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

func Prompt(pamh *C.pam_handle_t, style C.int, msg string) C.int {
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	return C.prompt(pamh, style, cs)
}

/* Authentication API's */
//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv C.cchar) C.int {
	runtime.LockOSThread()
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

	challenges := make(chan *oauth2.Challenge)
	errors := make(chan error)
	go func() {
		errors <- Auth(GoStringSlice(argc, argv), items, challenges)
	}()

	for challenge := range challenges {
		Prompt(pamh, C.PAM_TEXT_INFO, fmt.Sprintf("please visit %s and input %s",
			challenge.DeviceAuth.VerificationURI,
			challenge.DeviceAuth.UserCode))
	}

	err := <-errors
	if err != nil {
		Prompt(pamh, C.PAM_ERROR_MSG, fmt.Sprintf("authentication failure: %s", err))
		return C.PAM_AUTH_ERR
	}
	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv C.cchar) C.int {
	return C.PAM_IGNORE
}

func main() {}
