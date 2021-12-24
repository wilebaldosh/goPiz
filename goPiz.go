package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func fileNewDocument() {
	log.Println("New document func")
}

func fileSave() {
	log.Println("Save func")
}

func fileExit() {
	log.Println("Exit func")
	os.Exit(0)
}

func editCut() {
	log.Println("Cut func")
}

func editCopy() {
	log.Println("Copy func")
}

func editPaste() {
	log.Println("Paste func")
}

func editErase() {
	log.Println("Erase func")
}

func helpDisplay() {
	log.Println("Display help func")
}

func main() {
	a := app.New()
	w := a.NewWindow("Toolbar Widget")

	// Create the toolbar
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fileNewDocument()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			editCut()
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			editCopy()
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			editCopy()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			helpDisplay()
		}),
	)
	// Add the toolbar to container
	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	w.SetContent(content)

	// New menuItems for File
	fileNew := fyne.NewMenuItem("New", func() {
		fileNewDocument()
	})
	fileSave := fyne.NewMenuItem("Save", func() {
		fileSave()
	})
	fileExit := fyne.NewMenuItem("Exit", func() {
		fileExit()
	})
	// new File menu
	fileMenu := fyne.NewMenu("File", fileNew, fileSave, fileExit)

	// New menuItems for Edit
	mEditCut := fyne.NewMenuItem("Cut", func() {
		editCut()
	})
	mEditCopy := fyne.NewMenuItem("Copy", func() {
		editCopy()
	})
	mEditPaste := fyne.NewMenuItem("Paste", func() {
		editPaste()
	})
	mEditSeparator := fyne.NewMenuItemSeparator()
	mEditErase := fyne.NewMenuItem("Erase", func() {
		editErase()
	})
	// new Edit menu
	editMenu := fyne.NewMenu("Edit", mEditCut, mEditCopy, mEditPaste, mEditSeparator, mEditErase)

	// creatin new main menu
	menu := fyne.NewMainMenu(fileMenu, editMenu)
	// seting new menu
	w.SetMainMenu(menu)

	// Resize the window
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
