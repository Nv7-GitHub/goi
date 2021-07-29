package goi

import "github.com/gotk3/gotk3/gtk"

type TabSwitcher struct {
	widget     *gtk.Notebook
	isVertical bool

	pages []Component
}

func (t *TabSwitcher) getWidget() gtk.IWidget {
	return t.widget
}

func NewTabSwitcher() (*TabSwitcher, error) {
	t := &TabSwitcher{
		isVertical: false,
	}
	var err error
	t.widget, err = gtk.NotebookNew()
	return t, err
}

func NewTabSwitcherVertical() (*TabSwitcher, error) {
	t := &TabSwitcher{
		isVertical: false,
	}
	var err error
	t.widget, err = gtk.NotebookNew()
	if err != nil {
		return nil, err
	}
	t.widget.SetTabPos(gtk.POS_LEFT)
	return t, err
}

// AddTab adds a tab to the tab switcher
func (t *TabSwitcher) AppendTab(title string, c Component) error {
	t.pages = append(t.pages, c)
	label, err := gtk.LabelNew(title)
	if err != nil {
		return err
	}
	t.widget.AppendPage(c.getWidget(), label)
	return nil
}
