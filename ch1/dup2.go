package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建map集合赋值给counts变量
	counts := make(map[string]int)
	// 获取命令行参数1赋值给files
	files := os.Args[1:]
	// 判断是否传递命令行参数
	if len(files) == 0 {
		countLines(os.Stdout, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\t%s\n", os.Args[1], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
