package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type TextEditorPanel struct {
	widget.BaseWidget
	doc                         *Document
	firstVisibleLineGlobalIndex uint64
	cursorGlobalIndex           uint64
	maxCharactersPerLine        int
	lines                       []TextLine
}

func NewTextEditorPanel() *TextEditorPanel {
	editor := &TextEditorPanel{
		maxCharactersPerLine: 40,
		lines:                []TextLine{},
	}
	editor.ExtendBaseWidget(editor)
	return editor
}

func (editor *TextEditorPanel) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(color.White)
	return &TextEditorRenderer{
		editor:     editor,
		background: bg,
		objects:    []fyne.CanvasObject{bg},
	}
}

func (editor *TextEditorPanel) GetDocument() *Document {
	return editor.doc
}

type TextEditorRenderer struct {
	editor     *TextEditorPanel
	background *canvas.Rectangle
	objects    []fyne.CanvasObject
}

func (renderer *TextEditorRenderer) Layout(size fyne.Size) {
	renderer.background.Resize(size)
}

func (renderer *TextEditorRenderer) MinSize() fyne.Size {
	return fyne.NewSize(200, 200)
}

func (renderer *TextEditorRenderer) Refresh() {
	renderer.background.FillColor = color.White
	canvas.Refresh(renderer.editor)
}

func (renderer *TextEditorRenderer) BackgroundColor() color.Color {
	return color.White
}

func (renderer *TextEditorRenderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *TextEditorRenderer) Destroy() {}
