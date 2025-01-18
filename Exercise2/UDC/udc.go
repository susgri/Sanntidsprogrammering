package main

import (
	"fmt"
	"net"
)

func readFromConn(conn net.PacketConn, available_terminal chan bool){
	// making a variable to save recieved message 
	receivedMessage := make([]byte, 1024)
	
	for { 
		// Reading from connection
		_, addr, err := conn.ReadFrom(receivedMessage)

		//checking for error messages 
		if err != nil {
			fmt.Println(err)
			return
		}

		// Waiting for terminal to be available
		<- available_terminal 

		// Printing recieved message 
		fmt.Println("Received message: ", string(receivedMessage), ", address: ",addr)

		// Releasing terminal
		available_terminal <- true
	}
}

func writeToConn(conn net.PacketConn, address string, available_terminal chan bool){
	addr, err2 := net.ResolveUDPAddr("udp", address)

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	i := int(0)
	message := []byte("hello")

	for {
		i++

		// writing the message 
		_, err := conn.WriteTo(message, addr)

		// Checking for error message 
		if err != nil {
			fmt.Println(err)
			return
		}

		// Waiting for terminal to be available
		<- available_terminal

		// Printing a confirmation that message nr i is sent
		fmt.Println(i,". Message sent")

		// Releasing terminal
		available_terminal <- true
	}

}


func main() {

	// Channel to signal use of terminal print
	available_terminal := make(chan bool) 

	// The adddress of the other machine. IP is unknown, but port is known
	address := ":30000"

	// Listening to the port
	conn, err := net.ListenPacket("udp", address)

	// Checking for errors 
	if err != nil {
		fmt.Println(err)
		return
	}

	// Making sure the connection is closed when the main function exits
	defer conn.Close()

	// Using my readFromConn to read the recieved messages
	go readFromConn(conn, available_terminal)

	// Sending messages
	go writeToConn(conn, address, available_terminal)

	// Preventing the main function from exiting
	select{}
}