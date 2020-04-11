package controller

import (
	"bufio"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/cizixs/gohttp"
	"github.com/hashicorp/go-uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"mytool/02-craw/util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	wg      sync.WaitGroup
	t66yImg = `<img[\s\S]+?ess\-data='(http[\s\S]+?)'`
	title   = `<h4>([\s\S]+?)</h4>`
	//title   = `([\s\S]+?)</h4>`
)

type T66yImage struct {
	Title string
	Urls  []string
}

func GetImagesFromUrl(url string) (t66yImage T66yImage) {
	resp, err := GetBytesFromURL(url)
	defer resp.Body.Close()
	//eocoding := DetetmineEncoding(resp.Body)
	//reader := transform.NewReader(resp.Body, eocoding.NewDecoder())
	result, err := ioutil.ReadAll(resp.Body)
	result = util.ConvertToByte(string(result), "gbk", "utf-8")

	if err != nil {
		fmt.Println(err)
		return
	}

	retURLS := regexp.MustCompile(t66yImg)
	retTitle := regexp.MustCompile(title)
	allString := retURLS.FindAllStringSubmatch(string(result), -1)
	t66yTitle := retTitle.FindAllStringSubmatch(string(result), -1)

	for _, v := range allString {
		if strings.Contains(v[1], "gif") {
			log.Warn("图片为gif,可能为广告，丢弃", v[1])
			continue
		}
		if strings.Contains(v[1], ".th.") {
			v[1] = strings.Replace(v[1], ".th.", ".", -1)
		}
		t66yImage.Urls = append(t66yImage.Urls, v[1])
	}
	t66yImage.Title = t66yTitle[0][1]
	return t66yImage
}

func DownLoadImage(url string, title string, downLoadPath string) {
	resp, err := GetBytesFromURL(url)
	defer resp.Body.Close()

	imagePath := filepath.Join(downLoadPath, CutTitle(title))
	_, err = os.Stat(imagePath)
	if err != nil {
		fmt.Println("目录不存在，新建目录")
		err := os.MkdirAll(imagePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	filename, _ := uuid.GenerateUUID()
	savePath := filepath.Join(imagePath, fmt.Sprintf("%s.jpg", filename))
	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(savePath, imgBytes, 0644)
	if err != nil {
		log.Error(err)
	}
	log.Infof("下载成功:%s", filename)
}
func DownLoadImageAsync(url string, title string, downLoadPath string, paralChan chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		paralChan <- 123
		DownLoadImage(url, title, downLoadPath)
		<-paralChan
		defer wg.Done()
	}()
}

func GetBytesFromURL(url string) (*gohttp.GoResponse, error) {
	random := browser.Chrome()
	resp, err := gohttp.New().
		Proxy("http://127.0.0.1:1080").
		Timeout(time.Second*15).
		Header("user-agent", random).
		Get(url)
	if err != nil {
		fmt.Println(err)
	}

	return resp, err
}

/**
获得网页编码
**/
func DetetmineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		fmt.Println(err)
	}
	encodeing, _, _ := charset.DetermineEncoding(bytes, "")
	return encodeing
}

func CutTitle(title string) string {
	res := strings.Replace(title, "!", "", -1)
	res = strings.Replace(res, "]", "", -1)
	res = strings.Replace(res, "[", "", -1)
	res = strings.Replace(res, "【", "", -1)
	res = strings.Replace(res, "】", "", -1)
	res = strings.Replace(res, "/", "-", -1)
	res = strings.Replace(res, " ", "-", -1)
	return res
}

func GetImagNameFromURl(url string) {
	fmt.Println(url)
	//http://img599.net/images/2020/04/02/SKYHD-097.th.jpg
	//SKYHD-097.th.jpg
	//	title := `<h4>([\s\S]+?)</h4>`
	//imgret := regexp.MustCompile(`/\w+-[\s\S]+?\.(jpg)`)
	imgret := regexp.MustCompile(`/\w+-?\d+?\.?t?h?\.((jpg)|(jpeg))`)
	submatch := imgret.FindAllStringSubmatch(url, -1)
	fmt.Println(submatch)
}

func Run(url string, downLoadPath string, paral int) {
	paralChan := make(chan int, paral)
	t66yImage := GetImagesFromUrl(url)
	log.Infof("该页面总共有%d张图片\n", len(t66yImage.Urls))
	log.Infof("用%d个线程开始下载\n", paral)
	for _, url := range t66yImage.Urls {
		log.Infof("正在下载%s", url)
		DownLoadImageAsync(url, t66yImage.Title, downLoadPath, paralChan, &wg)
	}
	wg.Wait()
}

func GetImage(url string) {
	reg := `<img[\s\S]+?ess\-data='(http[\s\S]+?)'`
	imgret := regexp.MustCompile(reg)
	submatch := imgret.FindAllStringSubmatch(url, -1)
	for _, res := range submatch {
		fmt.Println(res[1])
	}
}

func GetTitle(url string) {
	//title   = `<title[\s\S]+?`
	//	t66yImg = `<img[\s\S]+?ess\-data='(http[\s\S]+?)'`
	//title   = `<h4>([\s\S]+?)</h4>`

	title = `<h4>([\s\S]+?)</h4>`

	imgret := regexp.MustCompile(title)
	//fmt.Println(url)
	submatch := imgret.FindAllStringSubmatch(url, -1)
	fmt.Println(submatch)
	for _, res := range submatch {
		fmt.Println(res)
	}
}
