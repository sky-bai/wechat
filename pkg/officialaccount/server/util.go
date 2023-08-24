package server

//
//import (
//	"encoding/xml"
//	"net/http"
//)
//
//var xmlContentType = []string{"application/xml; charset=utf-8"}
//var plainContentType = []string{"text/plain; charset=utf-8"}
//
//func writeContextType(w http.ResponseWriter, value []string) {
//	header := w.Header()
//	if val := header["Content-Type"]; len(val) == 0 {
//		header["Content-Type"] = value
//	}
//}
//
//// Render render from bytes
//func Render(w http.ResponseWriter, bytes []byte) error {
//	// debug
//	// fmt.Println("response msg = ", string(bytes))
//	w.WriteHeader(200)
//	_, err := w.Write(bytes)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// String render from string
//func String(str string) {
//	writeContextType(srv.Writer, plainContentType)
//	Render([]byte(str))
//} // 但是这个好像是适用于
//
//// XML render to xml
//func XML(obj interface{}) {
//	writeContextType(srv.Writer, xmlContentType)
//	bytes, err := xml.Marshal(obj) //
//	if err != nil {
//		panic(err)
//	}
//	srv.Render(bytes)
//}
//
//// Query returns the keyed url query value if it exists
//func Query(key string) string {
//	value, _ := srv.GetQuery(key)
//	return value
//}
//
//// GetQuery is like Query(), it returns the keyed url query value
//func GetQuery(key string) (string, bool) {
//	req := srv.Request
//	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
//		return values[0], true
//	}
//	return "", false
//}
