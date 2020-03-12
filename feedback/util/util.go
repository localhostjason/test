package util

import (
	"fmt"
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
