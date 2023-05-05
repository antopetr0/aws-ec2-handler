package main

import (
	"fmt"

	"github.com/antopetr0/aws-ec2-handler/ec2"
)

func main() {
	manager, err := ec2.NewEC2Manager()
	if err != nil {
		panic(err)
	}

	instances, err := manager.ListInstances()
	if err != nil {
		panic(err)
	}

	for _, instance := range instances {
		fmt.Printf("Instance ID: %s, State: %s\n", *instance.InstanceId, *instance.State.Name)
	}
}

