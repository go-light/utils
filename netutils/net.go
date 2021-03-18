package netutils

import (
	"fmt"
	"net"
)

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil{
		fmt.Println(err)
		return ""
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
