package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

func main() {
	counter_backup := 0

	cmd := exec.Command("cmd", "/C", "start", "powershell", "go", "run", "program/program.go")

	addr := ":20009"
	conn, err := net.ListenPacket("udp", addr)

	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Println("Connection made")

	message := make([]byte, 4)

	// ----- BACKUP PHASE -------
	fmt.Println("####### Backup phase #######")

	for {
		err_lostConnection := conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		if err_lostConnection != nil {
			break
		}
		_, _, err_reading := conn.ReadFrom(message)

		if err_reading != nil {
			fmt.Println("Taking over as primary...")
			break
		}

		counter_backup = int(binary.BigEndian.Uint32(message))
	}

	conn.Close() // making sure to close connection

	fmt.Println("####### Primary phase #######")
	cmd.Run()

	counter_primary := counter_backup

	conn_primary, err_primary := net.Dial("udp", addr)

	if err_primary != nil {
		fmt.Println("Error connecting: ", err)
	}

	defer conn_primary.Close()

	fmt.Println("Counting:")

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		message_send := make([]byte, 4)
		//message := []byte(string(counter))

		counter_primary++
		binary.BigEndian.PutUint32(message_send, uint32(counter_primary))

		// Send the message
		_, err_sending := conn_primary.Write(message_send)

		if err_sending != nil {
			fmt.Println("Error sending message:", err_sending)
			break
		}
		fmt.Println(counter_primary)

	}

	fmt.Println("DEAD")
	time.Sleep(time.Second * 2)
}
