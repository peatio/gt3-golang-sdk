package geetest

type ProcessRes struct {
	Success    int8   `json:"success"`
	CaptchaID  string `json:"gt"`
	Challenge  string `json:"challenge"`
	NewCaptcha int    `json:"new_captcha"`
}

type ValidateRes struct {
	Seccode          string            `json:"seccode"`
	ModelProbability int               `json:"model_probability"`
	Uniformity       bool              `json:"uniformity"`
	OriginInfo       map[string]string `json:"origin_info"`
	Duration         float64           `json:"duration"`
}
