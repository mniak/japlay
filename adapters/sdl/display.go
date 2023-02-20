package sdl

import (
	"image/color"

	"github.com/samber/lo"
)

func (ad *sdlAdapter) SetBackgroundImage(filename string) error {
	tex, err := LoadImageToTexture(ad.context.renderer, filename)
	if err != nil {
		return err
	}
	ad.context.data.Patch(RenderData{
		Background: tex,
	})
	return ad.context.Render()
}

func (ad *sdlAdapter) ShowTitle(title string) error {
	title = AdjustCase(title, ad.params.LetterCase)

	ad.context.data.Patch(RenderData{
		Text:      []string{title},
		TextColor: color.RGBA{R: 0xef, G: 0xb4},
	})
	return ad.context.Render()
}

func (ad *sdlAdapter) ShowVerse(lines ...string) error {
	lines = lo.Map(lines, func(line string, _ int) string {
		return AdjustCase(line, ad.params.LetterCase)
	})

	ad.context.data.Patch(RenderData{
		Text:      lines,
		TextColor: color.White,
	})
	return ad.context.Render()
}
