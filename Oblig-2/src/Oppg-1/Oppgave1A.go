package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func main (){
	File()
	fmt.Println ("Information about file text.txt")
	Size()
	Directory ()
	Regular ()
	Unix ()
	Append ()
	Device ()
	Symbolic ()
}

func File ()string{
	arg := os.Args[1:]
	var file string

	for _, v := range arg {
		file += v
	}
	return file
}

func Size() {
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Panicf("failed reading file: %s", err)
	}
	fmt.Printf("Size: %d bytes\n", len(data))
	i := len(data)
	f := float64(i)

	fmt.Printf("Size in KB: %f\n", f/1024)
	fmt.Printf("Size in MB: %f\n", f/1024e3)
	fmt.Printf("Size in GB: %e\n", f/1024e6)
}

func Directory () {
	fi, err := os.Lstat("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode ()

	if mode.IsDir () {
		fmt.Println("Is a directory")
	} else {
		fmt.Println("Is not a directory")
	}
}

func Regular () {
	fi, err := os.Lstat("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode ()

	if mode.IsRegular () {
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}
}

func Unix () {
	fi, err := os.Lstat("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode ()

	var unix = mode & os.ModePerm

	if unix == 0777 {
		fmt.Println("Has Unix permission bits -rwxrwxrwx")
	} else {
		fmt.Println("Does not have Unix permission bits -rwxrwxrwx")
	}
}

func Append () {
	fi, err := os.Lstat("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode ()

	var isappend = mode & os.ModeAppend

	if isappend != 0 {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}
}

func Device () {
	fi, err := os.Lstat("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode ()

	var device = mode & os.ModeDevice

	if device != 0 {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is not a device file")
	}
}

func Symbolic () {
	fi, err := os.Lstat("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode ()

	var symbolic = mode & os.ModeSymlink

	if symbolic != 0 {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}
}
