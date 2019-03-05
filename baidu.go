package main

import (
	"errors"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

/*
破解翻译接口  https://www.jianshu.com/p/5f3177943b91
https://www.codetd.com/article/2164560
https://www.jianshu.com/p/2c333f7ae1c2
*/

/**
 * 单词 百度翻译 ： 音标、中文、短语、例句
 */


func main(){


	// 网址能访问到东西 但只能在chrome里面 。。。。
	if true{
		word := "这个"
		slang := "zh"
		tlang := "en"
		sign := "69663.389934"
		token := "7593175d8bbcc09498870060b2685f7a"
		v2transapi := "https://fanyi.baidu.com/v2transapi?from="+slang+"&to="+tlang+"&query="+url.QueryEscape(word)+"&transtype=translang&simple_means_flag=3&sign="+sign+"&token="+token
		fmt.Println(v2transapi)
		bodyStr,_ := LoadURL(v2transapi)
		fmt.Println(bodyStr)
		ioutil.WriteFile("result.txt",[]byte(bodyStr),0777)
		return
	}

	//gtk
	//token
	//if true{
	//	jsFunc := getJSfunction()
	//	word := "这个" // every%20day
	//	v ,_ := jsFunc.Call(jsFunc,word,"320305.131321201")
	//	sign,_ := v.ToString()
	//	fmt.Println("sign := ",sign)
	//	return ;
	//}

	gtk := ""
	token := ""
	var result []string
	var body []byte
	var err error

	//if body,err = LoadURL("http://fanyi.baidu.com"); err!=nil{
	//	panic(err)
	//}

	body,err = ioutil.ReadFile("fanyi-baidu.com.txt")
	if err != nil{
		panic(err)
	}
	re := regexp.MustCompile("window.gtk = '(.*?)';")
	result = re.FindStringSubmatch(string(body))
	gtk = result[1]
	//for k,v := range re.FindStringSubmatch(string(body)){
	//	fmt.Println(k,"--",v)
	//}
	re = regexp.MustCompile("token: '(.*?)'")
	result = re.FindStringSubmatch(string(body))
	token = result[1]
	//for k,v := range re.FindStringSubmatch(string(body)){
	//	fmt.Println(k,"--",v)
	//}

	jsFunc := getJSfunction()
	word := "这个" // every%20day
	v ,_ := jsFunc.Call(jsFunc,word,gtk)
	sign,_ := v.ToString()
	fmt.Println("tk := ",sign)
	//fmt.Println(body)
	//ioutil.WriteFile("fanyi-baidu.com",[]byte(body),0777)
	slang := "zh"
	tlang := "en"
	sign = "69663.389934"
	token = "7593175d8bbcc09498870060b2685f7a"
	v2transapi := "http://fanyi.baidu.com/v2transapi?from="+slang+"&to="+tlang+"&query="+url.QueryEscape(word)+"&transtype=translang&simple_means_flag=3&sign="+sign+"&token="+token
	bodyStr,err := LoadURL(v2transapi)
	fmt.Println(bodyStr)
	ioutil.WriteFile("result.txt",[]byte(bodyStr),0777)
}

func LoadURL(url string) (string, error) {
	client := &http.Client{}
	request,err := http.NewRequest("get", url, strings.NewReader(""))
	request.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36")
	request.Header.Set("Cookie","'locale=zh; BAIDUID=FC2689968A662FA6104AA311FE89635B:FG=1; from_lang_often=%5B%7B%22value%22%3A%22en%22%2C%22text%22%3A%22%u82F1%u8BED%22%7D%2C%7B%22value%22%3A%22zh%22%2C%22text%22%3A%22%u4E2D%u6587%22%7D%5D; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; to_lang_often=%5B%7B%22value%22%3A%22zh%22%2C%22text%22%3A%22%u4E2D%u6587%22%7D%2C%7B%22value%22%3A%22en%22%2C%22text%22%3A%22%u82F1%u8BED%22%7D%5D'")
	resp, err := client.Do(request)//发送请求
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body),err
}


func getJSfunction() otto.Value{
	var signCode = `function a(r,o){for(var t=0;t<o.length-2;t+=3){var a=o.charAt(t+2);a=a>="a"?a.charCodeAt(0)-87:Number(a),a="+"===o.charAt(t+1)?r>>>a:r<<a,r="+"===o.charAt(t)?r+a&4294967295:r^a}return r}var C=null;var hash=function(r,_gtk){var o=r.length;o>30&&(r=""+r.substr(0,10)+r.substr(Math.floor(o/2)-5,10)+r.substr(-10,10));var t=void 0,t=null!==C?C:(C=_gtk||"")||"";for(var e=t.split("."),h=Number(e[0])||0,i=Number(e[1])||0,d=[],f=0,g=0;g<r.length;g++){var m=r.charCodeAt(g);128>m?d[f++]=m:(2048>m?d[f++]=m>>6|192:(55296===(64512&m)&&g+1<r.length&&56320===(64512&r.charCodeAt(g+1))?(m=65536+((1023&m)<<10)+(1023&r.charCodeAt(++g)),d[f++]=m>>18|240,d[f++]=m>>12&63|128):d[f++]=m>>12|224,d[f++]=m>>6&63|128),d[f++]=63&m|128)}for(var S=h,u="+-a^+6",l="+-3^+b+-f",s=0;s<d.length;s++)S+=d[s],S=a(S,u);return S=a(S,l),S^=i,0>S&&(S=(2147483647&S)+2147483648),S%=1e6,S.toString()+"."+(S^h)}`
	vm := otto.New()
	vm.Run(signCode)
	value, err := vm.Get("hash")
	if  err != nil {
		panic(errors.New("get javascript function tk from Otto fail"))
	}
	if value.IsFunction(){
		return value
	}else{
		panic(errors.New("get javascript tk from Otto is not a  function"))
	}
}
