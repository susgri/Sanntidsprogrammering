package main

import (
	"fmt"
	"net"
)

func readFromConn(conn net.PacketConn){
	// making a variable to save recieved message 
	receivedMessage := make([]byte, 1024)
	
	// Reading from connection
	_, addr, err := conn.ReadFrom(receivedMessage)


	//checking for error messages 
	if err != nil {
		fmt.Println(err)
		return
	}

	// Printing recieved message 
	fmt.Println("Received message: ", string(receivedMessage), ", address: ",addr)
}

func writeToConn(conn net.PacketConn, addr net.Addr, message []byte){
	// writing the message 
	_, err := conn.WriteTo(message, addr)

	// Checking for error message 
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Message sent")
}


func main() {
	// The adddress of the other machine. IP is unknown, but port is known
	address := ":30000"

	// Listening to the port
	conn, err := net.ListenPacket("udp", address)

	// Checking for errors 
	if err != nil {
		fmt.Println(err)
		return
	}

	// Using my readFromConn to read the recieved messages
	readFromConn(conn)

	// Sending a message 
	message := []byte("Hello there!")

	addr, err2 := net.ResolveUDPAddr("udp", address)

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	writeToConn(conn, addr, message)

	// Closing the connection
	conn.Close()
}