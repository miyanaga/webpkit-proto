package imagetype

import (
	"path/filepath"
	"testing"

	"github.com/phuslu/fastimage"
)

func TestFastImage(t *testing.T) {
	cases := []struct {
		path   string
		it     fastimage.Type
		width  uint32
		height uint32
	}{
		{
			path:   "simple/simple.jpg",
			it:     fastimage.JPEG,
			width:  240,
			height: 214,
		},
		{
			path:   "simple/simple.png",
			it:     fastimage.PNG,
			width:  240,
			height: 214,
		},
		{
			path:   "simple/simple.gif",
			it:     fastimage.GIF,
			width:  466,
			height: 521,
		},
		{
			path:   "simple/simple.webp",
			it:     fastimage.WEBP,
			width:  240,
			height: 214,
		},
	}

	for _, c := range cases {
		ft, err := FastImage(filepath.Join("../testdata", c.path))
		if err != nil {
			t.Errorf("Failed to get image type %s: %v", c.path, err)
		}
		if ft.Type != c.it {
			t.Errorf("Image type of %s is not %v but %v", c.path, c.it, ft.Type)
		}
		if ft.Width != c.width || ft.Height != c.height {
			t.Errorf("Dimensions of %s are not %dx%d but %dx%d", c.path, c.width, c.height, ft.Width, ft.Height)
		}
	}
}
