package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type Line struct {
	length          int64
	lengthWithEOL   int64
	lineIndex       int
	parts           []string
	carriageReturn  bool
	endsWithNewLine bool
}

// NewLine is a constructor for the Line struct
func NewLine(parts []string, lineIndex int) *Line {
	line := &Line{
		parts:     make([]string, len(parts)),
		lineIndex: lineIndex,
	}
	copy(line.parts, parts)
	for _, part := range line.parts {
		line.length += int64(len(part))
	}
	line.computeLengthWithEOL()
	return line
}

// Dump prints the details of the Line object
func (l *Line) Dump() {
	size := len(l.parts)
	fmt.Printf("--line (%d parts) lineIndex: %d length=%d, index=%d carriageReturn=%v endsWithNewLine:%v\n",
		size, l.lineIndex, l.length, l.lineIndex, l.carriageReturn, l.endsWithNewLine)
	for i, part := range l.parts {
		fmt.Printf("part %d:'%s'\n", i, part)
	}
}

func (l *Line) GetParts() []string {
	return l.parts
}

func (l *Line) Length() int64 {
	return l.length
}

func (l *Line) GetLengthWithEOL() int64 {
	return l.lengthWithEOL
}

func (l *Line) computeLengthWithEOL() {
	if l.endsWithNewLine {
		if l.carriageReturn {
			l.lengthWithEOL = l.length + 2
		} else {
			l.lengthWithEOL = l.length + 1
		}
	} else {
		l.lengthWithEOL = l.length
	}
}

func (l *Line) GetLineIndex() int {
	return l.lineIndex
}

func (l *Line) SetLineIndex(index int) {
	l.lineIndex = index
}

func (l *Line) CharAt(index int64) (rune, error) {
	if l.carriageReturn {
		if index == l.length {
			return '\r', nil
		} else if index == l.length+1 {
			return '\n', nil
		}
	} else {
		if index == l.length {
			return '\n', nil
		}
	}

	i := int64(0)
	for _, part := range l.parts {
		i += int64(len(part))
		if index < i {
			index2 := int(index - (i - int64(len(part))))
			return rune(part[index2]), nil
		}
	}
	return 0, errors.New(fmt.Sprintf("index %d is invalid, length is %d", index, l.length))
}

func (l *Line) SetUseCarriageReturn(b bool) {
	l.carriageReturn = b
	l.computeLengthWithEOL()
}

func (l *Line) HasCarriageReturn() bool {
	return l.carriageReturn
}

func (l *Line) GetText() string {
	return l.GetString(0, l.length)
}

func (l *Line) GetString(index, maxLength int64) string {
	var b strings.Builder
	maxIndex := min(l.length, maxLength+index)
	partStartIndex, partEndIndex := int64(0), int64(0)
	partIndex := 0

	size := len(l.parts)
	for partIndex < size && partEndIndex <= index {
		partStartIndex = partEndIndex
		partEndIndex += int64(len(l.parts[partIndex]))
		partIndex++
	}

	for i := index; i < maxIndex; i++ {
		if i >= partEndIndex {
			if partIndex >= size {
				break
			}
			partStartIndex = partEndIndex
			partEndIndex += int64(len(l.parts[partIndex]))
			partIndex++
		}
		indexInPart := int(i - partStartIndex)
		b.WriteByte(l.parts[partIndex-1][indexInPart])
	}
	return b.String()
}

func (l *Line) String() string {
	var text string
	if l.length > 100 {
		text = l.GetString(0, 100) + " ...[truncated]"
	} else {
		text = l.GetString(0, l.length)
	}
	end := "EOF"
	if l.endsWithNewLine {
		if l.carriageReturn {
			end = "CRLF"
		} else {
			end = "LF"
		}
	}
	return fmt.Sprintf("Line [index=%d length=%d, index=%d(%d parts) %s ] %s",
		l.lineIndex, l.length, l.lineIndex, len(l.parts), end, text)
}

func (l *Line) Append(b *strings.Builder, start int64, end int64) int64 {
	if start > end {
		panic(fmt.Sprintf("start index is greater than the end offset (%d>%d)", start, end))
	}
	if end-start > int64(int(^uint(0)>>1)) {
		panic("range is too big")
	}
	if end > l.GetLengthWithEOL() {
		panic(fmt.Sprintf("end index (%d) greater than length with new line (%d)", end, l.length+1))
	}

	b.WriteString(l.GetString(start, end-start))
	if end > l.length {
		b.WriteByte('\n')
	}
	return end - start
}

func (l *Line) Insert(indexInLine int64, text string) {
	if len(text) == 0 {
		return
	}
	if indexInLine == l.length {
		indexOfLast := len(l.parts) - 1
		s := l.parts[indexOfLast] + text
		l.parts[indexOfLast] = s
		l.length += int64(len(text))
		l.computeLengthWithEOL()
		return
	}

	i := int64(0)
	for partIndex := 0; partIndex < len(l.parts); partIndex++ {
		part := l.parts[partIndex]
		i += int64(len(part))
		if indexInLine < i {
			index2 := int(indexInLine - (i - int64(len(part))))
			before := part[:index2]
			after := part[index2:]
			newString := before + text + after
			l.parts[partIndex] = newString
			l.length += int64(len(text))
			l.computeLengthWithEOL()
			return
		}
	}
}

func (l *Line) SetEndsWithNewLine(b bool) {
	l.endsWithNewLine = b
	l.computeLengthWithEOL()
}

func (l *Line) EndsWithNewLine() bool {
	return l.endsWithNewLine
}

func (l *Line) Delete(start, end int64) {
	if start < 0 || end < start || end > l.GetLengthWithEOL() {
		panic(fmt.Sprintf("Invalid start or end index (%d, %d). Line length: %d (%d with EOL)", start, end, l.length, l.GetLengthWithEOL()))
	}

	if (end == l.length+1 || end == l.length+2) && l.endsWithNewLine {
		l.endsWithNewLine = false
	}

	currentIndex := int64(0)
	for i := 0; i < len(l.parts); i++ {
		part := l.parts[i]
		partLength := int64(len(part))

		if currentIndex < end && currentIndex+partLength > start {
			startIndexInPart := maxInt64(0, start-currentIndex)
			endIndexInPart := minInt64(partLength, end-currentIndex)

			l.parts[i] = part[:startIndexInPart] + part[endIndexInPart:]
			currentIndex += partLength
			l.length -= endIndexInPart - startIndexInPart
			l.computeLengthWithEOL()
		} else {
			currentIndex += partLength
		}
	}
}

func (l *Line) IndexOf(text string, fromIndex int64) int64 {
	stop := len(l.parts)
	textLength := int64(len(text))
	index := int64(0)

	for i := 0; i < stop; i++ {
		part := l.parts[i]
		if index+int64(len(part)) < fromIndex {
			index += int64(len(part))
			continue
		}

		if foundIndex := strings.Index(part[int(fromIndex-index):], text); foundIndex != -1 {
			return index + int64(foundIndex)
		}

		if i < stop-1 {
			nextPart := l.parts[i+1]
			if len(part) > int(textLength) && len(nextPart) >= int(textLength)-1 {
				start := len(part) - int(textLength)
				s := part[start:] + nextPart[:int(textLength)-1]
				if foundIndex := strings.Index(s, text); foundIndex != -1 {
					return index + int64(start+foundIndex)
				}
			}
		}
		index += int64(len(part))
	}
	return -1
}

func (l *Line) Find(text string, globalIndexOfLine int64, result *[]int64) {
	foundIndex := l.IndexOf(text, 0)
	for foundIndex >= 0 {
		*result = append(*result, foundIndex+globalIndexOfLine)
		foundIndex = l.IndexOf(text, foundIndex+int64(len(text)))
	}
}

func (l *Line) WriteTo(bOut *bufio.Writer, charset string, separator LineSeparator) error {
	stop := len(l.parts)
	for i := 0; i < stop; i++ {
		if _, err := bOut.WriteString(l.parts[i]); err != nil {
			return err
		}
	}
	if l.endsWithNewLine {
		switch separator {
		case AUTO:
			if l.carriageReturn {
				if _, err := bOut.WriteRune('\r'); err != nil {
					return err
				}
			}
			if _, err := bOut.WriteRune('\n'); err != nil {
				return err
			}
		case LF:
			if _, err := bOut.WriteRune('\n'); err != nil {
				return err
			}
		case CRLF:
			if _, err := bOut.WriteRune('\r'); err != nil {
				return err
			}
			if _, err := bOut.WriteRune('\n'); err != nil {
				return err
			}
		}
	}
	return nil
}

// Helper functions
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// Enum equivalent for LineSeparator
type LineSeparator int

const (
	AUTO LineSeparator = iota
	LF
	CRLF
)
