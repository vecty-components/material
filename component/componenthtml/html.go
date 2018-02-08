package componenthtml

import "agamigo.io/material/component"

func HTML(t component.Type) string {
	switch t {
	case component.Checkbox:
		return `
<div class="mdc-checkbox">
  <input class="mdc-checkbox__native-control" type="checkbox">
</div>`
	case component.Dialog:
		return `
<aside id="my-mdc-dialog"
  class="mdc-dialog"
  role="alertdialog"
  aria-labelledby="my-mdc-dialog-label"
  aria-describedby="my-mdc-dialog-description">
  <div class="mdc-dialog__surface">
    <header class="mdc-dialog__header">
      <h2 id="my-mdc-dialog-label" class="mdc-dialog__header__title">
        Dialog header
      </h2>
    </header>
    <section id="my-mdc-dialog-description" class="mdc-dialog__body">
      Dialog description
    </section>
    <footer class="mdc-dialog__footer">
      <button type="button" class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--cancel">Decline</button>
      <button type="button" class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--accept">Accept</button>
    </footer>
  </div>
  <div class="mdc-dialog__backdrop"></div>
</aside>`
	case component.PersistentDrawer:
		return `
<aside class="mdc-drawer mdc-drawer--persistent mdc-typography">
  <nav class="mdc-drawer__drawer">
    <header class="mdc-drawer__header">
      <div class="mdc-drawer__header-content">
        Header here
      </div>
    </header>
    <nav id="icon-with-text-demo" class="mdc-drawer__content mdc-list">
      <a class="mdc-list-item mdc-list-item--activated" href="#">
        <i class="material-icons mdc-list-item__graphic" aria-hidden="true">inbox</i>Inbox
      </a>
      <a class="mdc-list-item" href="#">
        <i class="material-icons mdc-list-item__graphic" aria-hidden="true">star</i>Star
      </a>
    </nav>
  </nav>
</aside>`
	case component.TemporaryDrawer:
		return `
<aside class="mdc-drawer mdc-drawer--temporary mdc-typography">
  <nav class="mdc-drawer__drawer">
    <header class="mdc-drawer__header">
      <div class="mdc-drawer__header-content">
        Header here
      </div>
    </header>
    <nav id="icon-with-text-demo" class="mdc-drawer__content mdc-list">
      <a class="mdc-list-item mdc-list-item--activated" href="#">
        <i class="material-icons mdc-list-item__graphic" aria-hidden="true">inbox</i>Inbox
      </a>
      <a class="mdc-list-item" href="#">
        <i class="material-icons mdc-list-item__graphic" aria-hidden="true">star</i>Star
      </a>
    </nav>
  </nav>
</aside>`
	case component.FormField:
		return `
<div class="mdc-form-field">
  <div class="mdc-checkbox">
    <input type="checkbox" id="my-checkbox"
    class="mdc-checkbox__native-control"/>
    </div>
  </div>
  <label for="my-checkbox" id="my-checkbox-label">This is my checkbox</label>
</div>`
	case component.GridList:
		return `
<div class="mdc-grid-list">
  <ul class="mdc-grid-list__tiles">
    <li class="mdc-grid-tile">
      <div class="mdc-grid-tile__primary">
        <img class="mdc-grid-tile__primary-content" src="my-image.jpg" />
      </div>
      <span class="mdc-grid-tile__secondary">
        <span class="mdc-grid-tile__title">Title</span>
      </span>
    </li>
  </ul>
</div>`
	case component.IconToggle:
		return `
<i class="mdc-icon-toggle material-icons" role="button" aria-pressed="false"
   aria-label="Add to favorites" tabindex="0"
   data-toggle-on='{"label": "Remove from favorites", "content": "favorite"}'
   data-toggle-off='{"label": "Add to favorites", "content": "favorite_border"}'>
  favorite_border
</i>`
	case component.LinearProgress:
		return `
<div role="progressbar" class="mdc-linear-progress">
  <div class="mdc-linear-progress__buffering-dots"></div>
  <div class="mdc-linear-progress__buffer"></div>
  <div class="mdc-linear-progress__bar mdc-linear-progress__primary-bar">
    <span class="mdc-linear-progress__bar-inner"></span>
  </div>
  <div class="mdc-linear-progress__bar mdc-linear-progress__secondary-bar">
    <span class="mdc-linear-progress__bar-inner"></span>
  </div>
</div>`
	case component.Menu:
		return `
<div class="mdc-menu-anchor">
  <button id="menu-button" class="mdc-button mdc-button--raised demo-button demo-button--normal">
    Show<span class="demo-button__normal-text"> Menu</span><span class="demo-button__long-text"> From Here Now!</span>
  </button>

  <div class="mdc-menu" style="position: absolute;" tabindex="-1" id="demo-menu">
    <ul class="mdc-menu__items mdc-list" role="menu" aria-hidden="true">
      <li class="mdc-list-item" role="menuitem" tabindex="0">Back</li>
      <li class="mdc-list-item" role="menuitem" tabindex="0">Forward</li>
      <li class="mdc-list-item" role="menuitem" tabindex="0">Reload</li>
      <li class="mdc-list-divider" role="separator"></li>
      <span class="demo-menu__long-items">
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 1</li>
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 2</li>
      </span>
      <span class="demo-menu__extra-long-items">
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 3</li>
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 4</li>
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 5</li>
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 6</li>
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 7</li>
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 8</li>
        <li class="mdc-list-item" role="menuitem" tabindex="0">Item 9</li>
      </span>
    </ul>
  </div>
</div>`
	case component.Radio:
		return `
<div class="mdc-radio">
  <input class="mdc-radio__native-control" type="radio" id="radio-1" name="radios" checked>
  <div class="mdc-radio__background">
    <div class="mdc-radio__outer-circle"></div>
    <div class="mdc-radio__inner-circle"></div>
  </div>
</div>
<label id="radio-1-label" for="radio-1">Radio 1</label>`
	case component.Ripple:
		return `
<div class="mdc-ripple-surface mdc-ripple-surface--primary my-surface" tabindex="0">
  Surface with a primary-colored ripple.
</div>`
	case component.Select:
		return `
<div class="mdc-select" role="listbox">
  <div class="mdc-select__surface" tabindex="0">
    <div class="mdc-select__label">Pick a Food Group</div>
    <div class="mdc-select__selected-text"></div>
    <div class="mdc-select__bottom-line"></div>
  </div>
  <div class="mdc-simple-menu mdc-select__menu">
    <ul class="mdc-list mdc-simple-menu__items">
      <li class="mdc-list-item" role="option" tabindex="0">
        Bread, Cereal, Rice, and Pasta
      </li>
      <li class="mdc-list-item" role="option" tabindex="0">
        Vegetables
      </li>
    </ul>
  </div>
</div>`
	case component.Slider:
		return `
<div class="mdc-slider" tabindex="0" role="slider"
     aria-valuemin="0" aria-valuemax="100" aria-valuenow="0"
     aria-label="Select Value">
  <div class="mdc-slider__track-container">
    <div class="mdc-slider__track"></div>
  </div>
  <div class="mdc-slider__thumb-container">
    <svg class="mdc-slider__thumb" width="21" height="21">
      <circle cx="10.5" cy="10.5" r="7.875"></circle>
    </svg>
    <div class="mdc-slider__focus-ring"></div>
  </div>
</div>`
	case component.Snackbar:
		return `
<div class="mdc-snackbar"
     aria-live="assertive"
     aria-atomic="true"
     aria-hidden="true">
  <div class="mdc-snackbar__text"></div>
  <div class="mdc-snackbar__action-wrapper">
    <button type="button" class="mdc-snackbar__action-button"></button>
  </div>
</div>`
	case component.Tab, component.TabBar:
		return `
<nav id="basic-tab-bar" class="mdc-tab-bar">
  <a class="mdc-tab mdc-tab--active" href="#one">Home</a>
  <a class="mdc-tab" href="#two">Merchandise</a>
  <a class="mdc-tab" href="#three">About Us</a>
  <span class="mdc-tab-bar__indicator"></span>
</nav>`
	case component.TabBarScroller:
		return `
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
	case component.TextField:
		return `
<div class="mdc-text-field">
  <input type="text" id="my-text-field" class="mdc-text-field__input">
  <label class="mdc-text-field__label" for="my-text-field">Hint text</label>
  <div class="mdc-text-field__bottom-line"></div>
</div>`
	case component.Toolbar:
		return `
<header class="mdc-toolbar">
  <div class="mdc-toolbar__row">
    <section class="mdc-toolbar__section mdc-toolbar__section--align-start">
      <a href="#" class="material-icons mdc-toolbar__menu-icon">menu</a>
      <span class="mdc-toolbar__title">Title</span>
    </section>
  </div>
</header>`
	}

	panic("Failed to get HTML for component type: " + t.String())
}
