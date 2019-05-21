package main

import (
	"image/png"
	"os"

	"github.com/tdewolff/canvas"
)

var dejaVuSerif canvas.Font

func main() {
	var err error
	dejaVuSerif, err = canvas.LoadFontFile("DejaVuSerif", canvas.Regular, "DejaVuSerif.woff")
	if err != nil {
		panic(err)
	}

	c := canvas.New(265, 100)
	Draw(c)

	pngFile, err := os.Create("text_example.png")
	if err != nil {
		panic(err)
	}
	defer pngFile.Close()

	img := c.WriteImage(144.0)
	err = png.Encode(pngFile, img)
	if err != nil {
		panic(err)
	}
}

func drawText(c *canvas.C, x, y float64, halign, valign canvas.TextAlign, indent float64) {
	face := dejaVuSerif.Face(6.0)
	phrase := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi egestas, augue eget blandit laoreet, dolor lorem interdum ante, quis consectetur lorem massa vitae nulla. Sed cursus tellus id venenatis suscipit. Nunc volutpat imperdiet ipsum vel varius. Pellentesque mattis viverra odio, ullamcorper iaculis massa tristique imperdiet. Aliquam posuere nisl tortor, in scelerisque elit eleifend sed."

	text := canvas.NewTextBox(face, phrase, 0.0, 35.0, halign, valign, indent)
	rect := text.Bounds()
	c.SetColor(canvas.WhiteSmoke)
	c.DrawPath(x, y, 0.0, rect.ToPath())
	c.SetColor(canvas.Black)
	c.DrawText(x, y, 0.0, text)
}

func Draw(c *canvas.C) {
	face := dejaVuSerif.Face(14.0)
	c.SetColor(canvas.Black)
	c.DrawText(132.5, 90.0, 0.0, canvas.NewTextBox(face, "Different horizontal and vertical alignments with indent", 0.0, 0.0, canvas.Center, canvas.Top, 0.0))

	drawText(c, 5.0, 80.0, canvas.Left, canvas.Top, 10.0)
	drawText(c, 70.0, 80.0, canvas.Center, canvas.Top, 10.0)
	drawText(c, 135.0, 80.0, canvas.Right, canvas.Top, 10.0)
	drawText(c, 200.0, 80.0, canvas.Justify, canvas.Top, 10.0)
	drawText(c, 5.0, 40.0, canvas.Left, canvas.Top, 10.0)
	drawText(c, 70.0, 40.0, canvas.Left, canvas.Center, 10.0)
	drawText(c, 135.0, 40.0, canvas.Left, canvas.Bottom, 10.0)
	drawText(c, 200.0, 40.0, canvas.Left, canvas.Justify, 10.0)
}
