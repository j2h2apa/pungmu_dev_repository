package datastruct

/*Hash : rolling hash
input 의 범위가 지정되어야 한다
output : input 에 대해 값은 정해져 있다
*/
func Hash(s string) int {
	var h int = 0
	A := 256
	B := 3571 // 범위

	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}

	return h
}

//////////////////////////////////////////////////////////////////////////////
type keyValue struct {
	key   string
	value string
}

/*Map : key, value 로 이루어진 datastruct*/
type Map struct {
	keyArray [3571][]keyValue
}

/*CreateMap : map create*/
func CreateMap() *Map {
	return &Map{}
}

/*Add : node append (hash map)*/
func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}

/*Get : node value 가져오기*/
func (m *Map) Get(key string) string {
	var h int = Hash(key)

	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}

	return "don't find value"
}
