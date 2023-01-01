package main

import (
	"bufio"
	"fmt"
	"hydrachat/configurator"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))
	fmt.Println("Starting hydraChatClient....")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)

	configuration := configurator.Configuration{}
	err := configuration.GetConfiguration("chat.conf")
	if err != nil {
		panic(err)
	}

	//name = configuration.Name
	proto := "tcp"
	if !configuration.TCP {
		proto = "udp"
	}

	fmt.Printf("Hello %s, connecting to the hydra chat system.... \n", name)
	conn, err := net.Dial(proto, configuration.RemoteAddr)
	if err != nil {
		log.Fatal("Could not connect to hydra chat system", err)
	}
	fmt.Println("Connected to hydra chat system")
	name += ":"
	defer conn.Close()
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	for err == nil {
		msg := ""
		fmt.Print(name)
		fmt.Scan(&msg)
		msg = name + msg + "\n"
		fmt.Println("Duplicate: " + msg)
		_, err = fmt.Fprintf(conn, msg)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, name+msg+"\n")
	}
}
