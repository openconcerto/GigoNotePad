package main

import (
	"fmt"
	"hash/fnv"
)

// Highlight represents a highlighted range with a start and end index.
type Highlight struct {
	startIndex int64
	endIndex   int64
}

// NewHighlight creates a new Highlight instance.
func NewHighlight(startIndex, endIndex int64) *Highlight {
	return &Highlight{
		startIndex: startIndex,
		endIndex:   endIndex,
	}
}

// GetStartIndex returns the start index of the highlight.
func (h *Highlight) GetStartIndex() int64 {
	return h.startIndex
}

// GetEndIndex returns the end index of the highlight.
func (h *Highlight) GetEndIndex() int64 {
	return h.endIndex
}

// CompareTo compares two Highlight instances by their start index.
func (h *Highlight) CompareTo(o *Highlight) int {
	if h.startIndex < o.startIndex {
		return -1
	} else if h.startIndex > o.startIndex {
		return 1
	}
	return 0
}

// Equals checks if two Highlight instances are equal.
func (h *Highlight) Equals(obj interface{}) bool {
	if other, ok := obj.(*Highlight); ok {
		return h.startIndex == other.startIndex && h.endIndex == other.endIndex
	}
	return false
}

// HashCode returns a hash code for the highlight.
func (h *Highlight) HashCode() int {
	hasher := fnv.New32a()
	hasher.Write([]byte(fmt.Sprintf("%d%d", h.startIndex, h.endIndex)))
	return int(hasher.Sum32())
}

// Contains checks if the highlight contains the given index.
func (h *Highlight) Contains(index int64) bool {
	return h.startIndex <= index && index < h.endIndex
}

// String returns a string representation of the highlight.
func (h *Highlight) String() string {
	return fmt.Sprintf("Highlight[%d, %d]", h.startIndex, h.endIndex)
}
