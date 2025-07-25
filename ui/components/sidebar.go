package components

import (
	"gomi-app/ui/themekit"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Sidebar struct {
    comicList          *widget.List
    chapterList        *widget.List
    currentComic       string
    currentChapter     string
    comicSearchEntry   *widget.Entry
    chapterSearchEntry *widget.Entry
    onComicSelected    func(string)
    onChapterSelected  func(string)
}

func NewSidebar() *Sidebar {
    return &Sidebar{}
}

func (s *Sidebar) Create() fyne.CanvasObject {
    // Comic List Section
    comicLabel := widget.NewLabelWithStyle("Comic List", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

    s.comicSearchEntry = widget.NewEntry()
    s.comicSearchEntry.SetPlaceHolder("Search...")

    comicSearchBar := container.NewBorder(
        nil, nil,
        widget.NewIcon(theme.SearchIcon()),
        widget.NewButtonWithIcon("", theme.MoreVerticalIcon(), func() {}),
        s.comicSearchEntry,
    )

    // Sample comic data
    comics := []string{
        "Isekai Demo Bunan ni Ikitai Shoukougun",
        "Kage no Jitsuryokusha ni Naritakute",
    }

    s.comicList = widget.NewList(
        func() int { return len(comics) },
        func() fyne.CanvasObject {
            return widget.NewLabel("Comic Name")
        },
        func(i widget.ListItemID, o fyne.CanvasObject) {
            o.(*widget.Label).SetText(comics[i])
        },
    )
    s.comicList.OnSelected = func(id widget.ListItemID) {
        s.currentComic = comics[id]
        if s.onComicSelected != nil {
            s.onComicSelected(comics[id])
        }
        s.updateChapterList()
    }

    borderComic := canvas.NewRectangle(themekit.ThemeColor(theme.ColorNameInputBorder))
    borderComic.StrokeWidth = 1
    borderComic.CornerRadius = 5
    borderComic.StrokeColor = themekit.ThemeColor(theme.ColorNameInputBorder)
    borderComic.FillColor = themekit.ThemeColor(theme.ColorNameInputBackground)

    comicList := container.NewStack(borderComic, s.comicList)
    paddingComic := container.NewPadded(comicList)

    comicSection := container.NewBorder(
        container.NewVBox(
            comicLabel,
            comicSearchBar,
        ),
        widget.NewLabel("Comics: 0/0"),
        nil,
        nil,
        container.NewScroll(paddingComic),
    )

    // Chapter List Section
    chapterLabel := widget.NewLabelWithStyle("Chapter List", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

    s.chapterSearchEntry = widget.NewEntry()
    s.chapterSearchEntry.SetPlaceHolder("Search...")

    chapterSearchBar := container.NewBorder(
        nil, nil,
        widget.NewIcon(theme.SearchIcon()),
        widget.NewButtonWithIcon("", theme.MoreVerticalIcon(), func() {}),
        s.chapterSearchEntry,
    )

    // Sample chapter data
    chapters := []string{
        "Chapter 01",
        "Chapter 02",
    }

    s.chapterList = widget.NewList(
        func() int { return len(chapters) },
        func() fyne.CanvasObject {
            return widget.NewLabel("Chapter")
        },
        func(i widget.ListItemID, o fyne.CanvasObject) {
            o.(*widget.Label).SetText(chapters[i])
        },
    )
    s.chapterList.OnSelected = func(id widget.ListItemID) {
        s.currentChapter = chapters[id]
        if s.onChapterSelected != nil {
            s.onChapterSelected(chapters[id])
        }
    }

    borderChapter := canvas.NewRectangle(themekit.ThemeColor(theme.ColorNameInputBorder))
    borderChapter.StrokeWidth = 1
    borderChapter.CornerRadius = 5
    borderChapter.StrokeColor = themekit.ThemeColor(theme.ColorNameInputBorder)
    borderChapter.FillColor = themekit.ThemeColor(theme.ColorNameInputBackground)

    chapterList := container.NewStack(borderChapter, s.chapterList)
    paddingChapter := container.NewPadded(chapterList)

    chapterSection := container.NewBorder(
        container.NewVBox(
            chapterLabel,
            chapterSearchBar,
        ),
        widget.NewLabel("Chapters: 0/0"),
        nil,
        nil,
        container.NewScroll(paddingChapter),
    )

    // Combine both sections with separator
    leftSidebar := container.NewVSplit(
        comicSection,
        chapterSection,
    )
    leftSidebar.SetOffset(0.5)

    // Add background
    bg := canvas.NewRectangle(themekit.ThemeColor(theme.ColorNameBackground))
    return container.NewStack(bg, leftSidebar)
}

func (s *Sidebar) SetOnComicSelected(callback func(string)) {
    s.onComicSelected = callback
}

func (s *Sidebar) SetOnChapterSelected(callback func(string)) {
    s.onChapterSelected = callback
}

func (s *Sidebar) updateChapterList() {
    // Update chapter list based on selected comic
    // This would typically load actual chapter data
}

func (s *Sidebar) Refresh() {
    if s.comicList != nil {
        s.comicList.Refresh()
    }
    if s.chapterList != nil {
        s.chapterList.Refresh()
    }
}
