package quickrequest

import (
	"encoding/json"
	"fmt"
	"github.com/jackcipher/quickrequest/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


/**
Normal Get
 */
func NewGet(url string, params map[string]string, headers map[string]string) []byte{
	var req *http.Request
	var resp *http.Response
	var err error
	var body []byte
	var inline []string
	for k,v := range params {
		inline = append(inline, fmt.Sprintf("%s=%s", k, v))
	}
	if len(inline)>0 {
		url += "?" + strings.Join(inline, "&")
	}
	if req,err = http.NewRequest("GET", url, nil); err!=nil {
		log.Fatalln(err)
	}
	for k,v := range headers {
		req.Header.Add(k,v)
	}
	client := http.Client{}
	if resp,err = client.Do(req); err!=nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if body,err = ioutil.ReadAll(resp.Body); err!=nil {
		log.Fatalln(err)
	}
	return body
}

/**
Form Post
 */
func PostParamsJson(url string, params map[string]string, headers map[string]string) ([]byte, int){
	var requestByte []byte
	var err error
	if requestByte,err = json.Marshal(params);err!=nil {
		log.Fatalln(err)
	}
	return util.RawPost(url, requestByte, headers, "application/json")
}

/**
Post Json
 */
func PostJson(url string, requestJson []byte, headers map[string]string) ([]byte,int) {
	return util.RawPost(url, requestJson, headers, "application/json")
}

func PostForm(url string, params map[string]string, headers map[string]string) ([]byte,int)  {
	var requestByte []byte
	var err error
	if requestByte,err = json.Marshal(params);err!=nil {
		log.Fatalln(err)
	}
	return util.RawPost(url, requestByte, headers, "x-www-form-urlencode")
}