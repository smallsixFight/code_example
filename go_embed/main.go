package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
)

//go:embed test.txt
var str string

//go:embed test.txt
//go:embed test2.txt
var f embed.FS

//go:embed test.txt test2.txt
var ff embed.FS

//go:embed dir
var dirF embed.FS

//go:embed *
var dirF2 embed.FS

func main() {
	fmt.Println(str)
	fmt.Println("=============  FS（fileSystem）  =================")
	f1, err := f.ReadFile("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(f1))
	f1, err = f.ReadFile("test2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(f1))
	fmt.Println("===============  ff 等同于 f  ===============")
	f1, err = ff.ReadFile("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(f1))
	f1, err = ff.ReadFile("test2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(f1))
	fmt.Println("===============  读取目录下的文件  ===============")
	f1, err = dirF.ReadFile("dir/hello.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(f1))
	f1, err = dirF.ReadFile("dir/hello22.txt")
	fmt.Println(string(f1))
	fmt.Println("===============  当前目录下的文件  ===============")
	f1, _ = dirF2.ReadFile("test.txt")
	fmt.Println(string(f1))
}
