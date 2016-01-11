package reg


import (
	"testing"
	"encoding/json"
	"os"
)

const(
	elsa = "http://192.168.99.100:8080/registration"
)


func TestRegisterer(t *testing.T) {

	// Local instance

	reg := NewRegistration("", "", 0, 0, "")
	json.NewEncoder(os.Stdout).Encode(reg)

	err := RegisterStruct(elsa, reg, Retry{Interval:1, Max:3})

	if err != nil {
		t.Log("Failed to register.")
		t.Log(err.Error())
//		t.Fail()
	} else {
		t.Log("Success!")
	}
}

func TestFileRegisterer(t *testing.T) {

	err := Register(elsa)

	if err != nil {
		t.Log("Failed to register.")
		t.Log(err.Error())
		//		t.Fail()
	} else {
		t.Log("Success File Reg!")
	}
}


func TestUnregisterer(t *testing.T) {

	// Local instance

	reg := NewRegistration("test-service", "", 0, 0, "")
	json.NewEncoder(os.Stdout).Encode(reg)

	err := RegisterStruct(elsa, reg, Retry{Interval:1, Max:3})

	if err != nil {
		t.Log("Failed to register.")
		t.Log(err.Error())
		t.Fail()
	} else {
		t.Log("Success!")
	}

	err = Unregister(elsa, reg)

	if err != nil {
		t.Log("Failed to unregister.")
		t.Log(err.Error())
		t.Fail()
	} else {
		t.Log("Success: Unregistration")
	}
}
