package imagetype

import (
	"bytes"
	"os"

	"github.com/ideamans/webpkit/l10n"
	"github.com/phuslu/fastimage"
)

var (
	jpegMagic  = []byte{0xFF, 0xD8}
	pngMagic   = []byte{0x89, 0x50, 0x4E, 0x47}
	gifMagic   = []byte{0x47, 0x49, 0x46}
	webPMagic1 = []byte{0x52, 0x49, 0x46, 0x46}
	webPMagic2 = []byte{0x57, 0x45, 0x42, 0x50}
)

func MagicType(filename string) (fastimage.Type, error) {
	file, err := os.Open(filename)
	if err != nil {
		return fastimage.Unknown, l10n.E("Failed to open file %s to detect magic number: %v", filename, err)
	}
	defer file.Close()

	magic := make([]byte, 12)
	_, err = file.Read(magic)
	if err != nil {
		return fastimage.Unknown, l10n.E("Failed to read magic number from file %s: %v", filename, err)
	}

	if bytes.Equal(magic[:2], jpegMagic) {
		return fastimage.JPEG, nil
	} else if bytes.Equal(magic[:4], pngMagic) {
		return fastimage.PNG, nil
	} else if bytes.Equal(magic[:3], gifMagic) {
		return fastimage.GIF, nil
	} else if bytes.Equal(magic[:4], webPMagic1) && bytes.Equal(magic[8:12], webPMagic2) {
		return fastimage.WEBP, nil
	}

	return fastimage.Unknown, nil
}
