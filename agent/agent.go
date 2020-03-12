package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

const (
	DistPath string = "/var/portal/html"
	httpUrl  string = "http://127.0.0.1"
)

func main() {
	// 定时执行
	interval := time.Duration(time.Second * 1)
	ticker_ := time.NewTicker(interval)
	defer ticker_.Stop()

	for {
		<-ticker_.C
		// 定义 目标目录
		t := time.Now().Format("2006-01-02")
		distPath := filepath.Join(DistPath, t)
		exist := PathExists(distPath)
		if !exist {
			_ = os.MkdirAll(distPath, os.ModePerm)
		}

		// 执行主体： 异常 增加、修改、删除文件
		ChangeFile(distPath)
	}
}

func ChangeFile(distPath string) {
	/*
		异常修改 增加 删除文件
	*/
	createPercent := SetPercent(50)
	if createPercent {
		files := _CreateFile(distPath, GenerateRangeNum(1, 5))
		time.Sleep(time.Duration(GenerateRangeNum(1, 3)) * time.Second)

		for i := range files {
			// 被web 访问
			t := time.Now().Format("2006-01-02")
			cmd := exec.Command("curl", fmt.Sprintf("%s/%s/%s", httpUrl, t, files[i]))
			_ = cmd.Wait()
		}
		time.Sleep(time.Duration(GenerateRangeNum(5, 10)) * time.Second)
	}

	updatePercent := SetPercent(50)
	if updatePercent {
		files := _UpdateFiles(distPath, GenerateRangeNum(1, 5))
		time.Sleep(time.Duration(GenerateRangeNum(2, 5)) * time.Second)

		for i := range files {
			// 被web 访问
			t := time.Now().Format("2006-01-02")
			cmd := exec.Command("curl", fmt.Sprintf("%s/%s/%s", httpUrl, t, files[i]))
			_ = cmd.Wait()
		}
		time.Sleep(time.Duration(GenerateRangeNum(5, 10)) * time.Second)
	}

	deletePercent := SetPercent(50)
	if deletePercent {
		files := _DeleteFiles(distPath, GenerateRangeNum(1, 5))
		time.Sleep(time.Duration(GenerateRangeNum(2, 5)) * time.Second)

		for i := range files {
			// 被web 访问
			t := time.Now().Format("2006-01-02")
			cmd := exec.Command("curl", fmt.Sprintf("%s/%s/%s", httpUrl, t, files[i]))
			_ = cmd.Wait()
		}
		time.Sleep(time.Duration(GenerateRangeNum(5, 10)) * time.Second)
	}
}

func _CreateFile(distPath string, num int) []string {
	/*
		创建文件
	*/
	t := time.Now().Unix()
	srtT := strconv.FormatInt(t, 10)

	var result []string
	for i := 0; i < num; i++ {
		fileName := fmt.Sprintf("%d_%d.html", t, i)
		file := filepath.Join(distPath, fileName)
		f, err := os.Create(file)
		if err != nil || f == nil {
			continue
		}
		_, _ = f.Write([]byte(srtT))
		f.Close()
		result = append(result, fileName)
	}
	fmt.Println(result)
	return result
}

func _UpdateFiles(distPath string, num int) []string {
	/*
		更新文件
	*/
	t := time.Now().Unix()
	srtT := strconv.FormatInt(t, 10)

	files, _ := filepath.Glob(filepath.Join(distPath, "*"))
	var result []string

	for i := range files {
		if i < num {
			f, err := os.OpenFile(files[i], os.O_WRONLY|os.O_TRUNC, 0600)
			if err != nil || f == nil {
				continue
			}
			_, _ = f.Write([]byte(srtT))
			f.Close()
		}
	}
	return result
}

func _DeleteFiles(distPath string, num int) []string {
	/*
		删除文件
	*/
	files, _ := filepath.Glob(filepath.Join(distPath, "*"))
	var result []string
	for i := range files {
		if i < num {
			_ = os.Remove(files[i])
		}
	}
	return result
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
