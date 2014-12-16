package pinejs
import (
	"github.com/bitly/go-simplejson"
)
// Predicates

// These are used to transform input before attempting to unmarshal and vary
// depending on the caller.

type transformJSONFunc func(*simplejson.Json) *simplejson.Json

// Retrieve an array of data at field "d" from input json, or nil if doesn't
// exist.
func theD(j *simplejson.Json) *simplejson.Json {
	// Check we have data to get.
	if j == nil {
		return nil
	} else if _, has := j.CheckGet("d"); !has {
		return nil
	}

	return j.Get("d")
}

// Retrieve the first element of an array of data at field "d" from input json,
// or nil if doesn't exist.
func first(j *simplejson.Json) *simplejson.Json {
	if j = theD(j); j == nil {
		return nil
	} else if arr, err := j.Array(); err != nil || len(arr) == 0 {
		return nil
	}

	return j.GetIndex(0)
}

// Simply return the input json. Useful for the POST response which just echos
// the object back.
func self(j *simplejson.Json) *simplejson.Json {
	return j
}
