//go:build linux
// +build linux

package recaptcha

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
	"io"
	"log"
)

func Recognition(captcha []byte, opt *Options) string {
	client := gosseract.NewClient()
	defer client.Close()
	captcha, err := ThresholdBinary(captcha) //图像二值化
	if err != nil {
		log.Println(err)
		return ""
	}
	client.SetImageFromBytes(captcha)
	client.SetWhitelist(opt.Whitelist)
	client.Languages = opt.Languages
	if len(opt.Languages) == 0 {
		client.Languages = []string{"eng"}
	}
	client.SetPageSegMode(gosseract.PageSegMode(opt.PageSegMode))
	text, _ := client.Text()
	return text
}

// RecCaptchaImageHandle
// https://gin-gonic.com/docs/examples/upload-file/multiple-file/
func RecCaptchaImageHandle(c *gin.Context) {

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(500, Result{Status: Error, Message: "Invalid multipart form"})
		return
	}
	opt := new(Options)
	if _, ok := form.Value["options"]; !ok {
		c.JSON(500, Result{Status: Error, Message: fmt.Sprintf("Invalid options: %s", form.Value["options"])})
		return
	}

	if err := json.Unmarshal([]byte(form.Value["options"][0]), opt); err != nil {
		c.JSON(500, Result{Status: Error, Message: "Invalid options: " + form.Value["options"][0]})
		return
	}

	files := form.File["captcha[]"]

	res := Result{
		Status:  Success,
		Message: "ok",
		Data:    make([]CaptchaText, 0),
	}

	for _, file := range files {
		op, err := file.Open()
		if err != nil {
			log.Println(err)
			c.JSON(500, Result{Status: Error, Message: "Invalid open file: " + file.Filename})
			return
		}
		content, err := io.ReadAll(op)
		if err != nil {
			c.JSON(500, Result{Status: Error, Message: "Invalid read file: " + file.Filename})
			return
		}

		res.Data = append(res.Data.([]CaptchaText), CaptchaText{
			ID:   file.Filename,
			Text: Recognition(content, opt),
		})
	}

	c.JSON(200, res)
}

func RecCaptchaBase64Handle(c *gin.Context) {
	rb := new(CaptchaBase64)
	if err := c.Bind(rb); err != nil {
		c.JSON(500, Result{Status: Error, Message: "Invalid json " + err.Error()})
		return
	}
	rbs := make([]CaptchaText, 0)

	for x := range rb.Captcha {
		// Decoding the string
		content, err := base64.StdEncoding.DecodeString(rb.Captcha[x].Text)
		var text string
		if err != nil {
			log.Println("Error decoding base64", err)
		} else {
			text = Recognition(content, &rb.Options)
		}

		rbs = append(rbs, CaptchaText{
			ID:   rb.Captcha[x].ID,
			Text: text,
		})
	}
	c.JSON(200, Result{Status: Success, Message: "ok", Data: rbs})
}
