package main
import (
	"os"
)
func main() {
	file, err := os.Create("test.txt")
	file1, err := os.Create("test1.txt")
	file2, err := os.Create("test2.txt")


	if err != nil {
		// handle the error here
		return

	}

    defer file.Close()
	defer file1.Close()
	defer file2.Close()

	file.WriteString("test")
	file1.WriteString("test")
	file2.WriteString("test")


}
