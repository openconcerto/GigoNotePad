package main

// TextLinePosition represents the position of a character within a TextLine.
type TextLinePosition struct {
	Line        *TextLine // Reference to the TextLine
	IndexInLine int       // Index of the character within the line
}
