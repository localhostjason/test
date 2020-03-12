package util

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func CreateFile(distPath string, num int) []string {
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

func UpdateFiles(distPath string, num int) []string {
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

func DeleteFiles(distPath string, num int) []string {
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
