package goi

import (
	"testing"
)

const width = 800
const height = 600

func TestTabSwitcher(t *testing.T) {
	win, err := NewWindow("Test Tab Switcher")
	if err != nil {
		t.Fatal(err)
	}

	switcher, err := NewTabSwitcherVertical()
	if err != nil {
		t.Fatal(err)
	}

	label, err := NewLabel("This is page 1")
	if err != nil {
		t.Fatal(err)
	}

	label2, err := NewLabel("This is page 2")
	if err != nil {
		t.Fatal(err)
	}

	err = switcher.AppendTab("Page 1", label)
	if err != nil {
		t.Fatal(err)
	}

	err = switcher.AppendTab("Page 2", label2)
	if err != nil {
		t.Fatal(err)
	}

	win.SetChild(switcher)
	win.SetDefaultSize(width, height)
	win.Show()
	Run()
}
