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

	fmt.Println("Getting one Application")

	resinApi.Get(&myApp)

	fmt.Printf("%+v\n", myApp)

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
