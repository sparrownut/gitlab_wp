package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGetWithCookie(url string) (string, error) {
	//cookies := &http.Session{
	//	Name:  "_gitlab_session",
	//	Value: "9ba4e61db11da8b0160892609c3c77f9",
	//}
	// 创建一个 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("创建 HTTP 请求失败：%v", err)
	}

	// 添加一个名为 "_gitlab_session" 值为 "9ba4e61db11da8b0160892609c3c77f9" 的 cookie
	cookie := &http.Cookie{
		Name:  "_gitlab_session",
		Value: "9ba4e61db11da8b0160892609c3c77f9",
	}
	req.AddCookie(cookie)

	// 使用 http.DefaultClient 发送 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {

		return "", fmt.Errorf("发送 HTTP 请求失败：%v", err)
	}
	defer resp.Body.Close()

	// 读取 HTTP 响应的 body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return "", fmt.Errorf("读取 HTTP 响应失败：%v", err)
	}

	values := ""
	for _, k := range resp.Header.Values("Set-Cookie") {
		values += k
	}
	return values + string(body), nil
}
func HttpGet(url string) (string, error) {
	//cookies := &http.Session{
	//	Name:  "_gitlab_session",
	//	Value: "9ba4e61db11da8b0160892609c3c77f9",
	//}
	// 创建一个 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("创建 HTTP 请求失败：%v", err)
	}

	// 添加一个名为 "_gitlab_session" 值为 "9ba4e61db11da8b0160892609c3c77f9" 的 cookie
	cookie := &http.Cookie{
		Name:  "_gitlab_session",
		Value: "9ba4e61db11da8b0160892609c3c77f9",
	}
	req.AddCookie(cookie)

	// 使用 http.DefaultClient 发送 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {

		return "", fmt.Errorf("发送 HTTP 请求失败：%v", err)
	}
	defer resp.Body.Close()

	// 读取 HTTP 响应的 body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return "", fmt.Errorf("读取 HTTP 响应失败：%v", err)
	}

	return string(body), nil
}

func ExtractStringFromResponse(resp *http.Response) (string, error) {
	// 读取响应的 Body
	body, err := ioutil.ReadAll(resp.Body)
	//println(body)
	if err != nil {
		return "", fmt.Errorf("读取响应 Body 失败：%v", err)
	}

	// 将字节数组转换为字符串
	str := string(body)
	return str, nil
}
