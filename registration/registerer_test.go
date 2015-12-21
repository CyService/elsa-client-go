package registration


import (
	"testing"
	"encoding/json"
	"os"
)


func TestRegisterer(t *testing.T) {

	reg := NewRegistration("", "", 0, 0, "")
	json.NewEncoder(os.Stdout).Encode(reg)

	err := RegisterService("", reg, RetrySetting{RetryInterval:1, RetryMax:3})

	if err != nil {
		t.Log("Failed to register.")
		t.Log(err.Error())
//		t.Fail()
	} else {
		t.Log("Success!")
	}
}
