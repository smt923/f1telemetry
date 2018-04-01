# f1telemetry
Go library for using the F1 2017 UDP telemetry, based on their post [here](http://forums.codemasters.com/discussion/comment/268912/#Comment_268912)

## Usage
For the most simple and probably unoptimal usage, you can call `ListenForTelemetryPackets` with an ip:port and a channel of TelemetryPacket to receive telemetry data as it's delivered:

```go
packets := make(chan t.TelemetryPacket)
t.ListenForTelemetryPackets("127.0.0.1:20777", packets)

for {
	packet := <-packets
	fmt.Printf("[Speed: %.02fmph | Engine RPM: %.02f, Engine temp: %.02fc]\n",
		packet.GetSpeedMPH(),
		packet.EngineRate,
		packet.EngineTemperature,
	)
}
```

For more advanced or generalized use, you can use `ParseBytesToPacket` on each packet (as a []byte) that you receive from your own UDP handling to receive a TelemetryPacket struct from your original packet

There are some helper methods and consts to make dealings with some of the data a little easier (`TyresTemperature[FrontLeft]` instead of remembering that front left is 2, and methods such as `GetTrackName()`)

## Notes
I kept as close to the Codemasters naming while sticking to Go conventions, this means "m_speed" becomes "Speed", however I kept the UK spelling of things such as "Tyre" to conform to what Codemasters own documentation says (link at the top of this readme)