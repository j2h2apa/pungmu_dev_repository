package main

import (
	"datastruct"
	"fmt"
)

type key struct {
	v int
}

type value struct {
	v int
}

func main() {
	fmt.Println("abcde = ", datastruct.Hash("abcde"))
	fmt.Println("abcde = ", datastruct.Hash("abcde"))
	fmt.Println("abcdf = ", datastruct.Hash("abcdf"))
	fmt.Println("abcdedfsdasd123 = ", datastruct.Hash("abcdedfsdasd123"))

	var m *datastruct.Map = datastruct.CreateMap()
	m.Add("AAA", "010-777-7777")
	m.Add("BBB", "010-777-7778")
	m.Add("CCC", "010-777-7779")
	m.Add("DDD", "010-777-7770")
	m.Add("EEE", "010-777-7771")

	fmt.Println(m.Get("AAA"))
	fmt.Println(m.Get("EEE"))
	fmt.Println(m.Get("CCC"))
	fmt.Println(m.Get("FFF"))

	/* reserve map 사용 기본 */
	// case 1
	var m1 map[string]string = make(map[string]string)
	m1["abc"] = "bbb"
	fmt.Println(m1["abc"])

	// case 2
	m2 := make(map[int]string)
	m2[53] = "ddd"
	fmt.Println(m2[53])

	fmt.Println("m2[55] : ", m2[55])

	// case 3
	m3 := make(map[int]int)
	m3[4] = 4

	// int value 기본 값은 0이다
	fmt.Println("m3[0] : ", m3[10])

	m3[5] = 0
	fmt.Println("m3[5] m3[10] : ", m3[5], m3[10])

	v, ok1 := m3[5]
	v1 := m3[4]
	// map node 는 2가지 리턴 value 와 값 셋팅여부(bool)
	v2, ok2 := m3[10]
	fmt.Println(v, ok1, v1, v2, ok2)

	// node remove
	delete(m3, 5)
	v, ok1 = m3[5]
	fmt.Println(v, ok1, v1, v2, ok2)

	// rounding
	m3[2] = 98
	m3[10] = 1245
	for key, value := range m3 {
		fmt.Println("m3[", key, "] = ", value)
	}

	// case 4 : bool value 는 참 값 셋팅 현황을 알 수 있다.
	m4 := make(map[int]bool)
	m4[4] = true
	fmt.Println(m4[4], m4[6])
}
