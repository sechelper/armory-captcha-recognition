//go:build linux
// +build linux

package recaptcha

import (
	"bytes"
	"gocv.io/x/gocv"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

func ThresholdBinary(buf []byte) ([]byte, error) {
	// 读取图片验证码
	img, err := gocv.IMDecode(buf, gocv.IMReadColor)
	if err != nil {
		return nil, err
	}

	// 转换为灰度图像
	gray := gocv.NewMat()
	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

	// 对图像进行二值化处理
	binary := gocv.NewMat()
	gocv.Threshold(gray, &binary, 200, 255, gocv.ThresholdBinary)

	// 对二值化后的图像进行降噪处理，使用高斯模糊
	result := gocv.NewMat()
	gocv.GaussianBlur(binary, &result, image.Point{X: 3, Y: 3}, 0, 0, gocv.BorderDefault)

	_, ext, err := image.DecodeConfig(bytes.NewReader(buf))
	if err != nil { // 未知文件类型
		return nil, err
	}

	// 将降噪后的图像编码为字节数据
	data, err := gocv.IMEncode(gocv.FileExt("."+ext), result)

	if err != nil {
		return nil, err
	}
	return data.GetBytes(), nil
	// 保存降噪后的图像
	//gocv.IMWrite("result.jpg", result)
}
