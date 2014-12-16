package pinejs

import (
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// Retrieve resource name from input struct - if contains a pinejs tag, use
// that, otherwise use the lowercase of the struct name.
func resourceNameFromStruct(v interface{}) string {
	// Look for pinejs tag, use it if we find it.
	for _, f := range structs.Fields(v) {
		if name := f.Tag("pinejs"); name != "" {
			return name
		}
	}

	// Otherwise, we default to the name of the struct in lower case.
	return strings.ToLower(structs.Name(v))
}

// Unwinds pointers, slices, and slices of pointers, etc. until we get to a
// struct then we hand off to resourceNameFromStruct.
func resourceName(v interface{}) string {
	ty := reflect.TypeOf(v)

	switch ty.Kind() {
	case reflect.Struct:
		// Stay Calm and Carry on
	case reflect.Ptr, reflect.Slice:
		// Create new pointer to pointer/slice type.
		ptr := reflect.New(ty.Elem())

		// Deref the pointer and recurse on that value until we get to a struct.
		el := ptr.Elem().Interface()
		return resourceName(el)
	default:
		// TODO: I think this is probably impossible due to guards elsewhere.
		logAlert.Fatalf("attempted to get resourceName of %s", ty.Kind())
	}

	return resourceNameFromStruct(v)
}

// Retrieve Id field from interface.
func resourceId(v interface{}) int {
	return structs.New(v).Field("Id").Value().(int)
}
