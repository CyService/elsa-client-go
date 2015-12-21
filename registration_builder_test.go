package registration


import (
	"testing"
	"encoding/json"
	"os"
	"strconv"
)

func TestEmptyRegistration(t *testing.T) {

	t.Log("Registration test start...\n")

	reg := NewRegistration("", "", 0, 0, "")

	json.NewEncoder(os.Stdout).Encode(reg)

	if reg.Version != defVer {
		t.Log(reg.Version)
		t.Fail()
	}

	if reg.Service != defId {
		t.Log(reg.Service)
		t.Fail()
	}

	inst := reg.Instances

	if len(inst) != 1 {
		t.Log("Num inst should be 1.")
		t.Fail()
	} else {
		t.Log("OK")
	}
}


func TestRegistration(t *testing.T) {

	t.Log("Registration test start...\n")

	id := "idmapping"
	ip := GetIpAddress()
	port := 3333
	cap := 5
	ver := "v3"

	reg := NewRegistration(id, ip, port, cap, ver)

	json.NewEncoder(os.Stdout).Encode(reg)

	if reg.Version != ver {
		t.Log(reg.Version)
		t.Fail()
	}

	if reg.Service != id {
		t.Log(reg.Service)
		t.Fail()
	}

	inst := reg.Instances

	if len(inst) != 1 {
		t.Log("Num inst should be", 1)
		t.Fail()
	} else {
		t.Log("OK")
	}

	entry := inst[0]

	if entry.Capacity != cap {
		t.Log(entry.Capacity)
		t.Fail()
	}

	loc := ip + ":" + strconv.Itoa(port)
	if entry.Location != loc {
		t.Log(entry.Location)
		t.Fail()
	}
}
