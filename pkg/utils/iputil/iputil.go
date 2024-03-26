package iputil

import (
	"github.com/sirupsen/logrus"
	"net"
)

func GetAvailableIP() []string {
	ips := []string{"127.0.0.1"}

	// Get all the local IP addresses
	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		logrus.Errorf("GetAvailableIP error: %v", err)
		return make([]string, 0)
	}

	for _, addr := range interfaces {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	return ips
}
