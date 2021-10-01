package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type HeroComponent struct {
	h vecty.ComponentOrHTML
	vecty.Core
}

func NewHeroComponent(h vecty.ComponentOrHTML) *HeroComponent {
	return &HeroComponent{h: h}
}

func (hc *HeroComponent) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/HeroComponent.css")

	return elem.Div(
		vecty.Markup(
			vecty.Class("hero-component"),
		),
		hc.h,
		// &HeroTabs{},
	)
}

type HeroTabs struct {
	vecty.Core
}

func (ht *HeroTabs) Render() vecty.ComponentOrHTML {
	return &TabBar{}
}

type TabBar struct {
	vecty.Core
}

func (tb *TabBar) Render() vecty.ComponentOrHTML {
	return nil
}
