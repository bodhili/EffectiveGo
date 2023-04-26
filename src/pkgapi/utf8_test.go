package pkgapi

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestRune(t *testing.T) {
	s := "hello, 世界"
	for _, r := range s {
		t.Logf("%c", r)
	}
}

func TestUtf8Cut(t *testing.T) {
	str := "Hello, 世界"
	fmt.Printf("字符串 \"%s\" 的长度为 %d 个字节\n", str, len(str))

	// 遍历字符串，输出每个字符的Unicode码点和UTF-8编码
	for i, r := range str {
		fmt.Printf("字符 #%d: %q  Unicode码点: %d  UTF-8编码: %x\n", i, r, r, []byte(string(r)))
	}

	// 计算UTF-8编码的字符数
	charCount := utf8.RuneCountInString(str)
	fmt.Printf("字符串 \"%s\" 的UTF-8编码字符数为 %d\n", str, charCount)

	// 截取字符串的一部分，并保证截断位置在一个UTF-8字符边界上
	subStr := str[:utf8.RuneLen(rune(str[7]))+7]
	fmt.Printf("截取字符串 \"%s\" 的一部分: \"%s\"\n", str, subStr)
}
