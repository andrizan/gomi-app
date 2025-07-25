package ui

import (
	"fmt"

	"gomi-app/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MangaReader struct {
    window      fyne.Window
    sidebar     *components.Sidebar
    mainContent *components.MainContent
    header      *components.Header
}

func NewMangaReader(window fyne.Window) *MangaReader {
    reader := &MangaReader{
        window:      window,
        sidebar:     components.NewSidebar(),
        mainContent: components.NewMainContent(),
        header:      components.NewHeader(),
    }

    // Set up callbacks
    reader.sidebar.SetOnComicSelected(func(comic string) {
        reader.updateChapterList()
    })

    reader.sidebar.SetOnChapterSelected(func(chapter string) {
        reader.loadChapterPages()
    })

    return reader
}

func (m *MangaReader) CreateUI() fyne.CanvasObject {
    // Create header
    // header := m.header.Create()

    // Create left sidebar with comic and chapter lists
    leftSidebar := m.sidebar.Create()

    // Create main content area for manga pages
    mainContent := m.mainContent.Create()

    // Combine left sidebar and main content
    splitContainer := container.NewHSplit(
        leftSidebar,
        mainContent,
    )
    splitContainer.SetOffset(0.25) // Set sidebar to 25% width

    // Combine header and main content
    mainContainer := container.NewBorder(
        nil, // header if needed
        nil,
        nil,
        nil,
        splitContainer,
    )

    return mainContainer
}

func (m *MangaReader) updateChapterList() {
    // Update chapter list based on selected comic
    fmt.Printf("Loading chapters for comic\n")
    m.sidebar.Refresh()
}

func (m *MangaReader) loadChapterPages() {
    // Load pages for selected chapter
    fmt.Printf("Loading pages for chapter\n")
    m.mainContent.LoadPages()
}
