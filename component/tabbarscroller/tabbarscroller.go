// The tabbarscroller package implements a material tabbarscroller component.
//
// See: https://material.io/components/web/catalog/tabs/
package tabbarscroller // import "agamigo.io/material/component/tabbarscroller"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div id="my-mdc-tab-bar-scroller" class="mdc-tab-bar-scroller">
  <div class="mdc-tab-bar-scroller__indicator mdc-tab-bar-scroller__indicator--back">
    <a class="mdc-tab-bar-scroller__indicator__inner material-icons" href="#" aria-label="scroll back button">
      navigate_before
    </a>
  </div>
  <div class="mdc-tab-bar-scroller__scroll-frame">
    <nav id="my-scrollable-tab-bar" class="mdc-tab-bar mdc-tab-bar-scroller__scroll-frame__tabs">
      <a class="mdc-tab mdc-tab--active" href="#one">Item One</a>
      <a class="mdc-tab" href="#two">Item Two</a>
      <a class="mdc-tab" href="#three">Item Three</a>
      <a class="mdc-tab" href="#four">Item Four</a>
      <a class="mdc-tab" href="#five">Item Five</a>
      <a class="mdc-tab" href="#six">Item Six</a>
      <a class="mdc-tab" href="#seven">Item Seven</a>
      <a class="mdc-tab" href="#eight">Item Eight</a>
      <a class="mdc-tab" href="#nine">Item Nine</a>
      <span class="mdc-tab-bar__indicator"></span>
    </nav>
  </div>
  <div class="mdc-tab-bar-scroller__indicator mdc-tab-bar-scroller__indicator--forward">
    <a class="mdc-tab-bar-scroller__indicator__inner material-icons" href="#" aria-label="scroll forward button">
      navigate_next
    </a>
  </div>
</div>`
)

// TBS is the interface for a material tabbarscroller component.
type TBS interface {
	component.C
}

// tabbarscroller is the internal implementation of TBS made available publicly via
// New().
type tabBarScroller struct {
	component.C
	html string
}

// New creates a material tabbarscroller component that implement the TBS interface.
// It is a wrapper around component.New.
func New() (c TBS, err error) {
	newTBS, err := component.New(component.TabBarScroller)
	if err != nil {
		return nil, err
	}
	return &tabBarScroller{newTBS, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (tbs *tabBarScroller) HTML() string {
	return tbs.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (tbs *tabBarScroller) SetHTML(html string) {
	tbs.html = html
}
