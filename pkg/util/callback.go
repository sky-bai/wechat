package util

// SignatureOptions 微信服务器验证参数
type SignatureOptions struct {
	Signature string `form:"signature"`
	TimeStamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
	Token     string `form:"token"`
}

func Validate(options SignatureOptions) bool {
	return options.Signature == Signature(options.Token, options.TimeStamp, options.Nonce, options.EchoStr)
}
