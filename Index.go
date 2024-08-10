package main

import "fmt"

type Index struct {
	lineIndex       int
	charIndexInLine int64
}

// NewIndex creates a new Index instance.
func NewIndex(lineIndex int, charIndexInLine int64) *Index {
	return &Index{
		lineIndex:       lineIndex,
		charIndexInLine: charIndexInLine,
	}
}

// GetLineIndex returns the line index.
func (i *Index) GetLineIndex() int {
	return i.lineIndex
}

// GetCharIndexInLine returns the character index within the line.
func (i *Index) GetCharIndexInLine() int64 {
	return i.charIndexInLine
}

// String returns a string representation of the Index.
func (i *Index) String() string {
	return fmt.Sprintf("[Line Index: %d charIndexInLine: %d]", i.GetLineIndex(), i.GetCharIndexInLine())
}
