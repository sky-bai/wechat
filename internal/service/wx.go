package service

import (
	"net/http"
)

func WxBase() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//// 0.先判断是哪一个公众号的请求 先获取到公众号信息
		//
		//// 1.先判断是否是微信服务器发来的请求
		//timestamp := server.Query("timestamp")
		//nonce := server.Query("nonce")
		//signature := server.Query("signature")
		////log.Debugf("validate signature, timestamp=%s, nonce=%s", timestamp, nonce)
		//if signature == util.Signature(, timestamp, nonce) {
		//
		//}
		//
		//// 2.判断请求方法 如果是Get请求 则返回echostr
		//if req.Method != "GET" {
		//	//
		//
		//} else {
		//	// 如果是post请求就解析xml数据
		//	var rawXMLMsgBytes []byte
		//	rawXMLMsgBytes, err := io.ReadAll(r.Body)
		//	if err != nil {
		//		w.WriteHeader(403)
		//		w.Write([]byte("读取body失败"))
		//	}
		//}
		//
		//w.WriteHeader(200)
		//w.Write([]byte("ok11111"))
	}
}

// 1.中间件里面是否可以获取数据库对象
// 2.中间件里面是否可以获取请求的参数
