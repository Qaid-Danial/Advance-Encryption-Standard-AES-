package main

import (
	keygen "AES/KeyGen"
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	filePath string = "RoundKeys.txt"
)

func main() {

	app := createWindow()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func createWindow() *tview.Application {
	window := tview.NewApplication().EnableMouse(true)

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	input := tview.NewInputField()
	input.SetLabel("Password: ")

	textPanel := tview.NewTextView()
	textPanel.SetBorder(true)

	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {

			textPanel.Clear()

			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				file.Close()
				panic(err)
			}

			password := input.GetText()
			roundKeysString, roundKeysByte := keygen.GenerateRoundKey(password)

			for i, element := range roundKeysString {
				fmt.Fprintf(textPanel, "Key %d: %s\n", i+1, element)
				fmt.Fprintf(file, "%x02\n", roundKeysByte[i])
			}

			input.SetText("")
			file.Close()

		}
	})

	flex.AddItem(input, 0, 1, true)
	flex.AddItem(textPanel, 0, 15, false)

	window.SetRoot(flex, true)

	return window
}
