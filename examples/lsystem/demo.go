package main

import (
	rv "renderview"
	"renderview/driver"
)

func main() {
	m := rv.NewBasicRenderModel()
	m.AddParameters(
		rv.SetHints(rv.HINT_SIDEBAR,
			rv.NewFloat64RP("left", -30),
			rv.NewFloat64RP("top", -30),
			rv.NewFloat64RP("right", 30),
			rv.NewFloat64RP("bottom", 30),
			rv.NewIntRP("width", 100),
			rv.NewIntRP("height", 100),
			rv.NewIntRP("options", rv.OPT_AUTO_ZOOM))...)
	m.AddParameters(
		rv.SetHints(rv.HINT_HIDE,
			rv.NewStringRP("LSystemResult", ""),
		)...)
	lsystemRP := rv.NewStringRP("lsystem", "FX\nX=X+YF+\nY=-FX-Y\n")
	m.AddParameters(rv.SetHints(rv.HINT_FULLTEXT, lsystemRP)...)
	m.AddParameters(
		rv.SetHints(rv.HINT_FOOTER,
			rv.NewFloat64RP("angle", 90),
			rv.NewIntRP("depth", 5))...)
	c := rv.NewChangeMonitor()
	c.AddParameters(m.Params[8], m.Params[10]) // lsystem, depth
	m.InnerRender = func() {
		m.Img = RenderLSystemModel(m, c)
		m.NeedsRender = true
	}
	m.NeedsRender = true
	driver.Main(m)
}
