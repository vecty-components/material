package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

type ComponentCatalogPanel struct {
	designLink  string
	description string
	demos       vecty.ComponentOrHTML
	docsLink    string
	hero        vecty.ComponentOrHTML
	sourceLink  string
	title       string
	vecty.Core
}

func NewComponentCatalogPanel(
	designLink string,
	description string,
	docsLink string,
	sourceLink string,
	title string,
	hero vecty.ComponentOrHTML,
	demos vecty.ComponentOrHTML,
) *ComponentCatalogPanel {
	return &ComponentCatalogPanel{
		designLink:  designLink,
		description: description,
		demos:       demos,
		docsLink:    docsLink,
		hero:        hero,
		sourceLink:  sourceLink,
		title:       title,
	}
}

func (cp *ComponentCatalogPanel) Render() vecty.ComponentOrHTML {

	heroComponent := cp.hero

	return elem.Section(
		vecty.Markup(
			vecty.Class("component-catalog-panel"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("component-catalog-panel__hero-area"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("component-catalog-panel__header"),
				),
				elem.Heading1(
					vecty.Markup(
						vecty.Class(
							"component-catalog-panel__header-elements", "mdc-typography--headline3",
						),
					),
					vecty.Text(cp.title),
				),
				elem.Paragraph(
					vecty.Markup(
						vecty.Class(
							"component-catalog-panel__header-elements", "mdc-typography--body1",
						),
					),
					vecty.Text(cp.description),
				),
				heroComponent,
				/*
				   <HeroOptionsComponent
				     className=' component-catalog-panel__header-elements component-catalog-panel__header__hero-options'
				     config={localConfig}
				     {...this.props}
				   />
				*/
			),
		),
		elem.Heading2(
			vecty.Markup(
				vecty.Class("demo-title", "mdc-typography--headline6"),
			),
			vecty.Text("Resources"),
		),
		elem.UnorderedList(
			vecty.Markup(
				vecty.Class("component-catalog-resources"),
			),

			cp.renderResource("Material Design Guidelines", "/assets/images/ic_material_design_24px.svg", cp.designLink),
			cp.renderResource("Documentation", "/assets/images/ic_drive_document_24px.svg", cp.docsLink),
			cp.renderResource("Source Code", "/assets/images/ic_code_24px.svg", cp.sourceLink),
		),

		elem.Heading2(
			vecty.Markup(
				vecty.Class("demo-title", "mdc-typography--headline6"),
			),
			vecty.Text("Demos"),
		),

		cp.demos,
	)
}

func (cp *ComponentCatalogPanel) renderResource(title, imageSource, url string) vecty.ComponentOrHTML {
	if url == "" {
		return elem.Span()
	}

	return elem.ListItem(
		elem.Anchor(
			vecty.Markup(
				prop.Href(url),
				vecty.Class("mdc-list-item"),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("mdc-list-item__graphic", "resources-graphic"),
				),
				elem.Image(
					vecty.Markup(
						prop.Src(imageSource),
						vecty.Class("resources-graphic"),
						prop.Alt(title+" icon"),
					),
				),
			),
		),
	)
}
