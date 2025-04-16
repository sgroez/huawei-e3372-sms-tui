package ui

import (
	"github.com/rivo/tview"
)

func CreateFrame(title string, content tview.Primitive) *tview.Frame {
	return tview.NewFrame(content).
	SetBorders(0, 0, 0, 0, 1, 1).
	AddText(title, true, tview.AlignCenter, tview.Styles.SecondaryTextColor)
}