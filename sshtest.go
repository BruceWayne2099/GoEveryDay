package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// 创建 ssh 配置
	sshConfig := &ssh.ClientConfig{
	User: "root",
	Auth: []ssh.AuthMethod{
	ssh.Password("password"),
	},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	Timeout:         5 * time.Second,
}

// 创建 client
client, err := ssh.Dial("tcp", "192.168.1.20:22", sshConfig)
checkErr(err)
defer client.Close()

	// 获取 session
	session, err := client.NewSession()
	checkErr(err)
	defer session.Close()

	// 拿到当前终端文件描述符
	fd := int(os.Stdin.Fd())
	termWidth, termHeight, err := terminal.GetSize(fd)

	// request pty
	err = session.RequestPty("xterm-256color", termHeight, termWidth, ssh.TerminalModes{})
	checkErr(err)

	// 对接 std
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	err = session.Shell()
	checkErr(err)
	err = session.Wait()
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
	fmt.Println(err)
	os.Exit(1)
	}
}