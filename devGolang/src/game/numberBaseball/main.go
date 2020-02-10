package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BallCount struct {
	strikes int
	balls   int
}

/* 0~9 사이 겹치지 않는 무작위 숫자 3개 반환 */
func makeNumbers() [3]int {
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					// 겹치면 re loop
					duplicated = true
					break
				}
			}

			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	// fmt.Println(rst)
	return rst
}

/* 키보드로부터 0~9 사이 겹치지 않는 숫자 3개를 입력받아 반환 */
func inputNumbers() [3]int {
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자를 입력하세요.")
		var no int
		// 키입력 버퍼에 엔터도 들어가기 때문에 \n 까지 읽는다
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					// 중복된 숫자 인지를 체크
					duplicated = true
					break
				}
			}

			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			// 숫자의 갯수 체크
			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}

		// 숫자의 갯수가 3보다 작은지 확인
		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
			continue
		}

		if !success {
			continue
		}

		break
	}

	rst[0], rst[2] = rst[2], rst[0]
	fmt.Println(rst)
	return rst
}

/* 2개의 숫자 3개르 비교하여 bool 반환 */
func compareNumber(numbers, inputNumbers [3]int) BallCount {
	// 두개의 숫자 3개를 비교하여 결과를 반환한다.
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	return BallCount{strikes, balls}
}

func printResult(result BallCount) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}

func isGameEnd(result BallCount) bool {
	return result.strikes == 3
}

// 무작위 함수
func main() {
	// rand 가 항상 변하는 값을 갖기 위하여 seed 에 변하는 값을 셋팅해야 한다.
	// 시간은 항상 변하기 때문에 시간 값으로 지정한다.
	rand.Seed(time.Now().UnixNano())
	numbers := makeNumbers()

	cnt := 0

	for {
		cnt++
		// 사용자 입력
		inputNumbers := inputNumbers()
		// 결과 비교
		var result BallCount = compareNumber(numbers, inputNumbers)

		// 출력
		printResult(result)

		// 3s 인가 ?
		if isGameEnd(result) {
			// game over
			break
		}
	}

	// How many times did you hit it?
	fmt.Printf("%d번 만에 맞췄습니다.\n", cnt)
}
