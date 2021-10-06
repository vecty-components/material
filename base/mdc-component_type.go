package base // import "github.com/vecty-material/material/base"

// ComponentType is a specific component type, as implemented by the
// material-components-web library.
//
// See: https://material.io/components/web/catalog/
type ComponentType struct {
	//  MDCClassName represents the name of the MDC class of the component. For
	//  example, a form-field is "MDCFormField".
	MDCClassName string

	// MDCCamelCaseName is the lower camel case version of an MDC component.
	// When using the all-in-one distribution of the MDC library, it is the
	// name of the object that holds the MDCComponent/MDCFoundation etc. For
	// example in "mdc.formField.MDCFormField" the MDCamelCaseName is
	// "formField".
	MDCCamelCaseName string
}

// String returns the ComponentType's MDCClassName field.
func (n ComponentType) String() string {
	return n.MDCClassName
}
