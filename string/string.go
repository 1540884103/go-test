package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	testContains()
	testCompare()
	testContainsAny()
	testContainsRune()
	testCount()
	testEqualFold()
	testFields()
	testFieldsFunc()
	testHasPrefix()
	testHasSuffix()
	testIndex()
	testIndexAny()
	testIndexFunc()
	testIndexRune()
	testJoin()
	testLastIndex()
	testMap()
	testRepeat()
	testReplace()
}

func testCompare() {
	fmt.Println("----------testCompare-------------")
	// a > b 返回1
	fmt.Println(strings.Compare("1235", "1234"))
	// a < b 返回-1
	fmt.Println(strings.Compare("123", "1234"))
	fmt.Println(strings.Compare("", "1234"))
	// a == b 返回 0，空串与空串相等，空串<任意字符串
	fmt.Println(strings.Compare("", ""))
}

func testContains() {
	fmt.Println("----------testContains-------------")
	fmt.Println(strings.Contains("asd", "a"))
	//字符串判断包含空串返回true
	fmt.Println(strings.Contains("sad", ""))
	fmt.Println(strings.Contains("asa", "as"))
	//空串判断包含空串返回true
	fmt.Println(strings.Contains("", ""))
}

//a是否包含b中任何字符
func testContainsAny() {
	fmt.Println("----------testContainsAny-------------")
	fmt.Println(strings.ContainsAny("asdasd", "aoioioi"))
	//空串中没有字符，所有会返回false
	fmt.Println(strings.ContainsAny("sad", ""))
	//空串不包含空字符，返回false
	fmt.Println(strings.ContainsAny("", ""))
}

//判断字段串中是否包含指定uniCode，看起来不怎么常用
func testContainsRune() {
	fmt.Println("----------testContainsRune-------------")
	fmt.Println(strings.ContainsRune("asdad", 1))
}

//判断字符串出现次数
func testCount() {
	fmt.Println("----------testCount-------------")
	fmt.Println(strings.Count("123", "12"))
	fmt.Println(strings.Count("12111212", "12"))
	//注意：空串出现次数 = 字符串长度+1
	fmt.Println(strings.Count("123", ""))
}

//比较字符串在Unicode case-folding下是否相等
//注意：不区分大小写
func testEqualFold() {
	fmt.Println("----------testEqualFold-------------")
	fmt.Println(strings.EqualFold("123", "12"))
	fmt.Println(strings.EqualFold("asas", "ASAS"))
	//空串等于空串
	fmt.Println(strings.EqualFold("", ""))
	fmt.Println(strings.EqualFold("sas", ""))
}

//将字符串根据空白字符分割，返回slice
func testFields() {
	fmt.Println("----------testFields-------------")
	fmt.Println(strings.Fields("123    ass  \ns"))
	//纯空格，会返回空数组
	fmt.Println(strings.Fields("          "))
}

//rune 用来处理字符的 int32
//FieldsFunc ，用函数结果分割，func(rune) bool，输入unicode返回true则为分割符号
func testFieldsFunc() {
	fmt.Println("----------testFieldsFunc-------------")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("sad ;';asd;'as;ddsa;..sasd123;12 3;2;3;23 '", f))
}

func testHasPrefix() {
	fmt.Println("----------testHasPrefix-------------")
	fmt.Println(strings.HasPrefix("asdsad", ""))
	fmt.Println(strings.HasPrefix("asdsad", "asas"))
	fmt.Println(strings.HasPrefix("asdsad", "A"))
	//空串判断是否空串前缀，返回true
	fmt.Println(strings.HasPrefix("", ""))
}

func testHasSuffix() {
	fmt.Println("----------testHasSuffix-------------")
	fmt.Println(strings.HasSuffix("asdsad", ""))
	fmt.Println(strings.HasSuffix("asdsad", "ad"))
	fmt.Println(strings.HasSuffix("asdsad", "A"))
	//空串判断是否空串后缀，返回true
	fmt.Println(strings.HasSuffix("", ""))
}

func testIndex() {
	fmt.Println("----------testIndex-------------")
	//空串匹配
	fmt.Println(strings.Index("asdsad", ""))
	//返回出现第一次的数组位置（初始游标为0）
	fmt.Println(strings.Index("asdsad", "ad"))
	//没出现过返回-1
	fmt.Println(strings.Index("asdsad", "A"))
	fmt.Println(strings.Index("", ""))
}

func testIndexAny() {
	fmt.Println("----------testIndexAny-------------")
	//空串不包含字符，故返回-1
	fmt.Println(strings.IndexAny("asdsad", ""))
	//返回给定字符串中任意字符出现第一次的数组位置（初始游标为0），区分大小写
	fmt.Println(strings.IndexAny("asdsad", "dA"))
	//没出现过返回-1
	fmt.Println(strings.IndexAny("asdsad", "A"))
	fmt.Println(strings.IndexAny("", ""))
}

//给定函数返回为true的字符，返回第一次出现的位置，没有则返回-1
func testIndexFunc() {
	fmt.Println("----------testIndexFunc-------------")
	f := func(c rune) bool {
		return unicode.IsNumber(c) || unicode.IsLetter(c)
	}
	fmt.Println(strings.IndexFunc("  sdds  ", f))
	fmt.Println(strings.IndexFunc("123", f))
	fmt.Println(strings.IndexFunc("*^*(*gjhgY&^&", f))
	fmt.Println(strings.IndexFunc("&*^*&^", f))
}

//和indexByte一样
func testIndexRune() {
	fmt.Println("----------testIndexRune-------------")
	fmt.Println(strings.IndexRune("  sdds  ", ' '))
	fmt.Println(strings.IndexByte("  sdds  ", ' '))
	fmt.Println(strings.IndexRune("123", 'd'))
	fmt.Println(strings.IndexByte("123", 'd'))
	fmt.Println(strings.IndexRune("*^*(*gjhgY&^&", '*'))
	fmt.Println(strings.IndexByte("*^*(*gjhgY&^&", '*'))
	fmt.Println(strings.IndexRune("&*^*&^", '^'))
	fmt.Println(strings.IndexByte("&*^*&^", '^'))
}

func testJoin() {
	fmt.Println("----------testJoin-------------")
	s := []string{"1", "2", "3", "  "}
	fmt.Println(strings.Join(s, "sadsd"))
	fmt.Println(strings.Join(s, ""))

}

//lastIndex系列，与index正好相反，最后一次出现的位置
func testLastIndex() {
	fmt.Println("----------testLastIndex-------------")
	//注意空串返回的位置，正的时候是0，反的时候是6不是5，这里这样理解，第五个位置不能匹配
	fmt.Println(strings.Index("asdsad", ""))
	fmt.Println(strings.LastIndex("asdsad", ""))
	//返回出现第一次的数组位置（初始游标为0）
	fmt.Println(strings.LastIndex("asdsad", "ad"))
	//没出现过返回-1
	fmt.Println(strings.LastIndex("asdsad", "A"))
	fmt.Println(strings.LastIndex("", ""))
}

//用函数对字符串对每个字符做映射
func testMap() {
	fmt.Println("----------testMap-------------")
	f := func(c rune) rune {
		if unicode.IsLower(c) {
			return unicode.ToUpper(c)
		} else {
			return c
		}
	}
	fmt.Println(strings.Map(f, "123IUO98798saddsd"))
}

func testRepeat() {
	fmt.Println("----------testRepeat-------------")
	fmt.Println(strings.Repeat("sads", 3))
	//返回空字符串，count为负数，panic
	fmt.Println(strings.Repeat("sads", 0) + "d")
	fmt.Println(strings.Repeat("", 100000000000000000) + "d")
}

//用新字符窜替换旧字符串，替换次数，负数为不限制次数等同于replaceAll
func testReplace() {
	fmt.Println("----------testReplace-------------")
	fmt.Println(strings.Replace("asas,asas,asas,asas", "as", "l", 2))
	fmt.Println(strings.Replace("asas,asas,asas,asas", "as", "l", -1))
	fmt.Println(strings.Replace("asas,asas,asas,asas", "as", "l", -2))
	fmt.Println(strings.Replace("asas,asas,asas,asas", "as", "", -2))
	//空串逻辑上只有一次出现
	fmt.Println(strings.Replace("", "", "l", -2))
}

