package l10n

import "testing"

func TestPhrase(t *testing.T) {
	Register("en", LexiconMap{
		"hello %s": "Hello, %s",
	})

	Register("ja", LexiconMap{
		"hello %s": "こんにちは、%s",
	})

	Language = "en"
	if T("hello %s") != "Hello, %s" {
		t.Errorf("T() failed")
	}
	if F("hello %s", "world") != "Hello, world" {
		t.Errorf("F() failed")
	}

	Language = "ja"
	if T("hello %s") != "こんにちは、%s" {
		t.Errorf("T() failed")
	}
	if F("hello %s", "世界") != "こんにちは、世界" {
		t.Errorf("F() failed")
	}

	Language = "unknown"
	if T("hello %s") != "hello %s" {
		t.Errorf("T() failed")
	}
	if F("hello %s", "world") != "hello world" {
		t.Errorf("F() failed")
	}
}
