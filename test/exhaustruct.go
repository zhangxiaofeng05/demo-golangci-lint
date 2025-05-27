package main

import "log"

type MockAlarm struct {
	Token string
}

type MockClient struct {
	Name      string
	MockAlarm *MockAlarm
}

func exhaustruct() {
	mc := &MockClient{
		Name: "test",
		//MockAlarm: &MockAlarm{
		//	Token: "mock token",
		//},
	}
	log.Println(mc.Name)
	log.Println(mc.MockAlarm.Token)
}
