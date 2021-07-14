package main

import (
	"fmt"
	"github.com/pbnjay/memory"
	"io/ioutil"
	"net"
	"os"
	"runtime"
)

const (
	BytesInMB = 1048576
	MBInGB = 1024

	HTMLFormat = `<!DOCTYPE html>
<html>
<body>
<h1>Hi!</h1>
<h3>Server stats:</h3>
<p>Hostname: %s<br>IP: %s<br>RAM: %s<br>Cores:%d<br></p>
</body>
</html>
`
)

// Returns the hostname or an empty string if it doesn't find any
func getHostname() string {
	host, _ := os.Hostname()
	return host
}

// Returns the public IP of the server (From https://www.cloudhadoop.com/2018/12/golang-examples-hostname-and-ip-address.html)
func getIP() string {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		fmt.Println(error)
	}

	defer conn.Close()
	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	return ipAddress.IP.String()
}

// Returns a prettified string of the total amount of memory in GB (Or MB if memory is smaller than 1 GB)
func getRAM() string {
	// Get total memory in bytes and div by number of bytes in MB
	ramMB := float64(memory.TotalMemory()) / BytesInMB
	// If total memory is larger than 1 GB, return result in GB, else return in MB
	if ramMB > MBInGB {
		return fmt.Sprintf("%.1fGB", ramMB / MBInGB)
	}
	return fmt.Sprintf("%.0fMB", ramMB)
}

// Returns the number of cores
func getCores() int {
	// NumOfPhysicalCores = NumOfVirtualCores / 2
	return runtime.NumCPU() / 2
}

func main() {
	// Format the HTML and insert the hostname, ip, memory and number of cores
	htmlResponse := fmt.Sprintf(HTMLFormat, getHostname(), getIP(), getRAM(), getCores())
	// Write html to file
	ioutil.WriteFile("index.html", []byte(htmlResponse), 0666)
}
