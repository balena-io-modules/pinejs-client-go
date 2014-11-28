package pinejs

import (
	"fmt"

	"github.com/fatih/structs"
)

func normalizeField(data interface{}, s *structs.Struct) interface{} {
	return nil
}

func normalize(data interface{}, s *structs.Struct) interface{} {
	switch vv := data.(type) {
	case []interface{}:
		fmt.Println("data was an array")
		for i, v := range vv {
			vv[i] = normalize(v, s)
		}
	case map[string]interface{}:
		fmt.Println("data was an object")
		fields := s.Map()
		for k, v := range vv {
			vv[k] = normalize(v, fields[k].(*structs.Struct))
		}
	}

	return nil
}
