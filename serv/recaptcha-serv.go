package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sechelper/recaptcha"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.POST("/v1/captcha/image", recaptcha.RecCaptchaImageHandle)
	r.POST("/v1/captcha/base64", recaptcha.RecCaptchaBase64Handle)

	r.Run(":60080")
}
