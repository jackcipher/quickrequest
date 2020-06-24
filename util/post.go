package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func PostForm(urlAddr string, params map[string]string, headers map[string]string) ([]byte, int) {
	data := url.Values{}
	for k,v := range params {
		data.Set(k,v)
	}

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlAddr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil,resp.StatusCode
	}
	defer resp.Body.Close()
	var body []byte
	if body,err=ioutil.ReadAll(resp.Body);err!=nil {
		log.Fatalln(err)
	}
	return body, resp.StatusCode
}
