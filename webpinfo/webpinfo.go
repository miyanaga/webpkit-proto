package webpinfo

/*
#cgo CFLAGS: -I. -I../libwebp -I../libwebp/src -I../libwebp/examples
#cgo LDFLAGS: -static -L../libwebp/build -lexampleutil -lextras -limagedec -limageenc -limageioutil -lwebp -lwebpdecoder -lwebpdemux -lwebpmux -lsharpyuv -lpthread -ljpeg -ltiff -lpng -lz -lm

#include <stdlib.h>
#include "webpinfo.c"

int webpinfo(int argc, const char* argv[]) {
	return internal_webpinfo(argc, argv);
}
*/
import "C"
import "unsafe"

func WebPInfo(args ...string) int {
	args = append([]string{"webpinfo"}, args...)
	argc := C.int(len(args))

	cArgs := make([]*C.char, len(args))
	for i, arg := range args {
		cArgs[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgs[i]))
	}

	argv := (**C.char)(unsafe.Pointer(&cArgs[0]))
	code := int(C.webpinfo(argc, argv))

	return code
}
