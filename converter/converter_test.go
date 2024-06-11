package converter

import (
	"path/filepath"
	"testing"
)

func TestConverter(t *testing.T) {
	cases := []struct {
		path      string
		outputExt string
	}{
		{
			path:      "simple/simple.jpg",
			outputExt: ".webp",
		},
		{
			path:      "simple/simple.png",
			outputExt: ".webp",
		},
		{
			path:      "simple/simple.gif",
			outputExt: ".webp",
		},
		{
			path:      "simple/simple.webp",
			outputExt: ".png",
		},
		{
			path:      "simple/simple.webp",
			outputExt: ".jpg",
		},
	}

	tmp := t.TempDir()
	converter := NewConverter(ConverterConfig{Verify: true})

	for _, c := range cases {
		inputPath := filepath.Join("../testdata", c.path)
		outputPath := filepath.Join(tmp, inputPath+c.outputExt)
		err := converter.Convert(inputPath, outputPath)
		if err != nil {
			t.Errorf("Failed to convert %s: %v", c.path, err)
		}
	}
}
