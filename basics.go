package goi

import "github.com/gotk3/gotk3/gtk"

type Label struct {
	Text string

	widget *gtk.Label
}

func (l *Label) getWidget() gtk.IWidget {
	return l.widget
}

func NewLabel(text string) (*Label, error) {
	var err error
	l := &Label{}
	l.Text = text
	l.widget, err = gtk.LabelNew(text)
	return l, err
}

func (l *Label) SetText(newText string) {
	l.Text = newText
	l.widget.SetText(newText)
}
