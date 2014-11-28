package pinejs

import (
	"testing"

	"bitbucket.org/rulemotion/pinejs-client-go/resin"
)

func TestApi(t *testing.T) {
	api := NewClient("https://alpha.resin.io/ewa", "bananasbananas")

	app := resin.Application{Id: 338}
	api.Get(&app)
}
