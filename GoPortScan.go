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

	var PORTS []int

	flag.StringVar(&targetIP, "targetIP", "127.0.0.1", "Target IP")
	flag.IntVar(&startPort, "startPort", 20, "Start Port")
	flag.IntVar(&endPort, "endPort", 1024, "End Port")
	flag.Parse()

	for i := startPort; i <= endPort; i++ {
		PORTS = append(PORTS, i)
	}
	var wg sync.WaitGroup

	for _, dstPort := range PORTS {
		wg.Add(1)
		go checkPortOpen(&wg, strconv.Itoa(dstPort))
	}
	wg.Wait()
}

func checkPortOpen(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	_, err := net.Dial("tcp", targetIP+":"+port)
	if err != nil {
		//log.Printf("Port %s closed", port)

	} else {
		log.Printf("Port %s open", port)
	}
}
