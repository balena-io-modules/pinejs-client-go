package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/resin-io/pinejs-client-go"
	"github.com/resin-io/pinejs-client-go/resin"
)

func main() {
	resinApi := &pinejs.Client{
		Endpoint: os.Getenv("API_ENDPOINT"),
		APIKey:   os.Getenv("API_KEY"),
	}

	var userApps []resin.Application
	var err error
	appId, _ := strconv.Atoi(os.Getenv("APP_ID"))
	userId, _ := strconv.Atoi(os.Getenv("USER_ID"))

	myApp := resin.Application{Id: appId}

	fmt.Println("Getting all apps")
	err = resinApi.List(&userApps)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got apps: %+v\n", userApps)
	}

	userAppMaps := make([]map[string]interface{}, 1)
	userAppMaps[0] = make(map[string]interface{})
	userAppMaps[0]["pinejs"] = "application"
	fmt.Println("Getting all apps with a slice of maps")
	err = resinApi.List(&userAppMaps)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got apps: %+v\n", userAppMaps)
	}

	fmt.Println("Getting one Application")

	resinApi.Get(&myApp)

	fmt.Printf("%+v\n", myApp)

	fmt.Println("Getting one Application with a map")
	appMap := map[string]interface{}{"pinejs": "application", "id": appId}
	resinApi.Get(&appMap)
	fmt.Printf("%+v\n", appMap)

	fmt.Println("Creating a device")
	dev := make(map[string]interface{})

	dev["pinejs"] = "device"
	dev["uuid"] = os.Getenv("TEST_UUID")
	dev["device_type"] = "raspberry-pi2"
	dev["application"] = appId
	dev["user"] = userId
	fmt.Printf("%+v", dev)

	err = resinApi.Create(&dev)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created device: %v", dev)
	}
}
