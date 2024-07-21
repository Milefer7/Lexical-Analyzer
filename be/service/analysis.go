package service

import (
	"errors"
	model2 "github.com/Milefer7/compliation_exp/dal/model"
	"github.com/Milefer7/compliation_exp/dao/mysql"
	"log"
	"strings"
	"unicode"
)

var DoubleDelimiters = map[string]string{":": "="}

// 自动机词法分析函数
func LexicalAnalysis(code string) ([]model2.LexicalAnalysis, error) {
	// 数据准备
	Keywords, err := mysql.NewKeywordsMySQL().ReadKeywords() // 准备关键字
	if err != nil {
		return nil, err
	}
	Delimiters, err := mysql.NewDelimiterMySQL().ReadDelimiters() // 准备分界符
	if err != nil {
		return nil, err
	}

	// 结果放在result中
	result := []model2.LexicalAnalysis{}
	// 将代码分割成行
	lines := strings.Split(code, "\n")
	lineNum := 0 // 添加行号计数器
	log.Printf("lines: %v", lines)
	for _, line := range lines {
		lineNum++
		i := 0
		for i < len(line) {
			ch := line[i]

			// 跳过空白字符
			if unicode.IsSpace(rune(ch)) {
				i++
				continue
			}

			// 处理数字
			if unicode.IsDigit(rune(ch)) {
				start := i
				for i < len(line) && unicode.IsDigit(rune(line[i])) {
					i++
				}
				value := line[start:i]
				result = append(result, model2.LexicalAnalysis{LineNum: lineNum, Content: value, Type: "无符号整数"})
				continue
			}

			// 处理字母（标识符或关键字）
			if unicode.IsLetter(rune(ch)) {
				start := i
				for i < len(line) && (unicode.IsLetter(rune(line[i])) || unicode.IsDigit(rune(line[i]))) {
					i++
				}
				value := line[start:i]
				if isKeyword(value, Keywords) {
					result = append(result, model2.LexicalAnalysis{LineNum: lineNum, Content: value, Type: "关键字"})
				} else {
					result = append(result, model2.LexicalAnalysis{LineNum: lineNum, Content: value, Type: "标识符"})
				}
				continue
			}

			// 处理单字符分界符
			if isDelimiter(string(ch), Delimiters) {
				result = append(result, model2.LexicalAnalysis{LineNum: lineNum, Content: string(ch), Type: "单字符分界符"})
				i++
				continue
			}

			// 处理双字符分界符
			if isDoubleDelimiter(string(ch)) {
				if i+1 < len(line) && isDoubleDelimiter(string(ch)+string(line[i+1])) {
					result = append(result, model2.LexicalAnalysis{LineNum: lineNum, Content: string(ch) + string(line[i+1]), Type: "双字符分界符"})
					i += 2
					continue
				} else {
					return nil, errors.New("Invalid token")
				}
			}

			// 处理注释
			if ch == '{' {
				i++
				for i < len(line) && line[i] != '}' {
					i++
				}
				if i < len(line) {
					i++ // 跳过结束的'}'
				} else {
					return nil, errors.New("Unclosed comment")
				}
				continue
			}

			// 处理数组下标界限符
			if ch == '.' {
				if i+1 < len(line) && line[i+1] == '.' {
					result = append(result, model2.LexicalAnalysis{LineNum: lineNum, Content: "..", Type: "双字符分界符"})
					i += 2
					continue
				} else {
					result = append(result, model2.LexicalAnalysis{LineNum: lineNum, Content: ".", Type: "程序结束"})
					i++
					continue
				}
			}

			// 未识别的符号
			return nil, errors.New("Invalid token: " + string(ch))
		}
	}
	return result, nil
}

// 判断是否是关键字
func isKeyword(word string, keywords []model2.Keywords) bool {
	for _, keyword := range keywords {
		if word == keyword.Keyword {
			return true
		}
	}
	return false
}

// 判断是否是单字符分界符
func isDelimiter(ch string, delimiters []model2.Delimiter) bool {
	for _, delimiter := range delimiters {
		if ch == delimiter.Name {
			return true
		}
	}
	return false
}

// 判断是否是双字符分界符
func isDoubleDelimiter(ch string) bool {
	for prefix, suffix := range DoubleDelimiters {
		if ch == prefix+suffix {
			return true
		}
	}
	return false
}
