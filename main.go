package main

import (
	"fmt"

	"gomi-app/ui"
	"gomi-app/ui/themekit"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
    myApp := app.New()
    myApp.Settings().SetTheme(&themekit.MercuryTheme{})

    myWindow := myApp.NewWindow("Gomi Comic Reader")
    myWindow.Resize(fyne.NewSize(1400, 800))

    myWindow.SetMainMenu(fyne.NewMainMenu(
        fyne.NewMenu("File",
            fyne.NewMenuItem("Open", func() { fmt.Println("Open clicked") }),
            // fyne.NewMenuItem("Exit", func() { myWindow.Close() }),
        ),
        fyne.NewMenu("Edit",
            fyne.NewMenuItem("Preferences", func() { fmt.Println("Preferences clicked") }),
        ),
        fyne.NewMenu("Help",
            fyne.NewMenuItem("About", func() { fmt.Println("About clicked") }),
        ),
    ))

    reader := ui.NewMangaReader(myWindow)
    content := reader.CreateUI()
    myWindow.SetContent(content)

    myWindow.ShowAndRun()
}
