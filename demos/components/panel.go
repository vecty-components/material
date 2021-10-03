package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/ul"
)

type ComponentCatalogPanel struct {
	designLink  string
	description string
	demos       vecty.ComponentOrHTML
	docsLink    string
	hero        *HeroComponent
	sourceLink  string
	title       string
	vecty.Core
}

func NewComponentPage(
	title string,
	description string,
	designLink string,
	docsLink string,
	sourceLink string,
	hero *HeroComponent,
	demos vecty.ComponentOrHTML,
) *ComponentPage {
	return &ComponentPage{
		panel: &ComponentCatalogPanel{
			designLink:  designLink,
			description: description,
			demos:       demos,
			docsLink:    docsLink,
			hero:        hero,
			sourceLink:  sourceLink,
			title:       title,
		},
	}
}

func (cp *ComponentCatalogPanel) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/ComponentCatalogPanel.css")

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
				cp.hero,
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

	return &ul.Item{
		Graphic: elem.Span(
			vecty.Markup(
				vecty.Class("resources-graphic"),
			),
			elem.Image(
				vecty.Markup(
					prop.Src(imageSource),
					vecty.Class("resources-graphic"),
					prop.Alt(title+" icon"),
				),
			),
		),
		Primary: elem.Anchor(
			vecty.Markup(
				prop.Href(url),
			),
			vecty.Text(title),
		),
	}
}
