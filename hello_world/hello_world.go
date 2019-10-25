// hello_world.go
package main // mast be main here

import (
	"fmt" // Package implementing formatted I/O.
	"os"
	"runtime"
)

func myPrint(msg string) {
	print(msg)
}

func main() {
	println("Hello", "world")
	fmt.Println("hello, world")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")

	msg := "here"
	println(msg)

	var name = "there"
	println(name)

	num := 5 // 赋值操作符，简短声明语法 :=
	println(num)
	println(num + 5)

	numFloat := 5.0
	println(int(numFloat))

	myPrint("what")

	var ( // 这种因式分解关键字的写法一般用于声明全局变量。
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)

	println(HOME, USER, GOROOT)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

}
