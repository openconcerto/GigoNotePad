package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const CHUNK_SIZE = 1 * 1024 * 1024

// Document struct containing a list of lines and other attributes
type Document struct {
	lines       []Line
	totalLength int64
}

func NewDocument() *Document {
	doc := &Document{
		lines:       make([]Line, 0),
		totalLength: -1,
	}
	emptyParts := []string{""}
	doc.lines = append(doc.lines, Line{parts: emptyParts})
	return doc
}

func (doc *Document) GetLines() []Line {
	return doc.lines
}

func (doc *Document) GetLineCount() int {
	return len(doc.lines)
}

func (doc *Document) PreLoadFrom(file string, skip int, charset string, max int) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if skip > 0 {
		_, err = f.Seek(int64(skip), io.SeekStart)
		if err != nil {
			return err
		}
	}

	a := make([]byte, 500000)
	_, err = f.Read(a)
	if err != nil && err != io.EOF {
		return err
	}

	start := string(a)
	doc.LoadFromString(start, max)
	return nil
}

func (doc *Document) LoadFrom(file string, skip int, charset string, max int) error {
	fmt.Printf("Document.LoadFrom() %s  from %d %s\n", file, skip, charset)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if skip == 0 {
		all := string(data)
		doc.LoadFromString(all, max)
	} else {
		b := bytes.NewBuffer(data[skip:])
		all := b.String()
		doc.LoadFromString(all, max)
	}

	return nil
}

func (doc *Document) LoadFromString(str string, maxPartSize int) {
	doc.lines = nil // clear the lines
	startTime := time.Now()

	size := len(str)
	var parts []string
	var b strings.Builder
	counter := 0
	returnFound := false

	for i := 0; i < size; i++ {
		c := str[i]

		if c == '\n' {
			if b.Len() > 0 {
				parts = append(parts, b.String())
			}
			line := Line{parts: parts, lineIndex: len(doc.lines)}
			line.endsWithNewLine = true
			line.carriageReturn = returnFound
			doc.lines = append(doc.lines, line)
			b.Reset()
			counter = 0
			parts = nil
			returnFound = false
		} else if c == '\r' {
			returnFound = true
		} else {
			if counter > maxPartSize {
				parts = append(parts, b.String())
				b.Reset()
				counter = 0
			}
			b.WriteByte(c)
		}

		counter++
	}

	if b.Len() > 0 {
		parts = append(parts, b.String())
	}
	line := Line{parts: parts, lineIndex: len(doc.lines)}
	doc.lines = append(doc.lines, line)

	fmt.Printf("Document.LoadFromString() took %dms\n", time.Since(startTime).Milliseconds())
	doc.invalidLength()
}

func (doc *Document) Dump(out io.Writer) {
	for _, line := range doc.lines {
		line.dump(out)
	}
}

func (doc *Document) invalidLength() {
	doc.totalLength = -1
}

func (doc *Document) Save(file string, charset string, lineSeparator LineSeparator) error {
	if charset == "" {
		return fmt.Errorf("null charset")
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, line := range doc.lines {
		err := line.writeTo(writer, charset, lineSeparator)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// Additional Methods like `createTextLines`, `getGlobalIndex`, `delete`, `insert`, etc. would need to be implemented in a similar fashion, but for brevity, only a subset of the Java methods have been translated.

func (line *Line) dump(out io.Writer) {
	for _, part := range line.parts {
		fmt.Fprintln(out, part)
	}
}

func (line *Line) writeTo(out *bufio.Writer, charset string, separator LineSeparator) error {
	for _, part := range line.parts {
		if _, err := out.WriteString(part); err != nil {
			return err
		}
	}

	if line.endsWithNewLine {
		switch separator {
		case AUTO, LF:
			if _, err := out.WriteRune('\n'); err != nil {
				return err
			}
		case CRLF:
			if _, err := out.WriteString("\r\n"); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	fmt.Println("Document main:")
	// Example usage
	doc := NewDocument()
	err := doc.LoadFrom("F:\\test-10000000.xml", 0, "UTF-8", CHUNK_SIZE)
	//err := doc.LoadFrom("F:\\test-100.xml", 0, "UTF-8", CHUNK_SIZE)
	if err != nil {
		fmt.Println("Error loading document:", err)
	}
	duration := time.Duration(1000) * time.Second // Pause for 10 seconds
	time.Sleep(duration)

	//doc.Dump(os.Stdout)
	if false {
		err = doc.Save("output.txt", "UTF-8", AUTO)
		if err != nil {
			fmt.Println("Error saving document:", err)
		}
	}
}
