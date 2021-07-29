package goi

import "github.com/gotk3/gotk3/gtk"

type TabSwitcher struct {
	switcher   gtk.IWidget
	stack      *gtk.Stack
	widget     *gtk.Box
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
	var err error
	t.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}
	t.switcher, err = gtk.StackSwitcherNew()
	if err != nil {
		return nil, err
	}
	t.switcher.(*gtk.StackSwitcher).SetStack(t.stack)
	return t, t.setupWidget()
}

func NewTabSwitcherVertical() (*TabSwitcher, error) {
	t := &TabSwitcher{
		isVertical: true,
		pages:      make(map[string]Component),
	}
	var err error
	t.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}
	t.switcher, err = gtk.StackSidebarNew()
	if err != nil {
		return nil, err
	}
	t.switcher.(*gtk.StackSidebar).SetStack(t.stack)
	return t, t.setupWidget()
}

// AddTab adds a tab to the tab switcher
func (t *TabSwitcher) AddTab(title, id string, c Component) {
	t.idList = append(t.idList, id)
	t.pages[id] = c
	t.stack.AddTitled(c.getWidget(), title, id)
}

func (t *TabSwitcher) setupWidget() error {
	orient := gtk.ORIENTATION_VERTICAL
	if t.isVertical {
		orient = gtk.ORIENTATION_HORIZONTAL
	}

	box, err := gtk.BoxNew(orient, 0)
	if err != nil {
		return err
	}

	box.PackStart(t.switcher, false, false, 0)
	box.PackStart(t.stack, true, true, 0)
	t.widget = box
	return nil
}
