package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("You entered: %s", input)
}
// go doc bufio.Reader.ReadString 軽量ドキュメント確認


