package widgets

import (
	"image"
	"supersonic/ui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Shows the current album art, track name, artist name, and album name
// for the currently playing track. Placed into the left side of the BottomPanel.
type NowPlayingCard struct {
	widget.BaseWidget

	trackName  *widget.Label
	artistName *CustomHyperlink
	albumName  *CustomHyperlink
	cover      *canvas.Image

	c fyne.CanvasObject
}

func NewNowPlayingCard() *NowPlayingCard {
	n := &NowPlayingCard{
		trackName:  widget.NewLabel(""),
		artistName: NewCustomHyperlink(),
		albumName:  NewCustomHyperlink(),
		cover:      &canvas.Image{},
	}
	n.ExtendBaseWidget(n)
	n.artistName.Hidden = true
	n.albumName.Hidden = true
	n.trackName.Wrapping = fyne.TextTruncate
	n.trackName.TextStyle = fyne.TextStyle{Bold: true}
	n.cover.SetMinSize(fyne.NewSize(85, 85))
	n.cover.FillMode = canvas.ImageFillContain

	n.c = container.New(&layouts.MaxPadLayout{PadLeft: -5},
		container.NewBorder(nil, nil, n.cover, nil,
			container.New(&layouts.MaxPadLayout{PadBottom: -3},
				container.New(&layouts.VboxCustomPadding{ExtraPad: -13}, n.trackName, n.artistName, n.albumName))),
	)
	return n
}

func (n *NowPlayingCard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(n.c)
}

func (n *NowPlayingCard) Update(track, artist, album string, cover image.Image) {
	n.trackName.Text = track
	n.artistName.SetText(artist)
	n.artistName.Hidden = artist == ""
	n.albumName.SetText(album)
	n.albumName.Hidden = album == ""
	n.cover.Image = cover
	n.c.Refresh()
}

func (n *NowPlayingCard) OnArtistNameTapped(f func()) {
	n.artistName.OnTapped = f
}

func (n *NowPlayingCard) OnAlbumNameTapped(f func()) {
	n.albumName.OnTapped = f
}
