package pinejs

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// Retrieve resource name from input struct - if contains a pinejs tag, use
// that, otherwise use the lowercase of the struct name.
func resourceNameFromStruct(v interface{}) string {
	// Only called from a function that asserts input is a struct.

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
func resourceName(v interface{}) (string, error) {
	ty := reflect.TypeOf(v)

	switch ty.Kind() {
	case reflect.Struct:
		return resourceNameFromStruct(v), nil
	case reflect.Ptr, reflect.Slice:
		// Create new pointer to pointer/slice type.
		ptr := reflect.New(ty.Elem())

		// Deref the pointer and recurse on that value until we get to a struct.
		el := ptr.Elem().Interface()
		return resourceName(el)
	}

	return "", fmt.Errorf("tried to retrieve resource name from non-struct %s",
		ty.Kind())
}

// Retrieve Id field from interface.
func resourceId(v interface{}) (ret int64, err error) {
	vl := reflect.ValueOf(v)

	for vl.Kind() == reflect.Ptr {
		vl = vl.Elem()
	}

	if vl.Kind() != reflect.Struct {
		return 0, fmt.Errorf("tried to retrieve resource id from non-struct %s",
			vl.Kind())
	}

	if field := vl.FieldByName("Id"); !field.IsValid() {
		err = errors.New("Id field not present")
	} else if !field.Type().ConvertibleTo(reflect.TypeOf(ret)) {
		err = fmt.Errorf("Id not convertible to %s", reflect.TypeOf(ret))
	} else {
		ret = field.Int()
	}

	return
}
