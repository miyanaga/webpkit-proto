package dwebp

/*
#cgo CFLAGS: -I. -I../libwebp -I../libwebp/src -I../libwebp/examples -I/usr/local/include
#cgo LDFLAGS: -static -L../libwebp/build -L/usr/local/lib -lexampleutil -lextras -limagedec -limageenc -limageioutil -lwebp -lwebpdecoder -lwebpdemux -lwebpmux -lsharpyuv -lpthread -ljpeg -ltiff -lpng -lz -lm

#include <stdlib.h>
#include "dwebp.c"

int dwebp(int argc, const char* argv[]) {
	return internal_dwebp(argc, argv);
}
*/
import "C"

import (
	"unsafe"
)

func DWebP(args ...string) int {
	args = append([]string{"dwebp"}, args...)
	argc := C.int(len(args))

	cArgs := make([]*C.char, len(args))
	for i, arg := range args {
		cArgs[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgs[i]))
	}

	argv := (**C.char)(unsafe.Pointer(&cArgs[0]))
	code := int(C.dwebp(argc, argv))

	return code
}
