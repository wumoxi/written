package stack

import "strconv"

const LIMIT = 10

type Stack struct {
	ix   int // 第一个元素是自由位置，所以data[ix] == 0
	data [LIMIT]int
}

// 入栈
func (s *Stack) Push(n int) {
	if s.ix+1 > LIMIT {
		return // 栈已满
	}
	s.data[s.ix] = n
	s.ix++
}

// 出栈
func (s *Stack) Pop() int {
	s.ix--
	return s.data[s.ix]
}

func (s *Stack) String() string {
	str := ""
	for ix := 0; ix < s.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(s.data[ix]) + "]"
	}
	return str
}
