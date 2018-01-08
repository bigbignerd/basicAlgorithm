package util

import (
	"io/ioutil"
	"log"
	"strings"
)

//读取文件：将单词存储进[]string
func ReadFile(filepath string) ([]string, error) {
	//分词结果
	result := []string{}
	byteContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Read file error:%v", err)
		return result, err
	}
	//byte content to string
	content := removeCommaAndDot(string(byteContent))
	for _, str := range strings.Split(content, " ") {
		noSpaceStr := strings.TrimSpace(str)
		if strings.TrimSpace(noSpaceStr) == "" {
			continue
		}
		//处理首字母大写的单词
		strLower := strings.ToLower(noSpaceStr)
		result = append(result, strLower)
	}
	return result, nil
}

//去掉字符串文本中的逗号和点
func removeCommaAndDot(s string) string {
	removeCommaStr := strings.Replace(s, ",", "", 0)
	removeDotStr := strings.Replace(removeCommaStr, ".", "", 0)
	return removeDotStr
}
