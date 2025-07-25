package components

import (
	"fmt"
	"image/color"

	"gomi-app/ui/themekit"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MainContent struct {
    pageContainer *fyne.Container
}

func NewMainContent() *MainContent {
    return &MainContent{}
}

func (mc *MainContent) Create() fyne.CanvasObject {
    // Create page container with grid layout for manga pages
    mc.pageContainer = container.NewGridWithColumns(2)

    // Add sample pages (placeholders)
    mc.addSamplePages()

    // Wrap in scroll container
    scrollContainer := container.NewScroll(mc.pageContainer)

    // Add dark background like manga reader
    bg := canvas.NewRectangle(themekit.ThemeColor(theme.ColorNameBackground))

    // Create margin using spacer containers
    leftMargin := canvas.NewRectangle(color.Transparent)
    leftMargin.SetMinSize(fyne.NewSize(20, 0))
    rightMargin := canvas.NewRectangle(color.Transparent)
    rightMargin.SetMinSize(fyne.NewSize(20, 0))
    topMargin := canvas.NewRectangle(color.Transparent)
    topMargin.SetMinSize(fyne.NewSize(0, 20))
    bottomMargin := canvas.NewRectangle(color.Transparent)
    bottomMargin.SetMinSize(fyne.NewSize(0, 20))

    // Wrap content with margins
    contentWithMargins := container.NewBorder(
        topMargin,    // top
        bottomMargin, // bottom
        leftMargin,   // left
        rightMargin,  // right
        scrollContainer,
    )

    return container.NewStack(
        bg,
        contentWithMargins,
    )
}

func (mc *MainContent) addSamplePages() {
    // Clear existing content
    if mc.pageContainer != nil {
        mc.pageContainer.Objects = nil

        // Add placeholder pages
        for i := 1; i <= 4; i++ {
            page := mc.createPagePlaceholder(fmt.Sprintf("Page %d", i))
            mc.pageContainer.Add(page)
        }
    }
}

type HoverOverlay struct {
    widget.BaseWidget
    overlay *canvas.Rectangle
    label   *widget.Label
}

func NewHoverOverlay(overlay *canvas.Rectangle, label *widget.Label) *HoverOverlay {
    h := &HoverOverlay{
        overlay: overlay,
        label:   label,
    }
    h.ExtendBaseWidget(h)
    return h
}

func (h *HoverOverlay) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(container.NewWithoutLayout())
}

func (h *HoverOverlay) MouseIn(_ *desktop.MouseEvent) {
    h.overlay.Show()
    h.label.Show()
}

func (h *HoverOverlay) MouseOut() {
    h.overlay.Hide()
    h.label.Hide()
}

func (h *HoverOverlay) MouseMoved(_ *desktop.MouseEvent) {}

func (mc *MainContent) createPagePlaceholder(text string) fyne.CanvasObject {
    // Background
    pageRect := canvas.NewRectangle(themekit.ThemeColor(theme.ColorNameInputBackground))
    pageRect.SetMinSize(fyne.NewSize(400, 600))

    // Label (awalnya tidak terlihat)
    label := widget.NewLabelWithStyle(text, fyne.TextAlignCenter, fyne.TextStyle{})
    label.Hide()

    // Border
    border := canvas.NewRectangle(color.Transparent)
    border.StrokeColor = themekit.ThemeColor(theme.ColorNameInputBorder)
    border.StrokeWidth = 1

    // Overlay hitam transparan untuk efek gelap saat hover
    overlay := canvas.NewRectangle(color.NRGBA{0, 0, 0, 50}) // transparan hitam
    overlay.Resize(fyne.NewSize(400, 600))
    overlay.Hide()

    // Container utama
    content := container.NewStack(
        pageRect,
        border,
        overlay,
        container.NewCenter(label),
    )

    // Hover tracker
    hover := NewHoverOverlay(overlay, label)

    return container.NewStack(content, hover)
}

func (mc *MainContent) LoadPages() {
    mc.addSamplePages()
}

func (mc *MainContent) GetContainer() *fyne.Container {
    return mc.pageContainer
}
