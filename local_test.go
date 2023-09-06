package recaptcha

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestTesseract_Run(t *testing.T) {

	tessdata, _ := filepath.Abs("testdata")
	os.Setenv("TESSDATA_PREFIX", tessdata)
	options := map[string]interface{}{
		"-c":    []string{"tessedit_char_whitelist=1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		"--oem": "1",
		"--psm": "13",
	}

	if err := DefaultTesseract.Run("testdata/captcha/1133.png", options); err != nil {
		log.Fatal(err)
	}
}
