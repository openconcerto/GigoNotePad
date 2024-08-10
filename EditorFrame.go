package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

type EditorFrame struct {
	window             fyne.Window
	labelFileName      *widget.Label
	labelSelection     *widget.Label
	labelCurrentLine   *widget.Label
	labelCurrentColumn *widget.Label
	labelCurrentIndex  *widget.Label
	editor             *TextEditorPanel
	file               *os.File
	charset            string
	needSave           bool
	lineSeparator      string
}

func NewEditorFrame(a fyne.App) *EditorFrame {
	frame := &EditorFrame{}

	// Initialize labels and editor
	frame.labelFileName = widget.NewLabel("")
	frame.labelSelection = widget.NewLabel("")
	frame.labelCurrentLine = widget.NewLabel("")
	frame.labelCurrentColumn = widget.NewLabel("")
	frame.labelCurrentIndex = widget.NewLabel("")
	frame.editor = NewTextEditorPanel()

	// Setup the main window
	frame.window = a.NewWindow("GigaNotePad")
	frame.window.SetContent(container.NewBorder(
		container.NewVBox(
			frame.labelFileName,
			container.NewHBox(frame.labelSelection, frame.labelCurrentIndex, frame.labelCurrentLine, frame.labelCurrentColumn),
		),
		nil, nil, nil,
		frame.editor,
	))

	// Setup menu
	frame.setupMenu()

	frame.window.Resize(fyne.NewSize(800, 480))
	frame.window.Show()

	return frame
}

func (frame *EditorFrame) setupMenu() {
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				doc := NewDocument()
				frame.showNewEditor(doc)
			}),
			fyne.NewMenuItem("Open", func() { frame.openFile() }),
			fyne.NewMenuItem("Save", func() { frame.saveFile() }),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Exit", func() { frame.window.Close() }),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				widget.ShowPopUp(widget.NewLabel("About GigaNotePad"), frame.window.Canvas())
			}),
		),
	)
	frame.window.SetMainMenu(menu)
}

func (frame *EditorFrame) showNewEditor(doc *Document) {
	// TODO
}

func (frame *EditorFrame) openFile() { /**
	dialog := widget.NewFileOpen(func(reader fyne.URIReadCloser, _ error) {
		if reader == nil {
			return
		}
		defer reader.Close()

		data, err := ioutil.ReadAll(reader)
		if err != nil {
			widget.ShowPopUp(widget.NewLabel(fmt.Sprintf("Error reading file: %v", err)), frame.window.Canvas())
			return
		}
		frame.editor.SetText(string(data))
		frame.labelFileName.SetText(filepath.Base(reader.URI().Path()))
		frame.file, _ = os.Open(reader.URI().Path())
		frame.needSave = false
	}, frame.window)
	dialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".md", ".go", ".java", ".py"}))
	dialog.Show()*/
}

func (frame *EditorFrame) saveFile() {
	if frame.file == nil {
		frame.saveFileAs()
		return
	}
	/**
	err := frame.editor.GetDocument().Save()
	if err != nil {
		widget.ShowPopUp(widget.NewLabel(fmt.Sprintf("Error saving file: %v", err)), frame.window.Canvas())
	} else {
		frame.needSave = false
	}

	*/
}

func (frame *EditorFrame) saveFileAs() { /**
	dialog := widget.NewFileSave(func(writer fyne.URIWriteCloser, _ error) {
		if writer == nil {
			return
		}
		defer writer.Close()

		data := []byte(frame.editor.Text)
		_, err := writer.Write(data)
		if err != nil {
			widget.ShowPopUp(widget.NewLabel(fmt.Sprintf("Error saving file: %v", err)), frame.window.Canvas())
		} else {
			frame.labelFileName.SetText(filepath.Base(writer.URI().Path()))
			frame.file, _ = os.Open(writer.URI().Path())
			frame.needSave = false
		}
	}, frame.window)
	dialog.Show()*/
}

func main() {
	a := app.New()
	frame := NewEditorFrame(a)
	frame.window.ShowAndRun()
}
