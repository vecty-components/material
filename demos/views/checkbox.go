package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-material/material/checkbox"
	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/typography"
)

func NewCheckboxPage() *components.ComponentPage {
	return components.NewComponentPage(
		"Checkbox",
		"Checkboxes allow the user to select multiple options from a set.",
		"https://material.io/go/design-checkboxes",
		"https://material.io/components/web/catalog/checkboxes/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-checkbox",
		components.NewHeroComponent(&CheckboxHero{}), &CheckboxDemos{},
	)
}

type CheckboxHero struct {
	vecty.Core
}

func (bh *CheckboxHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),

		&checkbox.CB{
			Root: vecty.Markup(
				vecty.Class("demo-checkbox"),
			),
		},

		&checkbox.CB{
			Root: vecty.Markup(
				vecty.Class("demo-checkbox"),
			),
		},

		//		<div className='mdc-checkbox mdc-checkbox--selected demo-checkbox' ref={this.initCheckbox}>
		//		<input type='checkbox'
		//			   defaultChecked={true}
		//			   className='mdc-checkbox__native-control'/>
		//		<div className='mdc-checkbox__background'>
		//		  <svg className='mdc-checkbox__checkmark'
		//			   viewBox='0 0 24 24'>
		//			<path className='mdc-checkbox__checkmark-path'
		//				  fill='none'
		//				  stroke='white'
		//				  d='M1.73,12.91 8.1,19.28 22.79,4.59'/>
		//		  </svg>
		//		  <div className='mdc-checkbox__mixedmark'></div>
		//		</div>
		//		<div className='mdc-checkbox__ripple'></div>
		//	  </div>
		//
		//	  <div className='mdc-checkbox demo-checkbox' ref={this.initCheckbox}>
		//		<input type='checkbox'
		//			   className='mdc-checkbox__native-control'/>
		//		<div className='mdc-checkbox__background'>
		//		  <svg className='mdc-checkbox__checkmark'
		//			   viewBox='0 0 24 24'>
		//			<path className='mdc-checkbox__checkmark-path'
		//				  fill='none'
		//				  stroke='white'
		//				  d='M1.73,12.91 8.1,19.28 22.79,4.59'/>
		//		  </svg>
		//		  <div className='mdc-checkbox__mixedmark'></div>
		//		</div>
		//		<div className='mdc-checkbox__ripple'></div>
		//	  </div>

	)
}

type CheckboxDemos struct {
	vecty.Core
}

func (bd *CheckboxDemos) Render() vecty.ComponentOrHTML {

	return elem.Div(
		typography.Subtitle1(
			vecty.Text("Unchecked"),
		),

		&checkbox.CB{
			Root: vecty.Markup(
				vecty.Class("demo-checkbox"),
			),
		},

		typography.Subtitle1(
			vecty.Text("Indeterminate"),
		),

		&checkbox.CB{
			Root: vecty.Markup(
				vecty.Class("demo-checkbox"),
			),
			Indeterminate: true,
		},

		typography.Subtitle1(
			vecty.Text("Checked"),
		),

		&checkbox.CB{
			Root: vecty.Markup(
				vecty.Class("demo-checkbox"),
			),
			Checked: true,
		},
	)

}
