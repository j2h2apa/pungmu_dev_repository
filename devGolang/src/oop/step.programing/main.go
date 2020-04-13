package main

/* 절차적 프로그래밍의 장 단 점 */

import (
	"fmt"
)

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonStrawberry struct {
}

type Sandwitch struct {
	val string
}

type OrangeJam struct {
	opended bool
}

func GetBreads(num int) []*Bread {
	var breads []*Bread = make([]*Bread, 2)

	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}

	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

/*OpenOrangeJam : 절차적 프로그램의 단점으로 strawberryJam 과 같은 기능
  을 하는 함수를 또 만들어야 한다.
*/
func OpenOrangeJam(jam *OrangeJam) {
	jam.opended = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonStrawberry {
	return &SpoonStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonStrawberry) {
	bread.val += " + Strawberry Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	var sandwitch *Sandwitch = &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}

	return sandwitch
}

func main() {
	var jam *StrawberryJam = &StrawberryJam{}

	// 1. 빵두개를 꺼낸다
	breads := GetBreads(2)
	// 2. 딸기잼 뚜껑을 연다
	OpenStrawberryJam(jam)
	// 3. 딸기잼을 한수푼 뜬다.
	var spoon *SpoonStrawberry = GetOneSpoon(jam)
	// 4. 딸기잼을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)
	// 5. 빵을 덮는다.
	var sandwitch *Sandwitch = MakeSandwitch(breads)
	// 6. 완성
	fmt.Println(sandwitch.val)
}
