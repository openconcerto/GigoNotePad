package main

import (
	"fmt"
)

// TextLine represents a line of text with its index information and related details.
type TextLine struct {
	globalIndexOfFirstChar int64
	indexOfFirstChar       int64
	line                   *Line
	str                    string
	isEndOfLine            bool
}

// NewTextLine creates a new TextLine instance with the given parameters.
func NewTextLine(globalIndexOfFirstChar, indexOfFirstChar int64, line *Line, str string, isEndOfLine bool) *TextLine {
	return &TextLine{
		globalIndexOfFirstChar: globalIndexOfFirstChar,
		indexOfFirstChar:       indexOfFirstChar,
		line:                   line,
		str:                    str,
		isEndOfLine:            isEndOfLine,
	}
}

// GetGlobalIndexOfFirstChar returns the global index of the first character.
func (tl *TextLine) GetGlobalIndexOfFirstChar() int64 {
	return tl.globalIndexOfFirstChar
}

// GetIndexOfFirstChar returns the index of the first character within the line.
func (tl *TextLine) GetIndexOfFirstChar() int64 {
	return tl.indexOfFirstChar
}

// GetLine returns the associated Line object.
func (tl *TextLine) GetLine() *Line {
	return tl.line
}

// GetText returns the text string of the line.
func (tl *TextLine) GetText() string {
	return tl.str
}

// IsEndOfLine returns whether this line is the end of a paragraph or not.
func (tl *TextLine) IsEndOfLine() bool {
	return tl.isEndOfLine
}

// SetEndOfLine sets whether this line is the end of a paragraph or not.
func (tl *TextLine) SetEndOfLine(isEndOfLine bool) {
	tl.isEndOfLine = isEndOfLine
}

// Length returns the length of the text string in the line.
func (tl *TextLine) Length() int {
	return len(tl.str)
}

// String returns a string representation of the TextLine.
func (tl *TextLine) String() string {
	return fmt.Sprintf("TextLine [globalIndexOfFirstChar=%d, indexOfFirstChar=%d, str=%s, isEndOfLine=%t, line=%v]",
		tl.globalIndexOfFirstChar, tl.indexOfFirstChar, tl.str, tl.isEndOfLine, tl.line)
}
