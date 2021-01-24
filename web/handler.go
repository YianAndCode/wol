package web

import (
	"fmt"
	"net"
	"net/http"

	"github.com/YianAndCode/wol/wol"
)

func handler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form.Get("mac"))

	localIP := ":7093"
	localAddr, err := net.ResolveUDPAddr("udp", localIP)
	if err != nil {
		fmt.Println(err)
		return
	}

	broadcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", localAddr, broadcastAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	packet, err := wol.New(req.Form.Get("mac"))
	if err != nil {
		fmt.Println(err)
		return
	}

	magicPacketBytes, err := packet.Marshal()
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := conn.Write(magicPacketBytes)
	if err != nil || n != 102 {
		fmt.Println(err)
		return
	}

	w.Write([]byte("Done"))
}
