package  main

/**
 * 15000英语日常口语词汇 采集 http://www.en8848.com.cn/kouyu/basic/15000kych/index.html
 * todo http://www.en8848.com.cn/kouyu/basic/five/
		http://www.en8848.com.cn/kouyu/slang/mgrmt-jtgyy/
		http://www.en8848.com.cn/kouyu/basic/kyyf/
		http://www.yygrammar.com
		http://en-grammar.xiao84.com/middle/
		https://wenku.baidu.com/view/050b2463caaedd3383c4d38a.html
		http://www.fltagrammar.com/
 */


import (
	"os"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
	"regexp"
	"net/url"
	"time"
)

func savePage(dom *html.Node) error {
	title := htmlquery.InnerText(dom) // 15000英语日常口语词汇 10-1-14 体操
	title = strings.Replace(title,"15000英语日常口语词汇","",-1)
	title = strings.Trim(title ," ")
	title = strings.Replace(title," ","",-1)
	title = strings.Replace(title,"\t","",-1)
	filename := fmt.Sprintf("./en8848/%s.txt",title)

	url := htmlquery.SelectAttr(dom, "href")
	time.Sleep(1)
	doc, err := htmlquery.LoadURL(url)
	if err != nil{
		return err
	}
	f,err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0755)
	if err != nil{
		return err
	}
	defer func() {
		fmt.Println("close file : ",filename)
		f.Close()
		if err := recover();err != nil {
			fmt.Println("detail page cant find content: ", url )
			fmt.Println(err)
			return
		}
	}()
	content := htmlquery.InnerText(htmlquery.FindOne(doc,"//textarea[@id='lrc_content']"))
	content = strings.Replace(content,"[al:15000英语单词口袋书]\n[by:小e英语www.en8848.com.cn]","",-1)
	re := regexp.MustCompile(`\[\d{2}:\d{2}\.\d{2}\]`)
	off, err := f.WriteString(strings.Trim(re.ReplaceAllString(content,"")," "))
	if err != nil {
		fmt.Printf("file write string err : %v\n", err)
	}else{
		fmt.Printf("file write string success  off is %v\n", off)
	}
	/*
	words := htmlquery.Find(doc,"//div[@id='articlebody']/div[@class='info-qh']/div[@class='qh_en']|/div[@id='articlebody']/div[@class='info-qh']/div[@class='qh_zg']")
	for i,n:=0,len(words); i < n; i+=2 { // 常见的 for 循环，支持初始化语句。
		off, err := f.WriteString(fmt.Sprintf("%s   %s\n",htmlquery.InnerText(words[i]),htmlquery.InnerText(words[i+1])))
		if err != nil {
			fmt.Printf("file2 write string err : %v\n", err)
		}else{
			fmt.Printf("file2 write string success , off is %v\n", off)
		}
	}*/
	return nil
}

func main(){

	var urlS string
	urlS = "http://www.en8848.com.cn/kouyu/basic/15000kych/index_5.html"
	for {
		if urlS != ""{
			time.Sleep(1)
			fmt.Println("###########################################\n load list url :  ", urlS)
			doc, err := htmlquery.LoadURL(urlS)
			if err != nil{
				panic(err)
			}
			list := htmlquery.Find(doc, "//div[@class='ch_content']/div[@class='ch_lii']/div[@class='ch_lii_left']/a")
			for i,row := range list{
				fmt.Printf("-- \n%d %s(%s)\n", i, htmlquery.InnerText(row), htmlquery.SelectAttr(row, "href"))
				if err := savePage(row); err != nil{
					panic(err)
				}
				//break // @todo
			}
			//break  // @todo
			urlS = getNextPageUrl(doc)
		}else{
			break
		}
	}

/*
	//doc, err := htmlquery.LoadURL("http://www.en8848.com.cn/kouyu/basic/15000kych/index.html")
	filePath := "./crawl.html"
	file , err := os.OpenFile(filePath,os.O_RDONLY,0755)
	if err != nil{
		panic(err)
	}
	//htmlquery.
	doc, err := htmlquery.Parse(file)
	if err != nil{
		panic(err)
	}
	fmt.Println(getNextPageUrl(doc))

 */

}

func getNextPageUrl(doc *html.Node) string {
	defer func() {
		if err := recover();err != nil {
			fmt.Println("cant find next page")
			fmt.Println(err)
			return
		}
	}()
	next := htmlquery.FindOne(doc, "//div[@class='ch_content']/div[@class='ch_pagebar']/a[last()]")
	//fmt.Println("next page ")
	//fmt.Println(next)
	if htmlquery.InnerText(next) != "下一页"{
		return ""
	}
	href := htmlquery.SelectAttr(next, "href")
	u,err := url.Parse(href)
	if err != nil{
		panic(err)
	}
	if u.Host == ""{
		u.Host = "www.en8848.com.cn"
		u.Scheme = "http"
	}
	return u.String()
}