package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	testSplit()
	testSplitAfter()
	testSplitAfterN()
	testTitle()
	testToLower()
	testToLowerSpecial()
	testToTitle()
	testToValidUTF8()
	testTrim()
	testTrimFunc()
}

//分割，分割符号丢弃
func testSplit() {
	fmt.Println("----------testSplit-------------")
	fmt.Println(strings.Split("2312asdasdsad", "a"))
	//用空串分割，则生成单个字符的字符串slice
	fmt.Println(strings.Split("2312asdasdsad", ""))
	fmt.Println(strings.Split("", ""))
}

//分割，在指定字符串后面分割，不丢弃分割符号
func testSplitAfter() {
	fmt.Println("----------testSplitAfter-------------")
	fmt.Println(strings.SplitAfter("2312asdasdsad", "a"))
	//用空串分割，则生成单个字符的字符串slice
	fmt.Println(strings.SplitAfter("2312asdasdsad", ""))
	fmt.Println(strings.SplitAfter("", ""))
}

//分割，在指定字符串后面分割，不丢弃分割符号，返回分割的数量，0则返回空slice，n>可分割数量，则返回全部分割的结果
//splitN用法相似
func testSplitAfterN() {
	fmt.Println("----------testSplitAfterN-------------")
	fmt.Println(strings.SplitAfterN("2312asdasdsad", "a", 1))
	//用空串分割，则生成单个字符的字符串slice
	fmt.Println(strings.SplitAfterN("2312asdasdsad", "", 2))
	fmt.Println(strings.SplitAfterN("", "", 9))
}

//首字母转为unicode的大写
func testTitle() {
	fmt.Println("----------testTitle-------------")
	fmt.Println(strings.Title("asdsads 123"))
	fmt.Println(strings.Title(""))
	fmt.Println(strings.Title("1212"))
	fmt.Println(strings.Title("*&^*&^"))
	fmt.Println(strings.Title("哈睡觉啊可视电话接口"))
	//每个单词首字母大写
	fmt.Println(strings.Title("sjh jkh jkhkj h jhsdsd   "))
}

//转小写，toUpper转大写，不写了
func testToLower() {
	fmt.Println("----------testToLower-------------")
	fmt.Println(strings.ToLower("JHGHJUYYTUI^&sad萨迪哈睡觉肯定"))

	fmt.Println(strings.ToLower(""))
}

//特殊的语言小写，基本用不上，special的都是酱紫，后面就不写demo了
func testToLowerSpecial() {
	fmt.Println("----------testToLowerSpecial-------------")
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
	fmt.Println(strings.ToLower(""))
}

//全部字母转为unicode的大写
func testToTitle() {
	fmt.Println("----------testToTitle-------------")
	fmt.Println(strings.ToTitle("sadsd sda sd "))
}

//把所有不合法的utf8替换为指定字符串，没找到不合法的case，但感觉这个函数在一些对外输入场景有用，防止一些非法字符进来
func testToValidUTF8() {
	fmt.Println("----------testToValidUTF8-------------")
	fmt.Println(strings.ToValidUTF8("sadsd sda 思考打火机烤红薯的sd ", ""))
}

//指定字符集合进行trim
//trimPrefix：指定前缀trim掉
//trimSuffix：指定后缀trim掉
//trimSpace：空trim掉
func testTrim() {
	fmt.Println("----------testTrim-------------")
	fmt.Println(strings.Trim("sdasdas\n", " \n"))
}

//返回true的字符trim
func testTrimFunc() {
	fmt.Println("----------testTrimFunc-------------")
	f := func(c rune) bool {
		return unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimFunc("asdad2131", f))
}
