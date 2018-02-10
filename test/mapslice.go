package main

type PortBinding struct {
	// HostIP is the host IP Address
	HostIP string `json:"HostIp"`
	// HostPort is the host port number
	HostPort string
}

// PortMap is a collection of PortBinding indexed by Port
type PortMap map[Port][]PortBinding

// Port is a string containing port number and protocol in the format "80/tcp"
type Port string

func main() {
	p := PortMap{}
	pi := "80"
	p[Port(pi)] = []PortBinding{{HostIP: "12", HostPort: "8080"}}
	println(p["80"][0].HostIP)
}
