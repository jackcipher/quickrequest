package main

import (
	"fmt"
	"github.com/jackcipher/quickrequest/util"
)

func main()  {
	resp,code := util.PostForm("http://127.0.0.1:8080", map[string]string{
		"name": "test",
	}, nil)
	fmt.Println(string(resp),code)
}
