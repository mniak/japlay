package sdl

import (
	"image/color"

	"github.com/mniak/japlayer/log"
	"github.com/samber/lo"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type displayAdapter struct {
	params  AdapterParams
	window  *sdl.Window
	context RenderContext
}

func DisplayAdapter(params AdapterParams) (ad *displayAdapter, err error) {
	ad = new(displayAdapter)
	defer func() {
		if err != nil {
			ad.Finish()
		}
	}()

	font, err := ttf.OpenFont(params.FontPath, int(params.FontSize))
	ad.context.font = font
	if err != nil {
		return
	}

	mode, err := sdl.GetCurrentDisplayMode(params.Display)
	if err != nil {
		return
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(
		mode.W, mode.H,
		sdl.WINDOW_HIDDEN|sdl.WINDOW_FULLSCREEN,
	)
	ad.window = window
	ad.context.renderer = renderer
	if err != nil {
		return
	}

	err = renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		return
	}

	err = ad.context.Render()
	if err != nil {
		return
	}
	ad.window.Show()
	return
}

func (ad *displayAdapter) Finish() {
	if ad.context.font != nil {
		ad.context.font.Close()
	}
	if err := ad.context.data.Destroy(); err != nil {
		log.Error(err, "failed to destroy render data")
	}
	if ad.context.renderer != nil {
		if err := ad.context.renderer.Destroy(); err != nil {
			log.Error(err, "failed to destroy background")
		}
	}
	if ad.window != nil {
		if err := ad.window.Destroy(); err != nil {
			log.Error(err, "failed to destroy window")
		}
	}
}

func (ad *displayAdapter) SetBackgroundImage(filename string) error {
	tex, err := LoadImageToTexture(ad.context.renderer, filename)
	if err != nil {
		return err
	}
	ad.context.data.Patch(RenderData{
		Background: tex,
	})
	return ad.context.Render()
}

func (ad *displayAdapter) ShowTitle(title string) error {
	title = AdjustCase(title, ad.params.LetterCase)

	ad.context.data.Patch(RenderData{
		Text:      []string{title},
		TextColor: color.RGBA{R: 0xef, G: 0xb4},
	})
	return ad.context.Render()
}

func (ad *displayAdapter) ShowVerse(lines ...string) error {
	lines = lo.Map(lines, func(line string, _ int) string {
		return AdjustCase(line, ad.params.LetterCase)
	})

	ad.context.data.Patch(RenderData{
		Text:      lines,
		TextColor: color.White,
	})
	return ad.context.Render()
}
