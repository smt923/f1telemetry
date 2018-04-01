package f1telemetry

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
	"net"
)

// Easy to memorize helpers for the various wheel arrays
const (
	RearLeft = iota
	RearRight
	FrontLeft
	FrontRight
)

// ParseBytesToPacket will take a byte slice and return a parsed packet struct
func ParseBytesToPacket(buf []byte) (TelemetryPacket, error) {
	r := bytes.NewReader(buf)
	packet := TelemetryPacket{}
	err := binary.Read(r, binary.LittleEndian, &packet)
	if err != nil {
		return TelemetryPacket{}, err
	}
	return packet, nil
}

// ListenForTelemetryPackets will connect to a server ("ip:port") and listen for data,
// then send a TelemetryPacket struct out to the given channel
//
// server: IP and port string, ex: "127.0.0.1:20777"
// channel: Channel of TelemetryPacket to send telemetry data to
func ListenForTelemetryPackets(server string, channel chan<- TelemetryPacket) {
	go listenLoop(server, channel)
}

func listenLoop(server string, channel chan<- TelemetryPacket) {
	addr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		log.Fatal("net.ResolveUDPAddr failure:\n    ", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("net.ListenUDP failure:\n    ", err)
	}
	defer conn.Close()

	var buf [1289]byte
	for {
		_, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			log.Println("conn.ReadFromUDP failure:\n    ", err)
			continue
		}

		packet, err := ParseBytesToPacket(buf[:])
		if err != nil {
			log.Println("Parse packet failure:\n    ", err)
			continue
		}
		channel <- packet
	}
}

// GetSpeedMPH returns the player's car's speed in MPH
func (p TelemetryPacket) GetSpeedMPH() float32 {
	return p.Speed * 2.2369
}

// GetPlayerCarInData returns the player's car from the CarData array
func (p TelemetryPacket) GetPlayerCarInData() CarData {
	return p.CarData[p.PlayerCarIndex]
}

// GetTrackName returns the string name of the track
func (p TelemetryPacket) GetTrackName() string {
	return Tracks[int(p.TrackNumber)]
}

// GetTyreName returns the string name of the track
func (p TelemetryPacket) GetTyreName() string {
	return Tyres[int(p.TyreCompound)]
}

// GetDriverName returns the string name of a driver from a CarData, handling era detection
func (p TelemetryPacket) GetDriverName(car CarData) (string, error) {
	if int(p.Era) == 1980 {
		return ClassicDrivers[int(car.DriverID)], nil
	}
	if int(p.Era) == 2017 {
		return Drivers[int(car.DriverID)], nil
	}
	return "", errors.New("Could not find a valid name, is era incorrect?")
}

// GetTeamName returns the string name of a team from a CarData, handling era detection
func (p TelemetryPacket) GetTeamName(car CarData) (string, error) {
	if int(p.Era) == 1980 {
		return ClassicTeams[int(car.DriverID)], nil
	}
	if int(p.Era) == 2017 {
		return Teams[int(car.DriverID)], nil
	}
	return "", errors.New("Could not find a valid team name, is era incorrect?")
}
