package main

import (
	"./gmail"
)

func main() {
	foo := gmail.New("wayne900619@gmail.com", "awbnlgfmcdwyubru")               //自己的邮箱地址和密码
	foo.Send("link test", "http://35.189.167.203:8888/", "39701965a@gmail.com") //邮件标题 邮件内容 需要发送到的邮箱地址
}
