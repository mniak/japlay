package sdl

import (
	"image/color"

	"github.com/mniak/japlayer/log"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type RenderData struct {
	Background *sdl.Texture
	Text       []string
	TextColor  color.Color
}

func (data *RenderData) Patch(new RenderData) {
	if new.Background != nil {
		if data.Background != nil {
			if err := data.Background.Destroy(); err != nil {
				log.Error(err, "faild to destroy background while replacing it")
			}
		}
		data.Background = new.Background
	}
	if new.Text != nil {
		data.Text = new.Text
	}
	if new.TextColor != nil {
		data.TextColor = new.TextColor
	}
}

func (data *RenderData) Destroy() error {
	if data.Background != nil {
		return data.Background.Destroy()
	}
	return nil
}

type RenderContext struct {
	renderer *sdl.Renderer
	font     *ttf.Font
	data     RenderData
}

func (rc *RenderContext) Render() error {
	err := rc.renderer.SetDrawColor(0, 0, 0, 255)
	if err != nil {
		return err
	}
	err = rc.renderer.Clear()
	if err != nil {
		return err
	}

	if rc.data.Background != nil {
		err = rc.renderer.Copy(rc.data.Background, nil, nil)
		if err != nil {
			return err
		}
	}

	width, height, err := rc.renderer.GetOutputSize()
	if err != nil {
		return err
	}
	centerX := float32(width) / 2
	centerY := float32(height) / 2
	var totalRect sdl.Rect

	type RenderLine struct {
		Texture *sdl.Texture
		Rect    sdl.Rect
		Target  sdl.Rect
	}

	// Pre-render texts and calculate position
	renderLines := make([]RenderLine, len(rc.data.Text))
	for lineIndex, line := range rc.data.Text {

		// line = adjustCase(line, rc.params.LetterCase)

		lineCenterY := centerY + float32((lineIndex*2-len(rc.data.Text))*rc.font.LineSkip())/2 + 60
		textSurface, err := rc.font.RenderUTF8Blended(line, Color(rc.data.TextColor))
		if err != nil {
			return err
		}
		textWidth := textSurface.W
		textHeight := textSurface.H
		renderLines[lineIndex].Rect = textSurface.ClipRect

		textTexture, err := rc.renderer.CreateTextureFromSurface(textSurface)
		textSurface.Free()
		if err != nil {
			return err
		}
		renderLines[lineIndex].Texture = textTexture

		targetRect := sdl.Rect{
			W: textWidth,
			H: textHeight,
			X: int32(centerX - float32(textWidth)/2),
			Y: int32(lineCenterY - float32(textHeight)/2),
		}

		renderLines[lineIndex].Target = targetRect
	}

	// Box
	for idx, line := range renderLines {
		if idx == 0 {
			totalRect = line.Target
		} else {
			totalRect = expandRect(totalRect, line.Target)
		}
	}
	totalRect = growRect(totalRect, 20, 30, 10, 30)
	err = rc.renderer.SetDrawColor(0, 0, 0, 200)
	if err != nil {
		return err
	}
	rc.renderer.FillRect(&totalRect)

	// Text
	for _, renderline := range renderLines {
		err = rc.renderer.Copy(renderline.Texture, &renderline.Rect, &renderline.Target)
		if err != nil {
			return err
		}
	}

	rc.renderer.Present()
	return nil
}
