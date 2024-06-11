package converter

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/ideamans/webpkit/cwebp"
	"github.com/ideamans/webpkit/dwebp"
	"github.com/ideamans/webpkit/gif2webp"
	"github.com/ideamans/webpkit/imagetype"
	"github.com/ideamans/webpkit/l10n"
	"github.com/phuslu/fastimage"
)

type ConverterConfig struct {
	Verify bool
}

type Converter struct {
	config ConverterConfig
}

type ConverterInterface interface {
	Convert(inputPath string, outputPath string) error
}

var (
	Singleton ConverterInterface
)

func init() {
	Singleton = NewConverter(ConverterConfig{Verify: true})
}

func NewConverter(config ConverterConfig) *Converter {
	return &Converter{config: config}
}

type SafeConvertCallback func(tmpPath string) error

func (c *Converter) safeConvert(inputPath, outputPath string, expectedImageType fastimage.Type, shouldBeSmaller bool, cb SafeConvertCallback) error {
	tmp, err := os.CreateTemp("", "webpkit")
	if err != nil {
		return l10n.E("Failed to create a tmp file: %v", err)
	}
	tmp.Close()
	defer os.RemoveAll(tmp.Name())

	err = cb(tmp.Name())
	if err != nil {
		return err
	}

	if shouldBeSmaller {
		srcStat, err := os.Stat(inputPath)
		if err != nil {
			return l10n.E("Failed to get stat of src file %s: %v", inputPath, err)
		}
		tmpStat, err := os.Stat(tmp.Name())
		if err != nil {
			return l10n.E("Failed to get stat of tmp file %s: %v", tmp.Name(), err)
		}
		if tmpStat.Size() > srcStat.Size() {
			return l10n.E("File size got larger by conversion %d > %d", tmpStat.Size(), srcStat.Size())
		}
	}

	if c.config.Verify {
		inputImage, err := imagetype.FastImage(inputPath)
		if err != nil {
			return l10n.E("Failed to get image detail of %s: %v", inputPath, err)
		}

		tmpImage, err := imagetype.FastImage(tmp.Name())
		if err != nil {
			return l10n.E("Failed to get image detail of %s: %v", tmp.Name(), err)
		}

		if tmpImage.Type != expectedImageType {
			return l10n.E("Image type is not %s expected %s", tmpImage.Type, expectedImageType)
		}

		if inputImage.Width != tmpImage.Width || inputImage.Height != tmpImage.Height {
			return l10n.E("Image dimensions are different between input(%dx%d) and converted(%dx%d)", inputImage.Width, inputImage.Height, tmpImage.Width, tmpImage.Height)
		}
	}

	os.MkdirAll(filepath.Dir(outputPath), 0777)
	os.Rename(tmp.Name(), outputPath)

	return nil
}

func (c *Converter) jpegToWebP(inputPath string, outputPath string) error {
	return c.safeConvert(inputPath, outputPath, fastimage.WEBP, true, func(tmpPath string) error {
		code := cwebp.CWebP("-quiet", "-q", "80", "-metadata", "icc", "-o", tmpPath, inputPath)
		if code != 0 {
			return l10n.E("cwebp command exited with code %d", code)
		}
		return nil
	})
}

func (c *Converter) pngToWebP(inputPath string, outputPath string) error {
	return c.safeConvert(inputPath, outputPath, fastimage.WEBP, true, func(tmpPath string) error {
		code := cwebp.CWebP("-quiet", "-lossless", "-metadata", "icc", "-o", tmpPath, inputPath)
		if code != 0 {
			return l10n.E("cwebp command exited with code %d", code)
		}
		return nil
	})
}

func (c *Converter) gifToWebP(inputPath string, outputPath string) error {
	return c.safeConvert(inputPath, outputPath, fastimage.WEBP, true, func(tmpPath string) error {
		code := gif2webp.Gif2WebP("-o", tmpPath, inputPath)
		if code != 0 {
			return l10n.E("gif2webp command exited with code %d", code)
		}
		return nil
	})
}

func (c *Converter) webPToPng(inputPath string, outputPath string) error {
	return c.safeConvert(inputPath, outputPath, fastimage.PNG, false, func(tmpPath string) error {
		code := dwebp.DWebP("-quiet", "-o", tmpPath, inputPath)
		if code != 0 {
			return l10n.E("dwebp command exited with code %d", code)
		}
		return nil
	})
}

func (c *Converter) webPToJpeg(inputPath string, outputPath string) error {
	return c.safeConvert(inputPath, outputPath, fastimage.JPEG, false, func(tmpPath string) error {
		tmp, err := os.CreateTemp("", "webpkit")
		if err != nil {
			return l10n.E("Failed to create a tmp file to write PNG data: %v", err)
		}
		tmp.Close()
		defer os.RemoveAll(tmp.Name())

		// Convert WebP to PNG at first
		code := dwebp.DWebP("-quiet", "-o", tmp.Name(), inputPath)
		if code != 0 {
			return l10n.E("dwebp command exited with code %d", code)
		}

		// Convert the PNG to JPEG
		pngFile, err := os.Open(tmp.Name())
		if err != nil {
			return l10n.E("Failed to open %s as PNG file: %v", tmp.Name(), err)
		}
		defer pngFile.Close()

		pngImg, err := png.Decode(pngFile)
		if err != nil {
			return l10n.E("Failed to decode PNG image: %v", err)
		}

		jpegFile, err := os.Create(tmpPath)
		if err != nil {
			return l10n.E("Failed to create %s as JPEG file: %v", tmpPath, err)
		}
		defer jpegFile.Close()

		err = jpeg.Encode(jpegFile, pngImg, &jpeg.Options{Quality: 80})
		if err != nil {
			return l10n.E("Failed to encode JPEG image: %v", err)
		}

		return nil
	})
}

func (c *Converter) Convert(inputPath string, outputPath string) error {
	it, err := imagetype.MagicType(inputPath)
	if err != nil {
		return l10n.E("Failed to get image type of %s by magic number: %v", inputPath, err)
	}

	destExt := filepath.Ext(outputPath)

	if it == fastimage.JPEG && destExt == ".webp" {
		return c.jpegToWebP(inputPath, outputPath)
	} else if it == fastimage.PNG && destExt == ".webp" {
		return c.pngToWebP(inputPath, outputPath)
	} else if it == fastimage.GIF && destExt == ".webp" {
		return c.gifToWebP(inputPath, outputPath)
	} else if it == fastimage.WEBP && destExt == ".png" {
		return c.webPToPng(inputPath, outputPath)
	} else if it == fastimage.WEBP && destExt == ".jpg" || destExt == ".jpeg" {
		return c.webPToJpeg(inputPath, outputPath)
	} else {
		return l10n.E("Unsupported conversion type pair from %s to %s", it, destExt)
	}
}
