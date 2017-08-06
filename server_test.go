package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"testing"
	"time"
)

func validateTime(response []byte) error {
	var seconds_since_1900 uint32
	buf := bytes.NewReader(response)
	if err := binary.Read(buf, binary.BigEndian, seconds_since_1900); err != nil {
		return err
	}

	seconds_since_1970 := seconds_since_1900 - 2208988800
	var response_time = time.Unix(int64(seconds_since_1970), 0)
	if time.Since(response_time).Seconds() > 0 {
		return fmt.Errorf("Time difference too great. Expected: %v. Got: %v", time.Now(), response_time)
	}

	return nil
}

func TestTCPServer(t *testing.T) {
	config := DefaultConfig()
	// change port to be able to run as non-root
	config.TCPPort = 12345
	server := TCPServer{}
	go server.Run(config)
	time.Sleep(500 * time.Millisecond)
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	buffer := make([]byte, 4)
	conn.Read(buffer)
	if err = validateTime(buffer); err != nil {
		t.Error(err)
	}
}

func TestUDPServer(t *testing.T) {
	config := DefaultConfig()
	// change port to be able to run as non-root
	config.UDPPort = 12345
	server := UDPServer{}
	go server.Run(config)
	time.Sleep(500 * time.Millisecond)
	conn, err := net.Dial("udp", "127.0.0.1:12345")
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	conn.Write([]byte(""))
	buffer := make([]byte, 4)
	conn.Read(buffer)
	if err = validateTime(buffer); err != nil {
		t.Error(err)
	}
}

func BenchmarkUDP(b *testing.B) {
	config := DefaultConfig()
	ResetLogging(config)
	// change port to be able to run as non-root
	config.UDPPort = 12345
	server := UDPServer{}
	go server.Run(config)
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("udp", "127.0.0.1:12345")
		if err != nil {
			b.Fatal(err)
		}
		conn.Write([]byte(""))
		buffer := make([]byte, 4)
		conn.Read(buffer)
		if err = validateTime(buffer); err != nil {
			b.Error(err)
		}
		conn.Close()
	}
}

func BenchmarkUDPParallel(b *testing.B) {
	config := DefaultConfig()
	// Reset the logger to INFO
	ResetLogging(config)
	// change port to be able to run as non-root
	config.UDPPort = 12345
	server := UDPServer{}
	go server.Run(config)
	time.Sleep(500 * time.Millisecond)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			conn, err := net.Dial("udp", "127.0.0.1:12345")
			if err != nil {
				b.Error(err)
			}
			conn.Write([]byte(""))
			buffer := make([]byte, 4)
			conn.Read(buffer)
			if err = validateTime(buffer); err != nil {
				b.Error(err)
			}
			// using defer conn.Close() ends in "too many open files"
			conn.Close()
		}
	})
}

func BenchmarkTCP(b *testing.B) {
	config := DefaultConfig()
	ResetLogging(config)
	// change port to be able to run as non-root
	config.TCPPort = 12345
	server := TCPServer{}
	go server.Run(config)
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:12345")
		if err != nil {
			b.Fatal(err)
		}
		buffer := make([]byte, 4)
		conn.Read(buffer)
		if err = validateTime(buffer); err != nil {
			b.Error(err)
		}
		conn.Close()
	}
}

func BenchmarkTCPParallel(b *testing.B) {
	config := DefaultConfig()
	// Reset the logger to INFO
	ResetLogging(config)
	// change port to be able to run as non-root
	config.TCPPort = 12345
	server := UDPServer{}
	go server.Run(config)
	time.Sleep(500 * time.Millisecond)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			conn, err := net.Dial("tcp", "127.0.0.1:12345")
			if err != nil {
				b.Error(err)
			}
			buffer := make([]byte, 4)
			conn.Read(buffer)
			if err = validateTime(buffer); err != nil {
				b.Error(err)
			}
			// using defer conn.Close() ends in "too many open files"
			conn.Close()
		}
	})
}
