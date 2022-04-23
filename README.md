# solid-rain
Iterative testing of downloaded software

用于下载的软件的迭代测试。

# solid-rainV0.1

使用``http.Get()`` 和 ``ioutil.WriteFile()``。
`http.Get()`下载文件到内存，
`ioutil.WriteFile()`将内存里的下载内容直接写到文件中。
很占用内存空间，也没有流式传输的一点点思维。

>go run downloadtest.go -downloadurl 网站名 -TheTargetLocation 下载到哪个位置（默认D:\file）(记得带上保存为什么名字)

# solid-rainV0.2

使用`http.Get()`和`io.copy`。
`http.Get()`下载文件到内存，
`io.copy`利用stdin 和 stdout 设备作为通道，在进程之间传递数据。

以流的方式高效处理数据，而不用考虑数据是什么，数据来自哪里，以及数据要发送到哪里的问题。

很多次转换打开（或是创造）的文件的格式，或者说是文件句柄的类型。

`resp, err := http.Get(url1)`
下载的文件句柄`resp`类型从`*http.Response`
到`resp.Body`也就是`io.Reader`


接收的文件句柄类型从`outfile, err2 := os.OpenFile(file2, os.O_RDWR|os.O_CREATE, os.ModePerm)`中的`outfile`，文件类型为`*os.File`
到`wt := bufio.NewWriter(outfile)`，文件句柄名为`wt`，类型是`*bufio.Writer`，这就是在转换类型的同时分配缓冲空间。
满足`Writer`类型了，现在是可以用在`io.Copy`函数输出值上。

>go run downloadtest.go -downloadurl 网站名 -TheTargetLocation 下载到哪个位置（默认D:\file）(记得带上保存为什么名字)

# solid-rainV0.3
加入bt下载功能。
