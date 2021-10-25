package main

import (
    // "bufio"
    // "fmt"
    // "io"
    // "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // dat, err := ioutil.ReadFile("./file/640.webp")
    // check(err)

		// f, err2 := os.Create("./file/my.webp")
    // check(err2)

		// f.Write(dat)

		f, err2 := os.Create("./file/hello.txt")
    check(err2)

		_, err := f.WriteString("writes\n")
    check(err)
    // fmt.Printf("wrote %d bytes\n", n3)
		
    // err2 := ioutil.WriteFile("./file/my.txt", dat, 0644)
    // check(err2)

    f.Close()
}