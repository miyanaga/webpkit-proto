package imagetype

import (
	"path/filepath"
	"testing"

	"github.com/phuslu/fastimage"
)

func TestImageType(t *testing.T) {
	cases := []struct {
		path string
		it   fastimage.Type
	}{
		{
			path: "simple/simple.jpg",
			it:   fastimage.JPEG,
		},
		{
			path: "simple/simple.png",
			it:   fastimage.PNG,
		},
		{
			path: "simple/simple.gif",
			it:   fastimage.GIF,
		},
		{
			path: "simple/simple.webp",
			it:   fastimage.WEBP,
		},
	}

	for _, c := range cases {
		it, err := MagicType(filepath.Join("../testdata", c.path))
		if err != nil {
			t.Errorf("Failed to get image type of %s: %v", c.path, err)
		}
		if it != c.it {
			t.Errorf("Image type of %s is not %v but %v", c.path, c.it, it)
		}
	}
}
