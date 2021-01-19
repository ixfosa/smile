package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

/********************爬虫步骤*******************/
/*
爬虫步骤
	明确目标（确定在哪个网站搜索）
	爬（爬下内容）
	取（筛选想要的）
	处理数据（按照你的想法去处理）
 */

//https://tieba.baidu.com/p/6051076813?pn=1
//https://tieba.baidu.com/p/6051076813?pn=2
//https://tieba.baidu.com/p/6051076813?red_tag=1573533731

//这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	reQQEmail = `(\d+)@qq.com`
)

// 爬邮箱
func GetEmail(url string, pn int)  {
	// 1.去网站拿数据
	resp, err := http.Get(url + "?pn=" + strconv.Itoa(pn))
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll err")
	// 字节转字符串
	pageStr := string(pageBytes)
	//fmt.Println(pageStr)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)

	//fmt.Println(results)
	// 遍历结果
	for _, res := range results {
		fmt.Println("email: ", res[0])
		fmt.Println("qq: ", res[1])
	}


}

func HandleError(err error, why string)  {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main121() {
	for i := 1; i <= 3; i++ {
		go GetEmail("https://tieba.baidu.com/p/6051076813", i)
	}
	for {
		;
	}
}




/***************************************/
var (
	// w代表大小写字母+数字+下划线
	// s?有或者没有s
	// +代表出1次或多次
	//\s\S各种字符
	// +?代表贪婪模式
	reEmail = `\w+@\w+\.\w+`

	reLink = `href="(https?://[\s\S]+?)"`

	rePhone = `1[3456789]\d\s?\d{4}\s?\d{4}`

	reIdCard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`

	reImg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func GetPageStr(url string) (page string) {
	resp, err := http.Get(url)
	HandleError(err,"http.Get ")
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll:")

	pageStr := string(pageBytes)

	return pageStr
}
func GetEmail2(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)

	results := re.FindAllStringSubmatch(pageStr, -1)

	for _, res := range results {
		fmt.Println(res)
	}
}

func GetLink(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reLink)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, res := range results {
		fmt.Println(res[1])
	}
}

func GetPhone(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, res := range results {
		fmt.Println(res)
	}
}

func GetIdCard(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reIdCard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, res := range results {
		fmt.Println(res)
	}
}
func GetImg(url string)  {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, res := range results {
		fmt.Println(res[0])
	}
}

func main122()  {
	//抽取的爬邮箱
	//GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")

	//爬链接
	//GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")

	//爬手机号
	//GetPhone("https://www.zhaohaowang.com/")

	//爬身份证号
	//GetIdCard("https://henan.qq.com/a/20171107/069413.htm")

	//爬图片
	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")


}



/*******************并发爬取美图********************/
// 下载图片，传入的是图片叫什么
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	HandleError(err, "ioutil.ReadAll")

	filename = "E:/imgs/" + filename

	err = ioutil.WriteFile(filename, bytes, 0666)

	if err != nil {
		return false
	}else {
		return true
	}
}

// 并发爬思路：
// 1.初始化数据管道
// 2.爬虫写出：26个协程向管道中添加图片链接
// 3.任务统计协程：检查26个任务是否都完成，完成则关闭数据管道
// 4.下载协程：从管道里读取链接并下载

var (
	chanImageUrls chan string // 存放图片链接的数据管道
	waitGroup sync.WaitGroup
	chanTask chan string // 用于监控协程
)

// 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	// 时间戳解决重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	return
}

// 下载图片
func DownLoadImg()  {
	for url := range chanImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	waitGroup.Done()
}

// 任务统计协程
func CheckOK()  {
	var count int
	for {
		url := <- chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

func GetImgs(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))
	for _, res := range results {
		url := res[0]
		urls = append(urls, url)
	}
	return
}

// 爬图片链接到管道
// url是传的整页链接
func getImgUrls(url string)  {
	urls := GetImgs(url)
	// 遍历切片里所有链接，存入数据管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	waitGroup.Done()
}

func main()  {
	// 1.初始化管道
	chanImageUrls = make(chan string, 26)
	// 2.爬虫协程
	for i := 1; i < 27; i++ {
		waitGroup.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}
	// 3.任务统计协程，统计26个任务是否都完成，完成则关闭管道
	waitGroup.Add(1)
	go CheckOK()
	// 4.下载协程：从管道中读取链接并下载
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go DownLoadImg()
	}
	waitGroup.Wait()
}