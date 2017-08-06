package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Server interface {
	Run(config Config) error
}

type TCPServer struct {
}

func getTime() []byte {
	now := time.Now()
	secs := now.Unix()
	secs += 2208988800
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(secs))
	return b
}

func handleTCPRequest(connection net.Conn) {
	log.Debugf("Handling tcp request from %v", connection.RemoteAddr())
	connection.Write(getTime())
	connection.Close()
}

func (self TCPServer) Run(config Config) error {
	if !config.TCPEnabled {
		log.Info("TCP disabled")
		return nil
	}
	log.Info("TCP enabled")
	server_addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", config.TCPHost, config.TCPPort))

	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", server_addr)
	if err != nil {
		return err
	}
	log.Infof("Listening TCP on %v:%v", config.TCPHost, config.TCPPort)
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Errorf("Error accepting connection: %v", err)
		}
		go handleTCPRequest(connection)
	}
}

type UDPServer struct {
}

func handleUDPRequest(connection *net.UDPConn, addr *net.UDPAddr) {
	log.Debugf("Handling udp request from %v", addr)
	connection.WriteTo(getTime(), addr)
}

func (self UDPServer) Run(config Config) error {
	if !config.UDPEnabled {
		log.Info("UDP disabled")
		return nil
	}
	log.Info("UDP enabled")
	server_addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%v", config.UDPHost, config.UDPPort))

	if err != nil {
		return err
	}

	server_conn, err := net.ListenUDP("udp", server_addr)
	if err != nil {
		return err
	}
	log.Infof("Listening UDP on %v:%v", config.UDPHost, config.UDPPort)
	defer server_conn.Close()

	buff := make([]byte, 1024)

	for {
		_, addr, err := server_conn.ReadFromUDP(buff)
		if err != nil {
			log.Error(err)
		}
		go handleUDPRequest(server_conn, addr)
	}
}
