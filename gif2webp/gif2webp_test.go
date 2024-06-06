package gif2webp

import (
	"testing"

	"github.com/ideamans/webpkit/imagetype"
	"github.com/phuslu/fastimage"
)

func TestGif2WebP(t *testing.T) {
	tmp := t.TempDir()

	srcPath := "../testdata/simple/simple.gif"
	destPath := tmp + "/sample.webp"
	code := Gif2WebP("-quiet", "-o", destPath, srcPath)
	if code != 0 {
		t.Errorf("gif2webp command exited with code %d", code)
	}

	gif, err := imagetype.FastImage(srcPath)
	if err != nil {
		t.Errorf("Failed to read gif file: %v", err)
	}
	webp, err := imagetype.FastImage(destPath)
	if err != nil {
		t.Errorf("Failed to read webp file: %v", err)
	}

	if webp.Type != fastimage.WEBP {
		t.Errorf("Output file is not a webp image: %v", webp.Type)
	}
	if gif.Width != webp.Width || gif.Height != webp.Height {
		t.Errorf("Image dimensions are not %dx%d but %dx%d", gif.Width, gif.Height, webp.Width, webp.Height)
	}
}
