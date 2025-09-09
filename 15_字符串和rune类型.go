package main

import (
    "fmt"
    "unicode/utf8"
)
// Go语言中的字符串是一个只读的byte类型的切片。 Go语言和标准库特别对待字符串 - 作为以 UTF-8 为编码的文本容器。 在其他语言当中， 字符串由”字符”组成。 在Go语言当中，字符的概念被称为 rune - 它是一个表示 Unicode 编码的整数
func strTest() {

    const s = "สวัสดี"

    fmt.Println("Len:", len(s))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
// 这是手动解码 UTF-8 的方法。
// DecodeRuneInString 会从 s[i:] 的开头解析出一个完整的 rune：
// runeValue：解码出的 Unicode 码点。

// width：这个 rune 占多少字节。
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width

        examineRune(runeValue)
    }
}

func examineRune(r rune) {

    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}