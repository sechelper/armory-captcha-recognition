package recaptcha

type Status string

const (
	Success Status = "success"
	Error          = "error"
)

type Result struct {
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CaptchaBase64 struct {
	Captcha []CaptchaText `json:"captcha"`
	Options
}

type CaptchaText struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Options struct {
	Languages   []string `json:"languages"`
	Whitelist   string   `json:"whitelist"`
	PageSegMode int      `json:"psm"`
}
