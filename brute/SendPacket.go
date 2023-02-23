package brute

import (
	"fmt"
	"gitlab_weakpassword/Data"
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

func GetTokenAndSession(Url string) (Data.T_S, error) {
	Url = strings.TrimSuffix(Url, "/")
	body, err := utils.HttpGetWithCookie(Url + "/users/sign_in")

	if err != nil {
		return Data.T_S{}, fmt.Errorf("can't connect gitlab")
	}
	token, session, _ := utils.ExtractValueFromAuthenticityToken(body)
	if err != nil {
		return Data.T_S{}, fmt.Errorf("no token")
	}
	return Data.T_S{
		Token:   token,
		Session: session,
	}, nil
}

func LoginToGitLab(Url, username, password string, TS Data.T_S) (*http.Client, error) {
	// 创建一个HTTP客户端，并启用Cookie支持
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar: jar,
	}
	//print("here")
	//get token
	token := TS.Token
	session := TS.Session
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
		//print("here")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		//print("here")
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	//fmt.Println(strconv.Itoa(resp.StatusCode) + username + ":" + password)
	// 否登录成功
	tmpBody, _ := utils.ExtractStringFromResponse(resp)
	//fmt.Println(tmpBody)

	if !(strings.Contains(tmpBody, "qa-your-groups-link")) {

		//response, err := utils.ExtractStringFromResponse(resp)
		//if err != nil {
		//	return nil, err
		//}
		//fmt.Println(response)
		return nil, fmt.Errorf("login failed with status code %d", resp.StatusCode)
	}
	println("成功")
	// 返回已登录的HTTP客户端
	return client, nil
}
