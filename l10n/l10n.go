package l10n

import (
	"fmt"
	"os"

	"golang.org/x/text/language"
)

type LexiconMap map[string]string
type WorldMap map[string]LexiconMap

var (
	Language = "en"
)

func Register(lang string, lex LexiconMap) {
	l, ok := World[lang]
	if !ok {
		World[lang] = lex
		return
	}

	for k, v := range lex {
		l[k] = v
	}
}

func T(phase string) string {
	lex, ok := World[Language]
	if !ok {
		return phase
	}

	t, ok := lex[phase]
	if !ok {
		return phase
	}

	return t
}

func F(phase string, args ...interface{}) string {
	return fmt.Sprintf(T(phase), args...)
}

func E(phase string, args ...interface{}) error {
	template := T(phase)
	return fmt.Errorf(F(template, args...))
}

func DetectLanguage() {
	langs := []string{
		os.Getenv("LANGUAGE"),
		os.Getenv("LC_ALL"),
		os.Getenv("LC_MESSAGES"),
		os.Getenv("LANG"),
	}

	matcher := language.NewMatcher([]language.Tag{
		language.Japanese,
		language.English,
	})

	Language = "en"
	for _, l := range langs {
		if l != "" {
			tag, _ := language.MatchStrings(matcher, l)
			if tag == language.Japanese {
				Language = "ja"
				break
			}
		}
	}
}

var World = WorldMap{
	"ja": LexiconMap{
		// app/exclusive.go
		"Failed to make directory for lock file %s: %v": "ロックファイル %s のディレクトリ作成に失敗しました: %v",
		"Failed to create lock file %s: %v":             "ロックファイル %s の作成に失敗しました: %v",

		// app/imagedouble.go
		"ImageDouble failed to get stat src file %s: %v":    "オリジナルファイル %s の情報取得に失敗しました: %v",
		"ImageDouble failed to get stat double file %s: %v": "代替ファイル %s の情報取得に失敗しました: %v",
		"ImageDouble failed to get stat error file %s: %v":  "エラーファイル %s の情報取得に失敗しました: %v",
		"Skip %s because of no update":                      "%s は更新がないためスキップします",
		"Failed to convert %s: %v":                          "%s の変換に失敗しました: %v",
		"Converted %s":                                      "%s を変換しました",

		// beside/beside.go
		"Failed to control exclusive: %v":              "排他制御に失敗しました: %v",
		"The Application seems to be already running.": "アプリケーションは既に起動しているようです",
		"Got an error while walking %s: %v":            "%s の走査中にエラーを取得しました: %v",
		"Failed to get relative path %s in %s: %v":     "%s の相対パス取得に失敗しました: %v",
		"Failed to ensure the image double for %s: %v": "%s の代替ファイルの確保に失敗しました: %v",

		// converter/converter.go
		"Failed to create a tmp file: %v":                                          "一時ファイルの作成に失敗しました: %v",
		"Failed to get stat of src file %s: %v":                                    "元ファイル %s の情報取得に失敗しました: %v",
		"Failed to get stat of tmp file %s: %v":                                    "一時ファイル %s の情報取得に失敗しました: %v",
		"File size got larger by conversion %d > %d":                               "変換によりファイルサイズが大きくなりました %d > %d",
		"Failed to get image detail of %s: %v":                                     "%s の画像情報取得に失敗しました: %v",
		"Image type is not %s expected %s":                                         "画像タイプが %s で期待される %s ではありません",
		"Image dimensions are different between input(%dx%d) and converted(%dx%d)": "画像サイズが入力(%dx%d)と変換後(%dx%d)で異なります",

		// imagetype/fastimage.go
		"Failed to open file %s for fastimage: %v":          "fastimage 用にファイル %s を開けませんでした: %v",
		"Failed to open file %s to detect magic number: %v": "マジックナンバー検出のためにファイル %s を開けませんでした: %v",
		"Failed to read magic number from file %s: %v":      "ファイル %s からマジックナンバーを読み取れませんでした: %v",

		// mirror/mirror.go
		// "Failed to control exclusive: %v": 						"排他制御に失敗しました: %v",
		"Failed to get absolute path for %s: %v": "絶対パス取得に失敗しました: %s: %v",
		// "Failed to get relative path %s in %s: %v": "%s の相対パス取得に失敗しました: %s: %v",
		"Failed to ensure image double for %s: %v": "%s の代替ファイルの確保に失敗しました: %v",

		// command.go
		"Toolkit for converting conventionally formatted Web images to WebP":                   "従来フォーマットの画像ファイルをWebPに変換するツールキット",
		"Log level to display (trace, debug, info, warn, error, fatal, silent)":                "表示するログレベル (trace, debug, info, warn, error, fatal, silent)",
		"Umask for file and directory creation":                                                "ファイルとディレクトリの作成のための Umask 値 (例 0022)",
		"Print the version number":                                                             "バージョン番号を表示",
		"Alias for cwebp command of libwebp":                                                   "libwebp の cwebp コマンドのエイリアス",
		"Alias for dwebp command of libwebp":                                                   "libwebp の dwebp コマンドのエイリアス",
		"Alias for gif2webp command of libwebp":                                                "libwebp の gif2webp コマンドのエイリアス",
		"Alias for webpinfo command of libwebp":                                                "libwebp の webpinfo コマンドのエイリアス",
		"Convert a single image file to WebP format":                                           "単一の画像ファイルをWebPに変換",
		"Convert image files under the directory to WebP format (or Reverse) as another tree":  "ディレクトリ配下の画像ファイルをWebP(またはその逆)に変換して別のツリー構造に同期",
		"Exclusive lock file path to control exclusive":                                        "排他制御のためのロックファイルパス",
		"Exclusive lock expires (e.g. 1h, 1d)":                                                 "排他制御の有効期限 (例: 1h, 1d)",
		"Failed to parse %s as a duration":                                                     "%s を期間として解釈できませんでした",
		"Failed to parse %s as a umask value":                                                  "%s をUmask値として解釈できませんでした",
		"Convert image files under the directory to WebP format (or Reverse) beside each file": "ディレクトリ配下の画像ファイルをWebP(またはその逆)に変換して元ファイルの隣に配置",
	},
}
