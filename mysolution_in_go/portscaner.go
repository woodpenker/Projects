package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

var (
	ip       string
	max_port int
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "input dest ip address,default is localhost")
	flag.IntVar(&max_port, "p", 2048, "input max scan port ,default is 2048")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("You can give your dest ip like: -ip 192.168.0.1 -p 1024 .Default ip: 127.0.0.1 port:2048")
	}
	fmt.Println(ip)
	c := make(chan int)
	for i := 1; i <= max_port; i++ {
		go scanport(ip, i, c)
		time.Sleep(10 * time.Millisecond)
	}
	count := 0
	for {
		count += <-c
		//	log.Printf("count:", count)
		if count >= max_port {
			break
		}
	}
	fmt.Println("Done!")

}
func scanport(ip string, i int, c chan int) {
	conn, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(i), 2*time.Second)
	if err == nil {
		fmt.Printf("%s's port %d is opening !\n", ip, i)
		conn.Close()
	} else if err != nil && conn != nil {
		//	log.Print("2:", err)
		conn.Close()
	}
	c <- 1
}
