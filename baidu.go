package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Baidu struct {
	CookieJar http.CookieJar
	Headers map[string]string
}


/**

HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Type: text/html
Date: Wed, 06 Mar 2019 03:14:33 GMT
P3p: CP=" OTI DSP COR IVA OUR IND COM "
Server: Apache
Set-Cookie: locale=zh; expires=Tue, 31-Dec-2019 03:14:33 GMT; path=/; domain=.baidu.com
BAIDUID=6E258D3289B51448249B81C1591A370B:FG=1; expires=Thu, 05-Mar-20 03:14:33 GMT; max-age=31536000; path=/; domain=.baidu.com; version=1
Vary: Accept-Encoding
Transfer-Encoding: chunked
 */
func (b *Baidu) init() {

	jar, _ := cookiejar.New(nil)
	client := http.Client{
		Jar:jar,
	}
	request, err := http.NewRequest("GET","https://fanyi.baidu.com/",strings.NewReader(""))
	if err != nil{
		panic(err)
	}
	for key,val := range b.Headers{
		request.Header.Set(key,val)
	}
	response, err := client.Do(request)
	fmt.Println(response.Cookies())
	url,err := url.Parse("https://fanyi.baidu.com/")
	if err != nil {
		log.Fatal(err)
	}
	jar.SetCookies(url,response.Cookies())
	fmt.Println(jar)
}

func main(){
	b := &Baidu{}
	b.Headers = map[string]string{
		"host":"fanyi.baidu.com",
		"User-Agent":"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:65.0) Gecko/20100101 Firefox/65.0",
		"Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection": "keep-alive",
		"Upgrade-Insecure-Requests": "1",
	}
	b.init()
}