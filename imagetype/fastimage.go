package imagetype

import (
	"os"

	"github.com/ideamans/webpkit/l10n"
	"github.com/phuslu/fastimage"
)

func FastImage(path string) (fastimage.Info, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return fastimage.Info{}, l10n.E("Failed to open file %s for fastimage: %v", path, err)
	}

	return fastimage.GetInfo(data), nil
}
