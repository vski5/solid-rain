package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

//go run downloadtest.go -downloadurl 网站名 -TheTargetLocation 下载到哪个位置（默认D:\）
var url1, TheTargetLocation, file2 string

func main() {
	flag.PrintDefaults() //打印 flag 的使用帮助信息
	flag.StringVar(&url1, "downloadurl", "err", "the max number of download film")
	flag.Parse() //这个函数会扫描命令行里的参数，并设置 flag,赋值给上面声明的url1
	flag.StringVar(&TheTargetLocation, "TheTargetLocation", "D:\\", "the max number of download film")
	flag.Parse() //这个函数会扫描命令行里的参数，并设置 flag,赋值给上面声明的TheTargetLocation
	DownloadFile(url1, TheTargetLocation)

}

func DownloadFile(url1, file2 string) {
	//url1 :="http://xxx/somebigfile"
	resp, err := http.Get(url1)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}
	defer resp.Body.Close()

	outfile, err2 := os.OpenFile(file2, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	//out, err := os.Create(TheTargetLocation)
	// 很重要：初始化一个 io.Writer
	wt := bufio.NewWriter(outfile)

	defer outfile.Close()

	jieshoudizhi, err := io.Copy(wt, resp.Body)
	if err != nil {
		panic(err)
	}
	wt.Flush()
	fmt.Fprintln(os.Stdout, "no err", jieshoudizhi)

}
