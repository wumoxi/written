package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// Error indicates an error in converting a word into an integer. (指示将单词转换为整数时出错)
type Error struct {
	Index int    // The index into the space-separated list of words. (空格分隔的单词列表的索引)
	Word  string // The word that generated the parse error. (生成解析错误的单词)
	Err   error  // The raw error that precipitated this error, if any. (引发此错误的原始错误(如果有的话))
}

// String returns a human-readable error message. (返回人类可读的错误消息)
func (e *Error) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in input as integers.(将输入中的空格分隔的单词解析为整数)
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2number(fields)
	return
}

func fields2number(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}

	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&Error{Index: idx, Word: field, Err: err})
		}
		numbers = append(numbers, num)
	}
	return
}
