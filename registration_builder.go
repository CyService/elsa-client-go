package registration

import (
	"flag"
	"log"
	"strconv"
)


const (
	defId = "ci-service"
	defIp = "127.0.0.1"
	defCap = 4
	defVer = "v1"
	defPort = 3000
)


func NewRegistrationFromCommandline() *Registration {

	// Required parameters
	name := flag.String("id", defId, "Service name")
	ip := flag.String("ip", defIp, "This API server's IP address")

	// Optional
	cap := flag.Int("cap", defCap, "Number of instances")
	ver := flag.String("ver", defVer, "API version")
	port := flag.Int("port", defPort, "This API server's port")

	flag.Parse()

	// Check required fields
	if *name == defId {
		log.Println("Warning: Missing service endpoint name. Endpoint set to ", defId)
	}

	var serverIp *string

	if *ip == defIp {
		log.Println("Warning: IP set to", defIp, " You must provide IP address with '-ip' option.")
		serverIp = ip
	} else {
		curIp := GetIpAddress()
		serverIp = &curIp
	}

	return getReg(*name, *serverIp, *port, *cap, *ver)
}


func NewRegistration(name string, ip string, port int, cap int, ver string) *Registration {

	if name == "" {
		name = defId
	}

	if ip == "" {
		ip = defIp
	}

	if port == 0 {
		port = defPort
	}

	if cap == 0 {
		cap = defCap
	}

	if ver == "" {
		ver = defVer
	}

	return getReg(name, ip, port, cap, ver)
}


func getReg(name string, ip string, port int, cap int, ver string) *Registration {

	serviceUrl := ip + ":" + strconv.Itoa(port)
	log.Println("Service API address:", serviceUrl)

	instance := Instance{
		Capacity: cap,
		Location:serviceUrl,
	}

	return &Registration{
		Service: name,
		Version: ver,
		Instances: []Instance{instance},
	}
}
