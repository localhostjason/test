package main

import (
	"./process"
	"./util"
	"flag"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	webPage      string = "/usr/feedback/html"
	upload       string = "/usr/feedback/upload"
	midWare      string = "/usr/feedback/db"
	confHttp     string = "/etc/httpd/conf/httpd.conf"
	confHttpDir  string = "/etc/httpd/conf"
	confPhp      string = "/etc/php/apache2/php.ini"
	confPhpDir   string = "/etc/php/apache2"
	iLockerShell string = "/usr/local/iguard6/ilocker/bin/ilocker"
)

var d = flag.Bool("d", false, "bool类型参数")

func main() {
	checkPath()
	flag.Parse()
	if *d {
		initData()
		return
	}

	process.InitHtml()

	// 定时执行
	interval := time.Duration(time.Second * 1)
	ticker_ := time.NewTicker(interval)
	defer ticker_.Stop()

	for {
		<-ticker_.C
		// 执行主体： process 正常 vim git 增加、修改、删除文件
		process.ChangeFile()
		// 执行主体： 异常 增加、修改、删除文件
		ChangeFile()
	}
}

func initData() {
	/*
		初始化一些假数据
	*/
	cmdStop := exec.Command(iLockerShell, "stop")
	_ = cmdStop.Wait()

	util.CreateFile(upload, 30)
	util.CreateFile(midWare, 5)

	cmdStart := exec.Command(iLockerShell, "start")
	_ = cmdStart.Wait()
}

func ChangeFile() {
	/*
		异常 修改 增加 删除 文件
		进程 都是 go ？
	*/
	webPageCreatePercent := SetPercent(1)
	if webPageCreatePercent {
		util.CreateFile(webPage, 1)
		time.Sleep(time.Duration(GenerateRangeNum(10, 50)) * time.Second)
	}
	webPageUpdatePercent := SetPercent(1)
	if webPageUpdatePercent {
		util.UpdateFiles(webPage, 1)
		time.Sleep(time.Duration(GenerateRangeNum(10, 50)) * time.Second)
	}
	webPageDeletePercent := SetPercent(1)
	if webPageDeletePercent {
		util.DeleteFiles(webPage, 1)
		time.Sleep(time.Duration(GenerateRangeNum(10, 50)) * time.Second)
	}

	uploadCreatePercent := SetPercent(50)
	if uploadCreatePercent {
		util.CreateFile(upload, GenerateRangeNum(1, 2))
		time.Sleep(time.Duration(GenerateRangeNum(5, 20)) * time.Second)
	}

	dbUpdatePercent := SetPercent(50)
	if dbUpdatePercent {
		util.UpdateFiles(midWare, GenerateRangeNum(1, 2))
		time.Sleep(time.Duration(GenerateRangeNum(5, 20)) * time.Second)
	}

	confHttpPercent := SetPercent(1)
	if confHttpPercent {
		t := time.Now().Unix()
		srtT := strconv.FormatInt(t, 10)
		f, _ := os.OpenFile(confHttp, os.O_WRONLY|os.O_TRUNC, 0600)
		_, _ = f.Write([]byte(srtT))
		defer f.Close()
	}

	confPhpPercent := SetPercent(1)
	if confPhpPercent {
		t := time.Now().Unix()
		srtT := strconv.FormatInt(t, 10)
		f, _ := os.OpenFile(confPhp, os.O_WRONLY|os.O_TRUNC, 0600)
		_, _ = f.Write([]byte(srtT))
		defer f.Close()
	}

}

func checkPath() {
	/*
		目录 判断
	*/
	cmdStop := exec.Command(iLockerShell, "stop")
	_ = cmdStop.Wait()

	existWebPage := PathExists(webPage)
	if !existWebPage {
		_ = os.MkdirAll(webPage, os.ModePerm)
	}

	existUpload := PathExists(upload)
	if !existUpload {
		_ = os.MkdirAll(upload, os.ModePerm)
	}

	existMidWare := PathExists(midWare)
	if !existMidWare {
		_ = os.MkdirAll(midWare, os.ModePerm)
	}

	existConfHttpDir := PathExists(confHttpDir)
	existConfHttp := PathExists(confHttp)
	if !existConfHttpDir {
		_ = os.MkdirAll(confHttpDir, os.ModePerm)
	}
	if !existConfHttp {
		f, _ := os.Create(confHttp)
		defer f.Close()
	}

	existConfPhpDir := PathExists(confPhpDir)
	existConfPhp := PathExists(confPhp)
	if !existConfPhpDir {
		_ = os.MkdirAll(confPhpDir, os.ModePerm)
	}
	if !existConfPhp {
		f, _ := os.Create(confPhp)
		defer f.Close()
	}

	cmdStart := exec.Command(iLockerShell, "start")
	_ = cmdStart.Wait()
}

func PathExists(path string) bool {
	/*
		检查 文件路径是否存在
	*/
	_, err := os.Stat(path)

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
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
