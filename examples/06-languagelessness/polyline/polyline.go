package main

import (
	"github.com/eriner/ghostscad/lib/shapes"
	"github.com/eriner/ghostscad/sys"

	. "github.com/go-gl/mathgl/mgl64"
)

func main() {
	sys.Initialize()
	sys.SetFn(120)
	points := []Vec2{{1, 2}, {-5, -4}, {-5, 3}, {5, 5}}
	sys.RenderOne(shapes.NewPolyline(points, 1).SetRound(true).Build())
}
