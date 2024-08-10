package main

import (
	"fmt"
)

// Selection represents a selection range with initial, start, and end indices.
type Selection struct {
	initIndex  int64
	startIndex int64
	endIndex   int64
}

// NewSelection creates a new Selection instance and initializes it with the given index.
func NewSelection(initIndex int64) *Selection {
	s := &Selection{}
	s.Init(initIndex)
	return s
}

// Init initializes the selection with the given index.
func (s *Selection) Init(initIndex int64) {
	s.initIndex = initIndex
	s.startIndex = initIndex
	s.endIndex = initIndex
}

// GetInitIndex returns the initial index of the selection.
func (s *Selection) GetInitIndex() int64 {
	return s.initIndex
}

// SetRange sets the start and end indices of the selection.
func (s *Selection) SetRange(startIndex, endIndex int64) {
	if startIndex > endIndex {
		panic(fmt.Sprintf("startIndex > endIndex (%d > %d)", startIndex, endIndex))
	}
	s.startIndex = startIndex
	s.endIndex = endIndex
}

// SetStartIndex sets the start index of the selection.
func (s *Selection) SetStartIndex(index int64) {
	s.startIndex = index
}

// GetStartIndex returns the start index of the selection.
func (s *Selection) GetStartIndex() int64 {
	return s.startIndex
}

// SetEndIndex sets the end index of the selection.
func (s *Selection) SetEndIndex(index int64) {
	s.endIndex = index
}

// GetEndIndex returns the end index of the selection.
func (s *Selection) GetEndIndex() int64 {
	return s.endIndex
}

// IsEmpty checks if the selection is empty (start index equals end index).
func (s *Selection) IsEmpty() bool {
	return s.startIndex == s.endIndex
}

// String returns a string representation of the selection.
func (s *Selection) String() string {
	return fmt.Sprintf("Selection [initIndex=%d, startIndex=%d, endIndex=%d]", s.initIndex, s.startIndex, s.endIndex)
}
