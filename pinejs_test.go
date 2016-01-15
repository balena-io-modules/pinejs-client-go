package pinejs

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/resin-io/pinejs-client-go/resin"
)

func init() {
	ShowDebugOutput()
}

func TestGet(t *testing.T) {
	if api, err := mockServerClientFromPath("appTest.json"); err != nil {
		t.Fatal(err)
	} else {
		app := resin.Application{Id: 338}
		if err := api.Get(&app, NewQueryOptions(Expand, "device")...); err != nil {
			t.Fatal(err)
		}

		// Some selected fields.
		expectedDevices := []resin.Device{
			resin.Device{
				Id:   486,
				UUID: "567fd03a5002a4cc2c4f178e2446a390d95b49b180aa6871e0acdbe0a723df",
			},
			resin.Device{
				Id:   476,
				UUID: "1f9af371c7113159893314fd1b14a9d239241550751240b0a943f2e48315ab",
			},
			resin.Device{
				Id:   477,
				UUID: "a2f33072393dbc9a3309c111a9705225050252ec0a454651333e7d6d29fb36",
			},
		}

		if app.Id != 338 {
			t.Errorf("Application ID changed to %d", app.Id)
		}

		if app.Devices == nil {
			t.Error("nil devices")
		} else if len(expectedDevices) != len(app.Devices) {
			t.Errorf("expected %d devices, got %d", len(expectedDevices),
				len(app.Devices))
		} else {
			actuals := make(map[int]resin.Device)
			for _, d := range app.Devices {
				actuals[d.Id] = d
			}

			for _, expected := range expectedDevices {
				if actual, has := actuals[expected.Id]; !has {
					t.Errorf("Missing expected Device id %d", actual.Id)
				} else {
					if actual.UUID != expected.UUID {
						t.Errorf("Device %d: Expected UUID '%s', got '%s'",
							expected.Id, expected.UUID, actual.UUID)
					}
				}
			}
		}

		expectedApp := resin.Application{
			Id:      338,
			AppName: "beagle",
		}

		if app.Id != expectedApp.Id {
			t.Errorf("Expected app id %d, got %d", expectedApp.Id, app.Id)
		}

		if app.AppName != expectedApp.AppName {
			t.Errorf("Expected app name '%s', got '%s'",
				expectedApp.AppName, app.AppName)
		}

		expectedUser := resin.User{
			Id: 5,
		}

		for _, d := range app.Devices {
			if d.Application == nil {
				t.Errorf("Device %d: Missing application", d.Id)
			} else if d.Application.Id != expectedApp.Id {
				t.Errorf("Device %d: Expected application %d, got %d",
					d.Id, expectedApp.Id, d.Application.Id)
			}

			if expectedUser.Id != d.User.Id {
				t.Errorf("Device %d: Expected user id %d, got %d",
					d.Id, expectedUser.Id, d.User.Id)
			}
		}
	}
}

func TestList(t *testing.T) {
	if api, err := mockServerClientFromPath("listTest.json"); err != nil {
		t.Fatal(err)
	} else {
		var apps []resin.Application
		if err := api.List(&apps, NewQueryOptions(Expand, "device")...); err != nil {
			t.Fatal(err)
		}
	}
}

type IdError1 struct {
}

type IdError2 struct {
	Id string
}

type IdOmitted struct {
	Id int `json:"id,omitempty"`
}

type IdNotOmitted struct {
	Id int `json:"id"`
}

var (
	omittedErrors = []interface{}{
		IdError1{},
		&IdError1{},
		IdError2{},
		&IdError2{},
	}
	areOmitted = []interface{}{
		IdOmitted{},
		&IdOmitted{},
	}
	notOmitted = []interface{}{
		IdNotOmitted{},
		&IdNotOmitted{},
	}
)

func TestIdOmitted(t *testing.T) {
	for _, val := range omittedErrors {
		if _, err := isIdOmitted(val); err == nil {
			t.Errorf("%#v didn't raise idOmitted Error", val)
		}
	}

	for _, val := range areOmitted {
		if omitted, err := isIdOmitted(val); err != nil {
			t.Errorf("%#v raised unexpected error %s", val, err)
		} else if !omitted {
			t.Errorf("%#v unexpectedly marked not omitted", val)
		} else if by, err := json.Marshal(val); err != nil {
			t.Errorf("%s", err)
		} else if strings.Contains(string(by), "id") {
			t.Errorf("%#v marked omitted but present in encoding/json output", val)
		}
	}

	for _, val := range notOmitted {
		if omitted, err := isIdOmitted(val); err != nil {
			t.Errorf("%#v raised unexpected error %s", val, err)
		} else if omitted {
			t.Errorf("%#v unexpectedly marked omitted", val)
		} else if by, err := json.Marshal(val); err != nil {
			t.Errorf("%s", err)
		} else if !strings.Contains(string(by), "id") {
			t.Errorf("%#v marked not omitted but not present in encoding/json output", val)
		}
	}
}
