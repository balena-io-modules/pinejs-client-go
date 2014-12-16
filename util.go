package pinejs
import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
)

func toJsonReader(v interface{}) (io.Reader, error) {
	if v == nil {
		return nil, nil
	}

	if buf, err := json.Marshal(v); err != nil {
		return nil, err
	} else {
		return bytes.NewReader(buf), nil
	}
}

// Some functionality that is strangely lacking from simplejson...

type jsonNodeType int

const (
	jsonObject jsonNodeType = iota
	jsonArray
	jsonValue // Anything else.
)

func getJsonNodeType(j *simplejson.Json) jsonNodeType {
	// TODO: Reuse returned values.
	if _, err := j.Map(); err == nil {
		return jsonObject
	} else if _, err := j.Array(); err == nil {
		return jsonArray
	} else {
		return jsonValue
	}
}

func getJsonFieldNames(j *simplejson.Json) (ret []string) {
	if obj, err := j.Map(); err != nil {
		// Caller should have checked.
		panic(err)
	} else {
		for name, _ := range obj {
			ret = append(ret, name)
		}
	}

	return
}

func getJsonFields(j *simplejson.Json) map[string]*simplejson.Json {
	ret := make(map[string]*simplejson.Json)

	for _, name := range getJsonFieldNames(j) {
		ret[name] = j.Get(name)
	}

	return ret
}

func getJsonArray(j *simplejson.Json) (ret []*simplejson.Json) {
	if arr, err := j.Array(); err != nil {
		// Caller should have checked.
		panic(err)
	} else {
		// TODO: This sucks. Don't want to remarshal just to use returned data
		// though.
		for i := 0; i < len(arr); i++ {
			ret = append(ret, j.GetIndex(i))
		}
	}

	return
}
