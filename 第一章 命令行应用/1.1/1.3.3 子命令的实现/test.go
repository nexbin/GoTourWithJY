package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse() // 它总是在所有命令行参数注册的最后进行调用，函数功能是解析并绑定命令行参数x`

	var name string

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError) // 该方法会返回带有指定名称和错误处理属性的空命令集给我们去使用,相当于就是创建一个新的命令集了去支持子命令了
		/*
			三种错误级别
			const (
				// 返回错误描述
				ContinueOnError ErrorHandling = iota
				// 调用 os.Exit(2) 退出程序
				ExitOnError
				// 调用 panic 语句抛出错误异常
				PanicOnError
			)
		*/
		goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
		_ = goCmd.Parse(args[1:])

	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
