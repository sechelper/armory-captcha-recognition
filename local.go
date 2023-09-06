package recaptcha

import (
	"io"
	"os"
	"os/exec"
)

var DefaultTesseract = New("tesseract")

type Tesseract struct {
	path   string
	stdout io.Writer
	stderr io.Writer
}

func New(path string) *Tesseract {
	return &Tesseract{path: path}
}

func (tes *Tesseract) SetPath(path string) error {
	path, err := exec.LookPath(path)
	tes.path = path
	return err
}

func (tes *Tesseract) SetTessdata(tessdata string) error {
	return os.Setenv("TESSDATA_PREFIX", tessdata)
}

// Run
//
//	options => map[string]interface{}{
//		"-c":    []string{"tessedit_char_whitelist=1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
//		"--oem": "1",
//		"--psm": "13",
//	}
func (tes *Tesseract) Run(captcha string, options map[string]interface{}, std ...*os.File) error {

	arg := []string{captcha, captcha}
	arg = append(arg, jsonOverOptions(options)...)

	cmd := exec.Command(tes.path, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if len(std) == 2 {
		cmd.Stdout = std[0]
		cmd.Stderr = std[1]
	}

	return cmd.Run()
}

func jsonOverOptions(options map[string]interface{}) []string {
	var arg []string
	for k, v := range options {
		switch v.(type) {
		case string:
			arg = append(arg, k)
			arg = append(arg, v.(string))
		case []string:
			for x := range v.([]string) {
				arg = append(arg, k)
				arg = append(arg, v.([]string)[x])
			}
		}
	}
	return arg
}
