//go:build linux
// +build linux

package recaptcha

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServerRun(t *testing.T) {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.POST("/v1/captcha/image", RecCaptchaImageHandle)
	r.POST("/v1/captcha/base64", RecCaptchaBase64Handle)

	//r.Run(":8085")
}
