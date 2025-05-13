package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func sendPatliteViaSocket(ipAddress string, alertByte byte) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:10000", ipAddress))
	if err != nil {
		return err
	}

	bytes := []byte{0x57, alertByte}
	_, err = conn.Write(bytes)
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return err
	}
	return nil
}

func sendPatliteViaHttp(ipAddress string, alertStatus string) error {
	url := fmt.Sprintf("http://%s/api/control?alert=%s", ipAddress, alertStatus)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(byteArray) != "Success." {
		return fmt.Errorf("Failed to send alert to patlite.")
	}

	return nil
}
