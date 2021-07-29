package goi

import "github.com/gotk3/gotk3/gtk"

type TabSwitcher struct {
	widget     gtk.IWidget
	isVertical bool

	pages  map[string]Component
	idList []string
}

func (t *TabSwitcher) getWidget() gtk.IWidget {
	return t.widget
}

func NewTabSwitcher() (*TabSwitcher, error) {
	t := &TabSwitcher{
		isVertical: false,
		pages:      make(map[string]Component),
	}
	stack, err := gtk.StackNew()
	if err != nil {
		return nil, err
	}
	t.widget, err = gtk.StackSwitcherNew()
	if err != nil {
		return nil, err
	}
	t.widget.(*gtk.StackSwitcher).SetStack(stack)
	return t, nil
}

func NewTabSwitcherVertical() (*TabSwitcher, error) {
	t := &TabSwitcher{
		isVertical: true,
		pages:      make(map[string]Component),
	}
	stack, err := gtk.StackNew()
	if err != nil {
		return nil, err
	}
	t.widget, err = gtk.StackSidebarNew()
	if err != nil {
		return nil, err
	}
	t.widget.(*gtk.StackSidebar).SetStack(stack)
	return t, nil
}

// AddPage adds a tab to the tab switcher, where the tab name is unique
func (t *TabSwitcher) AddTab(name string, c Component) {
	var widg *gtk.Stack
	if t.isVertical {
		widg = t.widget.(*gtk.StackSidebar).GetStack()
	} else {
		widg = t.widget.(*gtk.StackSwitcher).GetStack()
	}
	t.idList = append(t.idList, name)
	t.pages[name] = c
	widg.AddNamed(c.getWidget(), name)
}
