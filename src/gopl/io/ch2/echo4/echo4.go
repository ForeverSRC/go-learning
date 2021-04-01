package main

import (
	"flag"
	"fmt"
	"strings"
)

// -n 用于忽略行尾的换行符
var n = flag.Bool("n", false, "omit trailing newline") //指针

// -s sep 用于指定分隔符 默认为空格
var sep = flag.String("s", " ", "separator") //指针

// 使用示例
// go build后
// ./echo a bc def
// ./echo -s / a bc def
// .echo -n a bc def
func main() {
	flag.Parse() // 更新标志参数对应的值
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
