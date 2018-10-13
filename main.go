package main

import (
	"flag"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/ottojo/blnk"
	"github.com/ottojo/blnk/color"
	"log"
	"time"
)

var filename = flag.String("c", "/home/jonas/clients.json", "blnk System config file")
var animationSpeed = flag.Int("speed", 100, "Animation speed")
var framesPerSecond = flag.Int("fps", 20, "Framerate")

func main() {
	flag.Parse()
	system, err := blnk.LoadBlnkSystemFromFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	dps := *animationSpeed
	fps := *framesPerSecond
	for true {
		for _, c := range system.Clients {
			h := i
			step := 360.0 / len(c.Strip.NeoPixels)
			for i := range c.Strip.NeoPixels {
				col := colorful.Hsv(float64(h), 1, 1)
				r, g, b := col.LinearRgb()
				h = (h + step) % 360
				c.Strip.NeoPixels[i].SetColor(
					color.Color8bit{R: byte(r * 255), G: byte(g * 255), B: byte(b * 255)})
			}
			fmt.Println(".")
			c.Commit()
		}
		i += dps / fps
		time.Sleep(time.Duration(1e9 / fps))
	}
}
