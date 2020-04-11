package util

import (
	"fmt"
	"github.com/axgle/mahonia"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
	return filepath.Join(path,"images")
}


//src为要转换的字符串，srcCode为待转换的编码格式，targetCode为要转换的编码格式
func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}
