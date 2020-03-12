package process

import (
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	GithubUrl    string = "https://github.com/ishitam8/aksup-python-example.git"
	webPage      string = "/usr/feedback/html"
	GitShell     string = "/usr/bin/git"
	iLockerShell string = "/usr/local/iguard6/ilocker/bin/ilocker"
)

func ChangeFile() {
	/*
		正常进程 修改 增加 删除文件
	*/
	vimPercent := SetPercent(1)
	if vimPercent {
		vimChangeFile()
		time.Sleep(time.Duration(GenerateRangeNum(30, 50)) * time.Second)
	}

	gitPercent := SetPercent(3)
	if gitPercent {
		gitChangeFile()
		time.Sleep(time.Duration(GenerateRangeNum(30, 50)) * time.Second)
	}
}

func InitHtml() {
	/*
		git 初始化 文件
	*/
	cmdStop := exec.Command(iLockerShell, "stop")
	_ = cmdStop.Wait()

	cmdGit := exec.Command(GitShell, "clone", GithubUrl, webPage)
	_ = cmdGit.Wait()

	cmdStart := exec.Command(iLockerShell, "start")
	_ = cmdStart.Wait()
}

func gitChangeFile() {
	/*
		git 进程正常修改 web page 文件
	*/
	str, _ := os.Getwd()
	processFile := filepath.Join(str, "process", "process.sh")

	cmd := exec.Command("sh", processFile)
	_ = cmd.Wait()
}

func vimChangeFile() {
	/*
		vim 进程正常修改 httpd.conf 文件
	*/
	str, _ := os.Getwd()
	processFile := filepath.Join(str, "process", "vim.sh")

	cmd := exec.Command("sh", processFile)
	_ = cmd.Wait()
}

func GenerateRangeNum(min, max int) int {
	/*
		生成 2个数字之间的随机数
	*/
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func SetPercent(num int) bool {
	/*
		使用随机数 计算概率
	*/
	rand.Seed(time.Now().Unix())
	number := GenerateRangeNum(1, 100)
	if number <= num {
		return true
	}
	return false
}
