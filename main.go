package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var lineFlag = flag.Int("n", 10, "Number of lines to print")
	var byteFlag = flag.Int("c", 0, "Number of bytes to print")

	flag.Parse()
	files := flag.Args()
	fmt.Println(files)

	if len(files) > 0 {
		showFileName := len(files) > 1
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}

			if *byteFlag > 0 {
				printBytes(*byteFlag, f)
			} else if f != nil {
				printLines(*lineFlag, f, showFileName, file)
			}
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		printInput(*lineFlag)
	}
}

func printBytes(bytes int, f *os.File) {
	reader := bufio.NewReader(f)

	p := make([]byte, bytes)
	reader.Read(p)

	fmt.Print(string(p))
	fmt.Println()

}

func printLines(lines int, f *os.File, show bool, name string) {
	if show {
		fmt.Printf("===> %s <===\n", name)
	}

	reader := bufio.NewScanner(f)

	reader.Scan()

	i := 0

	for reader.Scan() && i < lines {
		fmt.Println(reader.Text())
		i++
	}
	fmt.Println()

}

func printInput(lines int) {
	stdin := bufio.NewReader(os.Stdin)
	in := make([]byte, 100)
	for range lines {
		stdin.Read(in)
		fmt.Print(string(in))

	}
}
