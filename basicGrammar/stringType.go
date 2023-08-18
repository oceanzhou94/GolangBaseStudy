/*字符串类型作为任何语言的重要数据类型重点学习，包括很多操作*/
package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//获取字符串长度len(), RuneCountInString()
	str1 := "hello"
	str2 := "华硕电脑ASUS" //utf-8编码，一个汉字3个字节，所以一个汉字的长度为3
	fmt.Println("str1长度", len(str1))
	fmt.Println("str2长度", len(str2))

	//统计 Uncode 字符数量
	fmt.Println("str1长度", utf8.RuneCountInString(str1))
	fmt.Println("str2长度", utf8.RuneCountInString(str2))

	//遍历字符串
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c->%d  ", str1[i], str1[i])
	}
	fmt.Println()

	//使用range遍历字符串 会正确输出包含中文的字符串
	for _, char := range str2 {
		fmt.Printf("%c->%d  ", char, char)
	}
	fmt.Println()

	//字符串查找string.Index(源字符串,子串),strings.LastIndex()从右往左查找子串
	str3 := "my name is mybob"
	s := "my"
	fmt.Println("find my in str3:", strings.Index(str3, s))
	fmt.Println("find my in last str3:", strings.LastIndex(str3, s))

	//字符串拼接 "+" 简单但是不高效
	// StringBuilder 来进行高效的字符串连接，声明缓冲-写入缓冲-形式转换
	str4 := "hello "
	str5 := "good morning"
	var stringBuilder bytes.Buffer //声明字节缓冲

	//把要拼接的字符串写入缓冲
	stringBuilder.WriteString(str4)
	stringBuilder.WriteString(str5)

	//将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())

}
