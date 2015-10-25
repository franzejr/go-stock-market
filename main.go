package main

import (
  "fmt"
	"os"
	"strings"
  curl "github.com/andelf/go-curl"
)

func main() {
	url := fmt.Sprintf("http://dev.markitondemand.com/MODApis/Api/v2/Quote/jsonp?symbol=%v&callback=my", os.Args[1])
  easy := curl.EasyInit()
  defer easy.Cleanup()
  easy.Setopt(curl.OPT_URL, url)

	handleData := func (buf []byte, userdata interface{}) bool{
		s := string(buf)
		print(s)
		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, handleData)
	easy.Perform()
}

func print(s string){
	s1 := strings.Split(s,"my({")
	s2 := strings.Split(s1[1],"})")
	splits := strings.Split(s2[0],",")

	for _,line := range splits {
		fmt.Println(line)
	}
}
