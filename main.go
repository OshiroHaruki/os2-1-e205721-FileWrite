package main

import (
	"flag"
	"fmt"
	"os-2-1/fileRead"
	"os-2-1/fileWrite"
)

func sub1() {
	fmt.Println("sub1")
}

func main() {
	//fmt.Println("Hello, World!")
	//fmt.Println(fileWrite.Hello("haru"))
	//u := fileWrite.FileIO{}
	//u = getOpts(u)

	fileRead.FileRead("testData.txt")
	fmt.Println("hello,世界")
}

func getOpts(u fileWrite.FileIO) fileWrite.FileIO {
	fmt.Println("Hello,getOpts")
	flag.BoolVar(&u.Buffering, "b", true, "buffering")
	flag.StringVar(&u.FileName, "filename", "testData.txt", "file name to write")
	flag.IntVar(&u.BufferSize, "buffersize", 2048, "buffer size")   //1024
	flag.IntVar(&u.WriteSize, "writesize", 102400000, "write size") //102400
	flag.Parse()
	return u
}
