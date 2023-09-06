package recaptcha

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	url := "http://127.0.0.1:60080/v1/captcha/image"

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var captcha, _ = filepath.Abs("testdata/captcha/1133.png")
	// Add the captcha images to the request
	captchaFiles := []string{captcha}
	for _, filename := range captchaFiles {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		// Create a form field for each image file
		part, err := writer.CreateFormFile("captcha[]", filename)
		if err != nil {
			panic(err)
		}

		// Copy the image file content to the form field
		_, err = io.Copy(part, file)
		if err != nil {
			panic(err)
		}

		file.Close()
	}

	// Add the options JSON to the request as a form field
	options := map[string]interface{}{
		"psm":       13,
		"languages": []string{"eng"},
		"whitelist": "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	optionsBytes, err := json.Marshal(options)
	if err != nil {
		panic(err)
	}
	err = writer.WriteField("options", string(optionsBytes))
	if err != nil {
		panic(err)
	}

	// Close multipart writer to finalize the request body
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}

	// Set the request header
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func TestClient_Recognition(t *testing.T) {
	cli := NewClient("http://127.0.0.1:60080/v1/captcha/image", Options{
		PageSegMode: 13,
		Languages:   []string{"eng"},
		Whitelist:   "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	})

	var captcha1381, _ = filepath.Abs("testdata/captcha/1381.png")
	file1381, err := os.Open(captcha1381)
	defer file1381.Close()
	if err != nil {
		panic(err)
	}

	cli.PushCaptcha(Captcha{Content: file1381, ID: file1381.Name()})

	var captcha1133, _ = filepath.Abs("testdata/captcha/1133.png")
	file1133, err := os.Open(captcha1133)
	defer file1133.Close()
	if err != nil {
		panic(err)
	}
	cli.PushCaptcha(Captcha{Content: file1133, ID: file1133.Name()})
	rec, err := cli.Image()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rec)
}
func TestClient_Recognition2(t *testing.T) {
	cli := NewClient("http://127.0.0.1:60080/v1/captcha/base64", Options{
		PageSegMode: 13,
		Languages:   []string{"eng"},
		Whitelist:   "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	})

	var captcha1381, _ = filepath.Abs("testdata/captcha/1381.png")
	file1381, err := os.ReadFile(captcha1381)
	if err != nil {
		panic(err)
	}

	encoded1381 := base64.StdEncoding.EncodeToString(file1381)

	cli.PushCaptcha(Captcha{Content: strings.NewReader(encoded1381), ID: captcha1381})

	var captcha1133, _ = filepath.Abs("testdata/captcha/1133.png")
	file1133, err := os.ReadFile(captcha1133)
	if err != nil {
		panic(err)
	}

	encoded1133 := base64.StdEncoding.EncodeToString(file1133)

	cli.PushCaptcha(Captcha{Content: strings.NewReader(encoded1133), ID: captcha1133})

	rec, err := cli.Base64()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rec)
}
func TestClient_Recognition3(t *testing.T) {
	cli := NewClient("http://127.0.0.1:60080/v1/captcha/image", Options{
		PageSegMode: 7,
		Languages:   []string{"eng"},
		Whitelist:   "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	})
	for x := 0; x < 10; x++ {

		resp, err := http.Get("http://192.168.111.133:8080/server/index.php?s=/api/common/verify&rand=0.721628902912116&rand=0.2664791129731017&rand=0.5896513056744962")
		if err != nil {
			panic(err)
		}
		cap, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		t := time.Now()

		id := fmt.Sprintf("%d", t.Nanosecond())

		cli.PushCaptcha(Captcha{Content: bytes.NewReader(cap), ID: id})

		// 保存到本地
		if err := os.WriteFile("testdata/output/"+id+".png", cap, 664); err != nil {
			panic(err)
		}
	}

	rec, err := cli.Image()
	if err != nil {
		log.Fatal(err)
	}

	for x := range rec {
		if rec[x].Text == "" {
			continue
		}
		os.Rename("testdata/output/"+rec[x].ID+".png", "testdata/output/"+rec[x].Text+".png")
	}
}
