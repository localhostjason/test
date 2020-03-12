package main

import (
	"./util"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	srcPath string = "/usr/cms/html"
)

type BadFile struct {
	camFile      string
	webShellFile string
}

func main() {
	var BadFile BadFile
	str, _ := os.Getwd()
	BadFile.camFile = filepath.Join(str, "bad_file", "file.html")
	BadFile.webShellFile = filepath.Join(str, "bad_file", "webshell.jsp")

	// 定时执行
	interval := time.Duration(time.Second * 1)
	ticker_ := time.NewTicker(interval)
	defer ticker_.Stop()

	for {
		<-ticker_.C
		// 定义 源目录
		t := time.Now().Format("2006-01-02")
		filePath := filepath.Join(srcPath, t)
		existSrcPath := util.PathExists(filePath)
		if !existSrcPath {
			_ = os.MkdirAll(filePath, os.ModePerm)
		}

		/*
			发布文件
			上传发布文件概率 {9:30,17:00}[5,10]->50%[1,5]
			修改发布文件概率 {9:00-17:00} [60]->10%[2,5]
			删除文件概率 {9:00-17:00} [60]->5%[1,2]
		*/
		UploadChangeFile(filePath, BadFile)
		UploadUpdateFile(filePath)
		UploadDeleteFile(filePath)

	}
}

func UploadChangeFile(srcPath string, badFile BadFile) {
	/*
		上传文件
	*/
	timeMinute := time.Now().Format("15:04")
	timeSecond := time.Now().Format("15:04:05")

	// {9:00}->[50,100] 上传
	if timeMinute == "09:00" {
		fileNum := util.GenerateRangeNum(50, 100)
		util.CreateFile(srcPath, fileNum)
		//	1% 概率 出现木马文件和 伪装文件
		camPercent := util.SetPercent(1)
		if camPercent {
			UploadBacFile(srcPath, "camouflage", badFile)
		}
		webShellPercent := util.SetPercent(1)
		if webShellPercent {
			UploadBacFile(srcPath, "webShell", badFile)
		}
	}

	// {13:00}->[10,20] 上传
	if timeMinute == "13:00" {
		fileNum := util.GenerateRangeNum(10, 20)
		util.CreateFile(srcPath, fileNum)
		//	1% 概率 出现木马文件和 伪装文件
		camPercent := util.SetPercent(1)
		if camPercent {
			UploadBacFile(srcPath, "camouflage", badFile)
		}
		webShellPercent := util.SetPercent(1)
		if webShellPercent {
			UploadBacFile(srcPath, "webShell", badFile)
		}
	}

	// {9:30,17:00}[5,10]->50%[1,5] 上传
	if timeSecond >= "09:30:00" && timeSecond <= "17:00:00" {
		fmt.Println(timeSecond)
		createPercent := util.SetPercent(50)
		if createPercent {
			fileNum := util.GenerateRangeNum(1, 5)
			util.CreateFile(srcPath, fileNum)
			//	1% 概率 出现木马文件和 伪装文件, 发生 将会 copy 一个文件 到发布目录
			camPercent := util.SetPercent(1)
			if camPercent {
				UploadBacFile(srcPath, "camouflage", badFile)
			}
			webShellPercent := util.SetPercent(1)
			if webShellPercent {
				UploadBacFile(srcPath, "webShell", badFile)
			}
		}
		time.Sleep(time.Duration(util.GenerateRangeNum(5, 10)) * time.Second)
	}
}

func UploadUpdateFile(srcPath string) {
	// {9:00-17:00} [60]->10%[2,5] 修改
	timeMinute := time.Now().Format("15:04")
	if timeMinute >= "09:00" && timeMinute <= "17:00" {
		updatePercent := util.SetPercent(10)
		if updatePercent {
			fileNum := util.GenerateRangeNum(2, 5)
			util.UpdateFiles(srcPath, fileNum)
		}
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func UploadDeleteFile(srcPath string) {
	// {9:00-17:00} [60]->5%[1,2] 删除
	timeMinute := time.Now().Format("15:04")
	if timeMinute >= "09:00" && timeMinute <= "17:00" {
		deletePercent := util.SetPercent(5)
		if deletePercent {
			fileNum := util.GenerateRangeNum(1, 2)
			util.DeleteFiles(srcPath, fileNum)
		}
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func UploadBacFile(srcPath, fileType string, badFile BadFile) {
	ex := "html"
	srcDir := badFile.camFile
	if fileType == "webShell" {
		ex = "jsp"
		srcDir = badFile.webShellFile
	}
	t := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s.%s", t, fileType, ex)
	dstPath := filepath.Join(srcPath, fileName)

	_, _ = CopyFile(dstPath, srcDir)
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	/*
		copy 文件
		dstName: 目标目录
		srcName: 源目录
	*/
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
