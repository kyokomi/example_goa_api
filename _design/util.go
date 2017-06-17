package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

func AttributeWithReturnName(name string, args ...interface{}) string {
	Attribute(name, args...)
	return name
}
