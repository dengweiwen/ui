// 14 february 2014
package ui

import (
	"sync"
)

// A LineEdit is a control which allows you to enter a single line of text.
type LineEdit struct {
	// TODO Typing event

	lock		sync.Mutex
	created	bool
	sysData	*sysData
	initText	string
}

// NewLineEdit makes a new LineEdit with the specified text.
func NewLineEdit(text string) *LineEdit {
	return &LineEdit{
		sysData:	mksysdata(c_lineedit),
		initText:	text,
	}
}

// SetText sets the LineEdit's text.
func (l *LineEdit) SetText(text string) (err error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.created {
		return l.sysData.setText(text)
	}
	l.initText = text
	return nil
}

// Text returns the LineEdit's text.
func (l *LineEdit) Text() string {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.created {
		return l.sysData.text()
	}
	return l.initText
}

func (l *LineEdit) make(window *sysData) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	err := l.sysData.make(l.initText, window)
	if err != nil {
		return err
	}
	l.created = true
	return nil
}

func (l *LineEdit) setRect(x int, y int, width int, height int) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.sysData.setRect(x, y, width, height)
}

func (l *LineEdit) preferredSize() (width int, height int, err error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	width, height = l.sysData.preferredSize()
	return
}
