package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]
	fmt.Printf("MD5 (%s) = %s \n", filePath, getFileMd5(filePath))

}

func getFileMd5(str string) string {

	f, err := os.Open(str)
	if err != nil {
		print("[Error] Please check that the file exists")
	}

	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}
	md5str := hex.EncodeToString(h.Sum(nil))

	return md5str
}
