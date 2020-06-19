package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func RawPost(url string, reqByte []byte, headers map[string]string, contentType string) ([]byte, int){
	var req *http.Request
	var resp *http.Response
	var err error
	var body []byte
	if req,err = http.NewRequest("POST", url, bytes.NewBuffer(reqByte)); err!=nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", contentType)
	for k,v := range headers {
		req.Header.Add(k,v)
	}
	client := http.Client{}
	if resp,err = client.Do(req); err!=nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if body,err=ioutil.ReadAll(resp.Body);err!=nil {
		log.Fatalln(err)
	}
	return body, resp.StatusCode
}
