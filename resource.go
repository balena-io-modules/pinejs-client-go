package pinejs

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
)

func resourceName(res interface{}) string {
	for _, f := range structs.Fields(res) {
		if name := f.Tag("pinejs"); name != "" {
			return name
		}
	}

	return strings.ToLower(structs.Name(res))
}

func resourceId(res interface{}) string {
	id := structs.New(res).Field("Id").Value().(int)
	return strconv.Itoa(id)
}

func resourceFromSlice(resSlice interface{}) interface{} {
	resType := reflect.TypeOf(resSlice).Elem().Elem()
	return reflect.New(resType).Interface()
}
