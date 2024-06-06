package webpinfo

import "testing"

func TestWebPInfo(t *testing.T) {
	code := WebPInfo("-quiet", "../testdata/simple/simple.webp")
	if code != 0 {
		t.Errorf("webpinfo failed with code %d", code)
	}
}
