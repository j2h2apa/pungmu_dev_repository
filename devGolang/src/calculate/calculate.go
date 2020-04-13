package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("숫자를 입력하세요.")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요.")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	/* case 1 */
	// if line == "+" {
	// 	fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	// } else if line == "-" {
	// 	fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	// } else if line == "*" {
	// 	fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	// } else if line == "/" {
	// 	if n2 == 0 {
	// 		fmt.Println("0으로 나눌 수 없습니다.")
	// 	} else {
	// 		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	// 	}
	// } else {
	// 	fmt.Printf("연산이 가능한 기호가 아니에요.")
	// }

	/* case 2 switch */

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		if n2 == 0 {
			fmt.Println("0으로 나눌 수 없습니다.")
		} else {
			fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
		}
	default:
		fmt.Println("잘못 입력하셨습니다.")
	}
}
