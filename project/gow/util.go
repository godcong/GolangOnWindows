package gow

import (
	"syscall"
)

func toTEXT(str string) *uint16 {
	return syscall.StringToUTF16Ptr(str)
}

/*
func getPropertyOnFile(s string) {
	f, _ := os.Open(s)
	defer f.Close()

	var str string
	r := bufio.NewReaderSize(f, 10)
	for {
		tmp := str
		line, isPrefix, err := r.ReadLine()
		if isPrefix {
			str2 := tmp + string(line)
			str = str2
		} else if len(line) > 0 {
			str2 := tmp + string(line) + "\n"
			str = str2
		}
		if err == io.EOF {
			break
		}

	}
	fmt.Println(str)
}
func main() {
	getPropertyOnFile("test.json")
}
*/
