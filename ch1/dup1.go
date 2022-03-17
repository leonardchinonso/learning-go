// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

package ch1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Dup1Function() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func TestVar() {
	hexString := "0x75bcd150"
	fmt.Println(hex2int(hexString))

	fmt.Println()
	hexString1 := "0x75bcd150"
	fmt.Println(hex2int1(hexString1))

	//hexString = "0x1"
	//fmt.Println(hex2int(hexString))
}

func hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)
	fmt.Println(cleaned)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return result
}

func hex2int1(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.ReplaceAll(hexStr, "0", "")
	fmt.Println(cleaned)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return result
}