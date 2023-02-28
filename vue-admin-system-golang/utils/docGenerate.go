package utils

import (
	"errors"
	"fmt"
	"github.com/nguyenthenguyen/docx"
	"math/rand"
	"time"
)

func DocGenerate(month string) (string, error) {
	r, err := docx.ReadDocxFile("./template.docx")
	if err != nil {
		fmt.Println("err", err)
		return "", errors.New("err")
	}
	docx1 := r.Editable()
	docx1.ReplaceHeader("header", month)
	fmt.Println("word", RandWord())
	filePath := fmt.Sprintf("./monthReport/%s.docx", RandWord())
	fmt.Println("filePath", filePath)
	err = docx1.WriteToFile(filePath)
	if err != nil {
		fmt.Println("err", err)
	}
	return filePath, nil
}

func RandWord() string {
	// 长度为62
	var bytes []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")
	// 保证每次生成的随机数不一样
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 20)
	for i := 0; i < 20; i++ {
		result[i] = bytes[rand.Int31()%62]
	}
	return string(result)
}
