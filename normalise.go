package pinejs

import (
	"fmt"

	"github.com/fatih/structs"
)

func normaliseField(data interface{}, s *structs.Struct) interface{} {
	return nil
}

func normalise(data interface{}, s *structs.Struct) interface{} {
	switch vv := data.(type) {
	case []interface{}:
		fmt.Println("data was an array")
		for i, v := range vv {
			vv[i] = normalise(v, s)
		}
	case map[string]interface{}:
		fmt.Println("data was an object")
		fields := s.Map()
		for k, v := range vv {
			vv[k] = normalise(v, fields[k].(*structs.Struct))
		}
	}

	return nil
}
