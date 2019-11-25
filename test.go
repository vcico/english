package  main

import (
	"strings"
	"fmt"
	//"regexp"
	//"net/url"
)

func main(){
	str := "15000英语日常口语词汇 6-2-6	一般的社会行为"
	fmt.Printf("%s.txt\n",str)
	str = strings.Replace(str,"15000英语日常口语词汇","",-1)
	fmt.Printf("%s.txt\n",str)
	str = strings.Trim(str ," ")
	fmt.Printf("%s.txt\n",str)
	str = strings.Replace(str," ","",-1)

	fmt.Printf("%s.txt\n",str)

	//re := regexp.MustCompile(`\[\d{2}:\d{2}\.\d{2}\]`)
	//fmt.Println(re.ReplaceAllString("[00:42.87]bear 忍受", ""))

	//u,err := url.Parse("http://www.guancha.cn/kouyu/basic/15000kych/index_2.html")
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(u)
	//fmt.Println(u.Host)
	////u.Host = "www.en8848.com.cn"
	//fmt.Println(u.Path)
	//fmt.Println(u.ForceQuery)
	//fmt.Println(u.Fragment)
	//fmt.Println(u.Opaque)
	//fmt.Println(u.RawPath)
	//fmt.Println(u.RawQuery)
	//fmt.Println("--",u.Scheme)
	//fmt.Println(u.User)
	//fmt.Println(u)

}