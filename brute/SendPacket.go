package brute

import (
	"fmt"
	"gitlab_weakpassword/utils"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type loginResponse struct {
	Token string `json:"token"`
}

func getTokenAndSession(Url string) (string, string, error) {
	Url = strings.TrimSuffix(Url, "/")
	body, err := utils.HttpGetWithCookie(Url + "/users/sign_in")

	if err != nil {
		return "", "", fmt.Errorf("can't connect gitlab")
	}
	token, session, _ := utils.ExtractValueFromAuthenticityToken(body)
	if err != nil {
		return "", "", fmt.Errorf("no token")
	}
	return token, session, nil
}

func LoginToGitLab(Url, username, password string) (*http.Client, error) {
	// 创建一个HTTP客户端，并启用Cookie支持
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar: jar,
	}

	//get token
	token, session, err := getTokenAndSession(Url)
	//println(session)
	if err != nil {
		return nil, err
	}
	// 创建登录表单
	data := url.Values{}
	data.Set("user[login]", username)
	data.Set("user[password]", password)
	data.Set("utf8", "✓")
	data.Set("authenticity_token", token) // token

	// 发送POST请求，登录GitLab
	loginUrl := fmt.Sprintf("%s/users/sign_in", Url)
	req, err := http.NewRequest("POST", loginUrl, strings.NewReader(data.Encode()))
	cookies := &http.Cookie{
		Name:  "_gitlab_session",
		Value: session,
	}
	req.AddCookie(cookies)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 检查是否登录成功
	if resp.StatusCode != http.StatusFound {
		//response, err := utils.ExtractStringFromResponse(resp)
		//if err != nil {
		//	return nil, err
		//}
		//fmt.Println(response)
		return nil, fmt.Errorf("login failed with status code %d", resp.StatusCode)
	}

	// 返回已登录的HTTP客户端
	return client, nil
}
