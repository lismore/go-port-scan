package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/color"
	_ "github.com/google/gopacket/layers"
)

var (
	author  string
	version string

	startPort int
	endPort   int
	targetIP  string
)

func banner() {
	name := fmt.Sprintf("go-port-scan (v.%s)", version)
	banner := `
	________                __________              __              _________                     
	/  _____/  ____          \______   \____________/  |_           /   _____/ ____ _____    ____  
   /   \  ___ /  _ \   ______ |     ___/  _ \_  __ \   __\  ______  \_____  \_/ ___\\__  \  /    \ 
   \    \_\  (  <_> ) /_____/ |    |  (  <_> )  | \/|  |   /_____/  /        \  \___ / __ \|   |  \
	\______  /\____/          |____|   \____/|__|   |__|           /_______  /\___  >____  /___|  /
		   \/                                                              \/     \/     \/     \/ 										
	`

	// window width
	all_lines := strings.Split(banner, "\n")
	w := len(all_lines[1])

	// Centered
	fmt.Println(banner)
	color.Green(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(name))/2, name)))
	color.Blue(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(author))/2, author)))
	fmt.Println()
}

func init() {

	banner()
}

func main() {

	//define array for port range
	var PORTS []int

	//capture arguments
	flag.StringVar(&targetIP, "targetIP", "127.0.0.1", "Target IP")
	flag.IntVar(&startPort, "startPort", 20, "Start Port")
	flag.IntVar(&endPort, "endPort", 1024, "End Port")
	flag.Parse()

	//generate a range of ports between start and end port values
	for i := startPort; i <= endPort; i++ {
		PORTS = append(PORTS, i)
	}

	//define a waitGroup
	var wg sync.WaitGroup

	//iterate through range of ports calling the checkPortOpen function
	for _, dstPort := range PORTS {
		wg.Add(1)
		go checkPortOpen(&wg, strconv.Itoa(dstPort))
	}

	//waiting for all goroutines to finish
	wg.Wait()
}

func checkPortOpen(wg *sync.WaitGroup, port string) {

	//deferring wait group for done
	defer wg.Done()

	//connects to the address on the target network
	_, err := net.Dial("tcp", targetIP+":"+port)

	//handle any error
	if err != nil {
		//TODO
		//log.Printf("Port %s closed", port)

	} else {
		//write out to console if port is open
		log.Printf("Port %s open", port)
	}
}
