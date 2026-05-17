package main

import (
	"bufio"
	"os"
)

func main() {
	// f, err := os.Open("exmaple.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileinfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(fileinfo.Name(), fileinfo.IsDir(), fileinfo.Size(), fileinfo.Mode(), fileinfo.ModTime())

	// read file
	// buf := make([]byte, 10)
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < d; i++ {
	// 	fmt.Println("Dsta", d, string(buf[i]) )
	// }
	// defer f.Close()

	// Read Method for Small Files
	// data, err := os.ReadFile("exmaple.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// Craete a file
	// f, err := os.Create("example2.txt")
	// defer f.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// // Append
	// f.WriteString("Hii go")
	// f.WriteString("Hii go 	Bye")

	// bytes := []byte("hello Golang")
	// f.Write(bytes)

	// reading data from one file and write into another )stram fashion

	sourceFile, err := os.Open("exmaple.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	deftFile, err := os.Create("hii.txt")
	if err != nil {
		panic(err)
	}
	defer deftFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(deftFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()
}
