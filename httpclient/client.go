package httpclient

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//test
//	resp, err := httpclient.HttpGet("http://121.42.35.23:8003/")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(resp)

//	resp, err := httpclient.HttpPost(TEST_POST_URL, "")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(resp)

//简单封装HTTP POST请求值
//输入url和Body
//输出响应Body
func HttpPost(urlStr string, reqBody string) (respBody string, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	return string(body), nil
}

//简单封装HTTP POST请求值
//输入url
//输出Body
func HttpGet(urlStr string) (respBody string, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}
	return string(body), nil
}
