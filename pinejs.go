package pinejs

import (
	"errors"
)

type Client struct {
	APIKey   string
	Endpoint string
}

func (a *Client) Get(res interface{}) error {
	path := resourceName(res) + "(" + resourceId(res) + ")"

	return a.request("GET", path, nil, nil, &[]interface{}{res})
}

func (a *Client) List(resSlice interface{}) error {
	path := resourceName(resourceFromSlice(resSlice))

	return a.request("GET", path, nil, nil, resSlice)
}

func (a *Client) Create(res interface{}) error {
	// Should POST
	return errors.New("Not implemented")
}

func (a *Client) Update(res interface{}) error {
	// Should PUT
	return errors.New("Not implemented")
}

func (a *Client) Patch(res interface{}) error {
	// Should PATCH
	return errors.New("Not implemented")
}
func (a *Client) Delete(res interface{}) error {
	// Should DELETE
	return errors.New("Not implemented")
}
