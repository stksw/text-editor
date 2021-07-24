package main

import (
	"io/ioutil"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	window := myApp.NewWindow("TextEditor")
	edit := widget.NewEntry()
	edit.MultiLine = true
	container := widget.NewScrollContainer(edit)
	inf := widget.NewLabel("information bar")


	newFile := func() {
		dialog.ShowConfirm("Alert", "Create New Document", func(f bool) {
			if f {
				edit.SetText("")
				inf.SetText("create new document")
			}
		}, window)
	}

	openFile := func() {
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Open file name", "OK", "Cancel", f, func(b bool) {
			if b {
				filename := f.Text + ".txt"
				byteArray, err := ioutil.ReadFile(filename)
				if err != nil {
					dialog.ShowError(err, window)
					return
				}
				edit.SetText(string(byteArray))
				inf.SetText("Open from file '" + filename + "'.")
			}
		}, window)
	}

	saveFile := func() {
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Save file name", "OK", "Cancel", f, func(b bool) {
			if b {
				filename := f.Text + ".txt"
				err := ioutil.WriteFile(filename, []byte(edit.Text), os.ModePerm)
				if err != nil {
					dialog.ShowError(err, window)
					return
				}
				inf.SetText("Save to file '" + filename + "'.")
			}
		}, window)
	}


	tf := true
	changeTheme := func() {
		if tf {
			myApp.Settings().SetTheme(theme.LightTheme())
			inf.SetText("change to Light Theme")
		}	else {
			myApp.Settings().SetTheme(theme.DarkTheme())
			inf.SetText("change to Dark Theme")
		}
		tf = !tf
	}

	quit := func() { 
		dialog.ShowConfirm("Alert", "Quit application?", func(b bool) {
			if b {
				myApp.Quit()
			}
		}, window)
	}

	createMenubar := func() *fyne.MainMenu {
		return fyne.NewMainMenu(
			fyne.NewMenu("File",
				fyne.NewMenuItem("New File", func() { newFile() }),
				fyne.NewMenuItem("Open File", func() { openFile() }),
				fyne.NewMenuItem("Save", func() { saveFile() }),
				fyne.NewMenuItem("Save", func() { changeTheme() }),
				fyne.NewMenuItem("Save", func() { quit() }),
			),
			fyne.NewMenu("Edit", 
				fyne.NewMenuItem("Cut", func() { 
					edit.TypedShortcut(&fyne.ShortcutCut{ 
						Clipboard: window.Clipboard(),
					})
					inf.SetText("Cut text.")
				}),
				fyne.NewMenuItem("Copy", func() {
					edit.TypedShortcut(&fyne.ShortcutCopy{
						Clipboard: window.Clipboard(),
					})
					inf.SetText("Copy text.")
				}),
				fyne.NewMenuItem("Paste", func() {
					edit.TypedShortcut(&fyne.ShortcutCopy{
						Clipboard: window.Clipboard(),
					})
					inf.SetText("Paste text.")
				}),
			),
		)
	}


	createToolbar := func() *widget.Toolbar {
		return widget.NewToolbar(
			widget.NewToolbarAction(
				theme.DocumentCreateIcon(), func() { newFile() },
			),
			widget.NewToolbarAction(
				theme.FolderOpenIcon(), func() { openFile() },
			),
			widget.NewToolbarAction(
				theme.DocumentSaveIcon(), func() { saveFile() },
			),
		)
	}

	menuBar := createMenubar()
	toolbar := createToolbar()

	window.SetMainMenu(menuBar)
	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(toolbar, inf, nil, nil), 
			toolbar, inf, container,
		),
	)

	window.Resize(fyne.NewSize(500, 500))
	window.ShowAndRun()
}