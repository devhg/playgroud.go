package continuecp

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func ContinueCopy() {
	srcFile := "a.txt"
	destFile := "b.txt"
	tempFile := destFile + "_temp.txt"

	src, _ := os.Open(srcFile)
	dest, _ := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	temp, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer dest.Close()
	defer src.Close()

	_, _ = temp.Seek(0, io.SeekStart)
	bytes := make([]byte, 1)

	n1, err := temp.Read(bytes)
	if err != nil {
		panic(err)
	}

	countStr := string(bytes[:n1])
	// fmt.Println(countStr)

	count, _ := strconv.ParseInt(countStr, 10, 64)
	// fmt.Println(count)

	_, _ = src.Seek(count, io.SeekStart)
	_, _ = dest.Seek(count, io.SeekStart)

	data := make([]byte, 1)
	n2, n3 := -1, -1
	total := int(count)

	for i := 0; i < 2; i++ {
		n2, err = src.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("end")
				temp.Close()
				break
			}
		}
		n3, _ = dest.Write(data[:n2])
		total += n3
		_, _ = temp.Seek(0, io.SeekStart)
		_, _ = temp.WriteString(strconv.Itoa(total))

		// break // 中断
	}
}
