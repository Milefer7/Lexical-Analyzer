package model

type LexicalCode struct {
	Code string `json:"code"`
}

type LexicalAnalysis struct {
	LineNum int    `json:"line_num"`
	Content string `json:"content"`
	Type    string `json:"type"`
}
