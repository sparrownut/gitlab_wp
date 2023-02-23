package utils

import (
	"fmt"
	"regexp"
)

func ExtractValueFromAuthenticityToken(token string) (string, string, error) {
	//println(token)
	// 使用正则表达式匹配 value 属性中的值
	r := regexp.MustCompile(`name="authenticity_token" value="(.*)" /`)
	match := r.FindStringSubmatch(token)
	//println(match[1])
	tokenN := ""
	sessionN := ""
	// 如果匹配到了值，则返回第一个子匹配项（即 * 的值）
	if len(match) > 1 {
		tokenN = match[1]
	}
	cookieR := regexp.MustCompile(`_gitlab_session=(.*); path=/`)
	cookieMatch := cookieR.FindStringSubmatch(token)
	if len(cookieMatch) > 1 {
		sessionN = cookieMatch[1]
		return tokenN, sessionN, nil
	}
	//_gitlab_session=0b49f821f9a449be63a8f382a4611c83;
	// 如果未匹配到任何值，则返回错误
	return "", "", fmt.Errorf("无法从字符串 %q 中提取值", token)
}
