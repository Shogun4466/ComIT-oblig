package main

import (
	"fmt"
	"log"
	"os"
)

func main (){
	FileInfo(file())
}


func file ()string{
	arg := os.Args[1:]
	var file string

	for _, v := range arg {
		file += v
	}
	return file
}



func FileInfo (file string) {

	fi, err := os.Lstat(file)
	if err != nil {
		log.Fatal(err)
	}
	info, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode()
	fmt.Println ("Information about file",file)

	var bytes int64 = info.Size()
	var kb float64 = (float64)(bytes/1024)
	var mb float64 = (kb/1024)
	var gb float64 = (mb/1024)

fmt.Println("Size:",bytes,"Bytes,",kb,"KB,",mb,"MB,",gb,"GB")

	if mode.IsDir () {
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}

	if mode.IsRegular () {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}

	var isappend = mode & os.ModeAppend
	if isappend != 0 {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}

	var unix = mode & os.ModePerm
	if unix == 0777 {
	fmt.Println("Has Unix permission bits -rwxrwxrwx")
	} else {
	fmt.Println("Does not have Unix permission bits -rwxrwxrwx")
	}

	if mode.IsRegular () {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}

	var device = mode & os.ModeDevice
	if device != 0 {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is not a device file")
	}

	var symbolic = mode & os.ModeSymlink
	if symbolic != 0 {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}
}
