package main

import (
	"flag"
	"io/ioutil"
	"net/http"
)

//go run downloadtest.go -downloadurl 网站名 -TheTargetLocation 下载到哪个位置（默认D:\）
var url1, TheTargetLocation string

func main() {
	flag.PrintDefaults() //打印 flag 的使用帮助信息
	flag.StringVar(&url1, "downloadurl", "err", "the max number of download film")
	flag.Parse() //这个函数会扫描命令行里的参数，并设置 flag,赋值给上面声明的url1
	flag.StringVar(&TheTargetLocation, "TheTargetLocation", "D:\\", "the max number of download film")
	flag.Parse() //这个函数会扫描命令行里的参数，并设置 flag,赋值给上面声明的TheTargetLocation
	download(url1, TheTargetLocation)

}

func download(url, TheTargetLocation string) error {
	//url := "某网站"
	urlinneed, err := http.Get("url")
	if err != nil {
		return err
	}
	defer urlinneed.Body.Close()
	data, err := ioutil.ReadAll(urlinneed.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(TheTargetLocation, data, 0755)

}
