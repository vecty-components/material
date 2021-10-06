package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-components/material/demos/components"
	"github.com/vecty-components/material/typography"
	"github.com/vecty-components/material/ul"
)

func NewListPage() *components.ComponentPage {
	return components.NewComponentPage(
		"List",
		"Lists present multiple line items vertically as a single continuous element.",
		"https://material.io/go/design-lists",
		"https://material.io/components/web/catalog/lists/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-list",
		components.NewHeroComponent(&ListHero{}), &ListDemos{},
	)
}

type ListHero struct {
	vecty.Core
}

func (bh *ListHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),
		&ul.L{
			Root: vecty.Markup(
				vecty.Class("demo-list"),
			),
			Items: []vecty.ComponentOrHTML{
				&ul.Item{Primary: vecty.Text("Line Item")},
				&ul.Item{Primary: vecty.Text("Line Item")},
				&ul.Item{Primary: vecty.Text("Line Item")},
			},
		},
	)
}

type ListDemos struct {
	vecty.Core
}

func (bd *ListDemos) Render() vecty.ComponentOrHTML {

	return elem.Div(

		//		<ListVariant title='Single-Line'>
		//		<ListItem lineOne='Line item' tabIndex='0'/>
		//		<ListItem lineOne='Line item' />
		//		<ListItem lineOne='Line item' />
		//		</ListVariant>

		typography.Subtitle1(
			vecty.Text("Single-Line"),
		),

		&ul.L{
			Root: vecty.Markup(
				vecty.Class("demo-list"),
			),
			Items: []vecty.ComponentOrHTML{
				&ul.Item{Primary: vecty.Text("Line Item")},
				&ul.Item{Primary: vecty.Text("Line Item")},
				&ul.Item{Primary: vecty.Text("Line Item")},
			},
		},

	//		<ListVariant title='Two-Line' twoLines>
	//		<ListItem lineOne='Line item' lineTwo='Secondary text' tabIndex='0'/>
	//		<ListItem lineOne='Line item' lineTwo='Secondary text' />
	//		<ListItem lineOne='Line item' lineTwo='Secondary text' />
	//		</ListVariant>

	//		<ListVariant title='Leading Icon'>
	//		<ListItem lineOne='Line item' leadingIcon='wifi' tabIndex='0'/>
	//		<ListItem lineOne='Line item' leadingIcon='bluetooth' />
	//		<ListItem lineOne='Line item' leadingIcon='data_usage' />
	//		</ListVariant>
	//
	//		<ListVariant title='List with activated item'>
	//		<ListItem lineOne='Inbox' leadingIcon='inbox' />
	//		<ListItem activated lineOne='Star' leadingIcon='star' />
	//		<ListItem lineOne='Send' leadingIcon='send' />
	//		<ListItem lineOne='Drafts' leadingIcon='drafts' />
	//		</ListVariant>
	//
	//		<ListVariant title='List with shaped activated item' className='demo-list-item-shaped'>
	//		<ListItem lineOne='Inbox' leadingIcon='inbox' />
	//		<ListItem activated lineOne='Star' leadingIcon='star' />
	//		<ListItem lineOne='Send' leadingIcon='send' />
	//		<ListItem lineOne='Drafts' leadingIcon='drafts' />
	//		</ListVariant>
	//
	//		<ListVariant title='Trailing Icon'>
	//		<ListItem lineOne='Line item' trailingIcon='info' tabIndex='0'/>
	//		<ListItem lineOne='Line item' trailingIcon='info' />
	//		<ListItem lineOne='Line item' trailingIcon='info' />
	//		</ListVariant>
	//
	//		<ListVariant title='Two-Line with Leading and Trailing Icon and Divider' twoLines avatars>
	//		<ListItem lineOne='Dog Photos' lineTwo='9 Jan 2018' leadingIcon='folder' trailingIcon='info' tabIndex='0'/>
	//		<ListItem lineOne='Cat Photos' lineTwo='22 Dec 2017' leadingIcon='folder' trailingIcon='info' />
	//		<ListDivider />
	//		<ListItem lineOne='Potatoes' lineTwo='30 Noc 2017' leadingIcon='folder' trailingIcon='info' />
	//		<ListItem lineOne='Carrots' lineTwo='17 Oct 2017' leadingIcon='folder' trailingIcon='info' />
	//		</ListVariant>
	//
	//		<ListVariant title='List with Trailing Checkbox' avatars>
	//		<ListItem lineOne='Dog Photos' trailingCheckbox tabIndex='0'/>
	//		<ListItem lineOne='Cat Photos' trailingCheckbox />
	//		<ListDivider />
	//		<ListItem lineOne='Potatoes' trailingCheckbox />
	//		<ListItem lineOne='Carrots'  trailingCheckbox />
	//		</ListVariant>
	//
	//		<ListVariant title='List with Trailing Radio Buttons' avatars>
	//		<ListItem lineOne='Dog Photos' trailingRadio tabIndex='0'/>
	//		<ListItem lineOne='Cat Photos' trailingRadio />
	//		<ListDivider />
	//		<ListItem lineOne='Potatoes' trailingRadio />
	//		<ListItem lineOne='Carrots'  trailingRadio />
	//		</ListVariant>
	)

}
