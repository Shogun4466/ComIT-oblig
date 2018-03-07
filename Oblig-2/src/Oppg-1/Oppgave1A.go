package main

import (

	"fmt"

	"os"

)

func main() {

	// can handle symbolic link, but will no follow the link
	fileInfo, err := os.Lstat("text.txt")

	if err != nil {
		panic(err)
	}

	var modePerm = fileInfo.Mode()
	var bytes int64 = fileInfo.Size()

	fmt.Println("Size: ", bytes,)

	fmt.Println("Name : ", fileInfo.Name())

	fmt.Println("Size : ", fileInfo.Size())

	fmt.Println("Is a directory? : ", fileInfo.IsDir())

	fmt.Println("Is a regular file? : ", fileInfo.Mode().IsRegular())

	fmt.Println("Unix permission bits? : ", fileInfo.Mode().Perm())

	if modePerm.IsDir() {

		fmt.Println("Is a directory")

	} else{

		fmt.Println("Has not UNIX permission bits")

	}
		// --- check if file is a symlink

		if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			fmt.Println("File is a symbolic link")
		}

		fmt.Println("Permission in string : ", fileInfo.Mode().String())

		fmt.Println("What else underneath? : ", fileInfo.Sys())

	}
