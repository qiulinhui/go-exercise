package container

import (
	"fmt"
	"strings"
)

// Strings strings 包的使用
func Strings() {
	/*
		strings 包下的关于字符串的函数
	*/

	s1 := "helloword"
	//1.是否包含指定的内容 bool
	fmt.Println(strings.Contains(s1, "hello"))
	fmt.Println(strings.Contains(s1, "abc"))
	//2.是否包含chars中任意的一个字符即可
	fmt.Println(strings.ContainsAny(s1, "abcd"))
	//3.统计 substr 在 s 中出现的次数
	fmt.Println(strings.Count(s1, "lloo"))
	//4.以xxx前缀开头，以xxx后缀结尾
	s2 := "20190511课堂笔记.txt"
	if strings.HasPrefix(s2, "201905") {
		fmt.Println("2019年5月份的文件")
	}
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("文本文档")
	}

	fmt.Println(strings.Index(s1, "lloo"))     //查找 substr 在 s 中的位置，如果不存在就返回 -1
	fmt.Println(strings.IndexAny(s1, "abcdh")) //查找 chars 中任意的一个字符出现在 s中的位置，如果不存在就返回 -1
	fmt.Println(strings.LastIndex(s1, "l"))    //查找 substr 在 s 中最后一次出现的位置

	//字符串的拼接
	ss1 := []string{"abc", "world", "hello", "ruby"}
	s3 := strings.Join(ss1, "*")
	fmt.Println(s3)

	//字符串切割
	s4 := "123,456,789"
	ss2 := strings.Split(s4, ",")
	fmt.Println(ss2)

	//重复，自己拼接自己 count 次
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5)

	//替换
	s6 := strings.Replace(s1, "l", "*", -1)
	fmt.Println(s6)

	s7 := "hello WorlD**123"
	fmt.Println(strings.ToLower(s7)) //字符串所有大写字母转小写
	fmt.Println(strings.ToUpper(s7)) //字符串所有小写字母转大写

	/*
		截取子字符串：
		str[start:end] -->substr
		包含 start，不包含 end 下标
	*/

	fmt.Println(s1)
	s8 := s1[:5]
	s9 := s1[5:]
	fmt.Println(s8)
	fmt.Println(s9)
}
