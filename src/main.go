package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	md, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Errorf("agebox.UpdateMarkdown: Error reade a file: %w", err)
	}

	start := []byte("<!-- age-box start -->")
	before := md[:bytes.Index(md, start)+len(start)]
	end := []byte("<!-- age-box end -->")
	after := md[bytes.Index(md, end):]

	newMd := bytes.NewBuffer(nil)
	newMd.Write(before)
	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Errorf("agebox.UpdateMarkdown: error read time: %w", err)
	}
	newMd.Write([]byte(strconv.Itoa(time.Now().Local().Year() - year)))

	newMd.Write(after)
	fmt.Println("================================")
	fmt.Println("agebox.UpdateMarkdown: Successfully updated")
}
