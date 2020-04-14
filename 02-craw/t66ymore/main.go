package main

import (
	"fmt"
	"io/ioutil"
	"mytool/02-craw/t66y/cmd"
	"mytool/02-craw/t66y/controller"
	"mytool/02-craw/util"
	"os"
	"regexp"
	"strconv"
	"sync"
)

var (
	pageRe = `<a href="(htm_data/2\d{3}?.+)" target=.*id.*>(.+).*</a>`
	//pageRe = `<a href="(htm_data/2\d{3}?.+)" target=.*id="">(.+).*|<font>(.+)</font></a>`
	//pageRe = `<a href="(.+)" target=.*id=""><font.+>(.+)</font>.*`
	wg sync.WaitGroup
)

func working(start, end int) {
	baseURL := "http://t66y.com/thread0806.php?fid=16&page=%d"
	pageChan := make(chan int, 2)
	fmt.Printf("开始下载%d到%d页的数据...\n", start, end)
	for i := start; i <= end; i++ {
		wg.Add(1)
		pageChan <- i
		go SpiderPage(baseURL, i, pageChan)
	}
}

func SpiderPage(baseURL string, i int, pageChan chan int) {
	pagePath := util.GetFilePath("02-craw/images")
	fmt.Println(pagePath)
	url := fmt.Sprintf(baseURL, i)
	fmt.Println(url)
	result, err := util.HttpGet(url)
	if err != nil {
		fmt.Println("httpGet err:", err)
		return
	}
	//saveHtmlToFile(err, pagePath, i, result)
	getURLInfoFromBaseURL(result)
	defer wg.Done()
	<-pageChan
}

func saveHtmlToFile(err error, pagePath string, i int, result string) {
	file, err := os.Create(pagePath + "/第" + strconv.Itoa(i) + "页.html")
	defer file.Close()
	if err != nil {
		fmt.Println("os.Create err", err)
	}
	file.WriteString(result)
}

func main() {
	working(1, 1)
	//Test_refile()
	wg.Wait()
}

func Test_refile() {
	file, err := os.OpenFile("02-craw/images/第1页.html", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	s, _ := ioutil.ReadAll(file)
	content := string(s)
	getURLInfoFromBaseURL(content)
}

func getURLInfoFromBaseURL(content string) {

	re := regexp.MustCompile(pageRe)
	submatch := re.FindAllStringSubmatch(content, -1)
	for _, v := range submatch {
		url := fmt.Sprintf("http://t66y.com/%s", v[1])
		controller.Run(url, cmd.DownLoadPath, 5)
	}
}
