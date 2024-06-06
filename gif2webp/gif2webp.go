package gif2webp

/*
#cgo CFLAGS: -I. -I../libwebp -I../libwebp/src -I../libwebp/examples
#cgo LDFLAGS: -static libwebp/build/CMakeFiles/gif2webp.dir/examples/gifdec.c.o -L../libwebp/build -lexampleutil -lextras -limagedec -limageenc -limageioutil -lwebpdecoder -lwebpdemux -lwebpmux -lpthread -lwebp -lsharpyuv -lgif -ljpeg -ltiff -lpng -lz -lm

#define WEBP_HAVE_GIF 1

#include <stdlib.h>
#include "gif2webp.c"

int gif2webp(int argc, const char* argv[]) {
	return internal_gif2webp(argc, argv);
}
*/
import "C"

import (
	"unsafe"
)

func Gif2WebP(args ...string) int {
	args = append([]string{"gif2webp"}, args...)
	argc := C.int(len(args))

	cArgs := make([]*C.char, len(args))
	for i, arg := range args {
		cArgs[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgs[i]))
	}

	argv := (**C.char)(unsafe.Pointer(&cArgs[0]))
	code := int(C.gif2webp(argc, argv))

	return code
}
