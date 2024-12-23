package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		// handle the error here
		panic(err)
	}

	fileInfo, err := f.Stat() // get the file stats
	if err != nil {
		// handle the error here
		panic(err)
	}

	// output the file name
	fmt.Println(fileInfo.Name())

	// output the file size
	fmt.Println(fileInfo.Size())

	// output the file permissions
	fmt.Println(fileInfo.Mode())

	// output the last modified time
	fmt.Println(fileInfo.ModTime())

	// output the file is a directory or not
	fmt.Println(fileInfo.IsDir())

	// output the system interface
	fmt.Println(fileInfo.Sys())

	// output the file is a regular file or not
	fmt.Println(fileInfo.Mode().IsRegular())

	// output the file is a directory or not
	fmt.Println(fileInfo.Mode().IsDir())

	// read file
	buf := make([]byte, fileInfo.Size())
	n, err := f.Read(buf)
	if err != nil {
		// handle the error here
		panic(err)
	}
	fmt.Print("File content: ")
	fmt.Println(string(buf[:n]))

	// data, err = os.ReadFile("example.txt")
	// if err != nil {
	// 	// handle the error here
	// 	panic(err)
	// }

	// fmt.Print("File content: ")
	// fmt.Println(string(data))

	// close the file
	defer f.Close()

	dir, err := os.Open(".")
	if err != nil {
		// handle the error here
		panic(err)
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		// handle the error here
		panic(err)
	}

	// list all files in the current directory
	fmt.Println(" Files in the current directory: ")
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// create a new file
	newFile, err := os.Create("example2.txt")
	if err != nil {
		// handle the error here
		panic(err)
	}

	// write some text to the file
	n, err = newFile.WriteString("Hello, World!")
	if err != nil {
		// handle the error here
		panic(err)
	}

	fmt.Printf("Wrote %d bytes to the file\n", n)

	// append some text to the file
	n, err = newFile.WriteString("Hello, Go!")
	if err != nil {
		// handle the error here
		panic(err)
	}

	fmt.Printf("Wrote %d bytes to the file\n", n)

	// close the file
	defer newFile.Close()

	// stream a file to another file
	oldFile, err := os.Open("example.txt")
	if err != nil {
		// handle the error here
		panic(err)
	}

	newFile, err = os.Create("example3.txt")
	if err != nil {
		// handle the error here
		panic(err)
	}

	// copy the content of the old file to the new file
	reader := bufio.NewReader(oldFile)
	writer := bufio.NewWriter(newFile)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		writer.WriteString(line)
	}

	fmt.Printf("Copied %d bytes to the new file\n", n)

	// close the files
	defer oldFile.Close()
	defer newFile.Close()

	// remove a file
	err = os.Remove("example3.txt")
	if err != nil {
		// handle the error here
		panic(err)
	}

}
