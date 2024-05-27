package model

type LexicalCode struct {
	Code string `json:"code"`
}

type LexicalAnalysis struct {
	LineNum int    `json:"lineNum"`
	Content string `json:"content"`
	Type    string `json:"type"`
}
