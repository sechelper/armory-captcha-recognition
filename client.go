package recaptcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
)

type Captcha struct {
	ID      string
	Content io.Reader
}

type RecaptchaClient struct {
	url     string
	captcha []Captcha
	options Options
}

func NewClient(url string, options Options) *RecaptchaClient {
	return &RecaptchaClient{
		url:     url,
		options: options,
	}
}

func (c *RecaptchaClient) PushCaptcha(captcha Captcha) {
	c.captcha = append(c.captcha, captcha)
}

func (c *RecaptchaClient) Image() ([]CaptchaText, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for x := range c.captcha {

		part, err := writer.CreateFormFile("captcha[]", c.captcha[x].ID)
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(part, c.captcha[x].Content); err != nil {
			return nil, err
		}
	}

	opt, err := json.Marshal(c.options)
	if err != nil {
		return nil, err
	}
	err = writer.WriteField("options", string(opt))
	if err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}
	return c.request(body, writer.FormDataContentType())
}

func (c *RecaptchaClient) Base64() ([]CaptchaText, error) {
	var ct []CaptchaText
	for x := range c.captcha {
		captchaBase64, err := io.ReadAll(c.captcha[x].Content)
		if err != nil {
			return nil, err
		}
		ct = append(ct, CaptchaText{
			ID:   c.captcha[x].ID,
			Text: string(captchaBase64),
		})
	}
	body, err := json.Marshal(CaptchaBase64{
		Options: c.options,
		Captcha: ct,
	})
	if err != nil {
		return nil, err
	}
	return c.request(bytes.NewReader(body), "application/json")
}

func (c *RecaptchaClient) request(body io.Reader, contentType string) ([]CaptchaText, error) {
	req, err := http.NewRequest("POST", c.url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := new(Result)
	if err = json.Unmarshal(b, result); err != nil {
		return nil, err
	}
	if result.Status != Success {
		return nil, errors.New(result.Message)
	}
	rcs := make([]CaptchaText, 0)
	for x := range result.Data.([]interface{}) {
		rcs = append(rcs, CaptchaText{
			ID:   result.Data.([]interface{})[x].(map[string]interface{})["id"].(string),
			Text: result.Data.([]interface{})[x].(map[string]interface{})["text"].(string),
		})
	}
	return rcs, nil
}
