package registration


import (
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
	"time"
)


const (
	bodyType = "application/json"
	defAgentUrl = "http://localhost:8080/registration"

	interval time.Duration = 10 // Retry every 10 sec.
	retryMax = 10 // Retry 10 times before give up.
)


type RetrySetting struct {
	RetryMax int
	RetryInterval time.Duration
}


func RegisterService(agentUrl string, reg *Registration, retry RetrySetting) error {

	if retry.RetryMax == 0 {
		retry.RetryMax = retryMax
	}

	if retry.RetryInterval == 0 {
		retry.RetryInterval = interval
	}

	log.Println("Service Registration start...")

	if agentUrl == "" {
		log.Println("WARNING: Submit Agenet address is missing.")
		log.Println("WARNING: Use Default: ", defAgentUrl)
		agentUrl = defAgentUrl
	}

	// Can post multiple services.
	var regs []*Registration
	regs = append(regs, reg)

	retryCount := 0

	var regError error

	intSec := retry.RetryInterval * time.Second

	for retryCount < retry.RetryMax {
		regError = registerServices(agentUrl, regs, retry)

		if regError == nil {
			log.Println("Registered: ", *reg)
			return nil
		} else {
			log.Println("Retry in", intSec)
			time.Sleep(intSec)
		}
		retryCount++
	}

	log.Println("Could not register to Submit Agent:", agentUrl)
	log.Println("Make sure the Submit Agent is up and running.")
	log.Println("Running in single server mode...")

	return regError
}


// Register multiple services at once
func registerServices(agentUrl string, regs []*Registration, retry RetrySetting) error {

	// Encode JSON
	regJson, err := json.Marshal(regs)
	if err != nil {
		return err
	}
	log.Println("Registering service to the Agent...")

	// POST this service to submit agent
	res, err := http.Post(agentUrl, bodyType, bytes.NewReader(regJson))

	if err == nil {
		defer res.Body.Close()
		resBody, err := ioutil.ReadAll(res.Body)

		log.Println(string(resBody))
		return err
	} else {
		return err
	}
}
