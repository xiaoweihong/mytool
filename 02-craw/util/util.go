package util

import (
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/axgle/mahonia"
	"github.com/cizixs/gohttp"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func Proxy(proxyStr string) *http.Client {

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyStr)
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}
	return httpClient
}

// 获取下载图片的路径
func GetDownLoadPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return filepath.Join(path, "images")
}

//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func GetFilePath(childPath string) (imagePath string) {
	basePath, _ := os.Getwd()
	imagePath = filepath.Join(basePath, childPath)
	_, err := os.Stat(imagePath)
	if err != nil {
		fmt.Println("目录不存在，新建目录")
		err := os.MkdirAll(imagePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		return imagePath
	}
	return imagePath
}

func HttpGet(url string) (result string, err error) {
	browser := browser.Random()
	resp, err := gohttp.New().
		Timeout(time.Second*20).
		//Proxy("http://127.0.0.1:1080").
		Header("user-agent", browser).
		Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err != nil && err != io.EOF {
			return "", nil
		}
		//累加循环读到的buf数据，存入result，一次性返回
		result += string(buf[:n])
	}
	result = string(ConvertToByte(result, "gbk", "utf-8"))

	return result, nil
}
