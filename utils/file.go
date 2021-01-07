package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 读取文件内容按行生成数组
func ReadFileToArr(file_path string) []string {
	array := make([]string, 0)

	fi, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return array
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		array = append(array, string(a))
	}
	return array
}

// 把一个数组写文件， 每个成员占一行
func WriteArrToFile(arr []string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, value := range arr {
		fmt.Fprintln(w, value)
	}
	return w.Flush()
}

// 读取本地文件或者远程文件
func FileGetContents(filePth string) ([]byte, error) {
	if strings.Index(filePth, "http") == -1 {
		f, err := os.Open(filePth)
		if err != nil {
			return nil, err
		}
		return ioutil.ReadAll(f)
	} else {
		res, err := http.Get(filePth)
		if err != nil {
			return nil, err
		}
		contents, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		return contents, err
	}
}
