package f1telemetry

import "testing"

// TestTyreConsts tests if the tyre consts are correct
func TestTyreConsts(t *testing.T) {
	if RearLeft != 0 {
		t.Errorf("RearLeft is incorrect, want: %d, got: %d", 0, RearLeft)
	}
	if RearRight != 1 {
		t.Errorf("RearRight is incorrect, want: %d, got: %d", 1, RearRight)
	}
	if FrontLeft != 2 {
		t.Errorf("FrontLeft is incorrect, want: %d, got: %d", 2, FrontLeft)
	}
	if FrontRight != 3 {
		t.Errorf("FrontRight is incorrect, want: %d, got: %d", 3, FrontRight)
	}
}

func TestGetSpeedMPH(t *testing.T) {
	var telemetry = TelemetryPacket{}
	telemetry.Speed = 50 // m/s
	// the conversion is only rough anyway, within 1mph is acceptable enough for a test
	// (to make sure we're not wildly out of range)
	if int(telemetry.GetSpeedMPH()) != 111 {
		t.Errorf("GetSpeedMPH incorrect, got %d, want %d", int(telemetry.GetSpeedMPH()), 111)
	}
}

func TestPlayerCarInData(t *testing.T) {
	var telemetry = TelemetryPacket{}
	var cardata = CarData{}
	telemetry.PlayerCarIndex = 0
	telemetry.CarData[0] = cardata
	if telemetry.GetPlayerCarInData() != cardata {
		t.Errorf("GetPlayerCarInData failed, want %v, got %v", cardata, telemetry.GetPlayerCarInData())
	}
}

func TestGetTrackName(t *testing.T) {
	var telemetry = TelemetryPacket{TrackNumber: 7}
	if telemetry.GetTrackName() != "Silverstone" {
		t.Errorf("GetTrackName incorrect, want %s, got %s", "Silverstone", telemetry.GetTrackName())
	}
}

func TestGetTyreName(t *testing.T) {
	var telemetry = TelemetryPacket{TyreCompound: 3}
	if telemetry.GetTyreName() != "Medium" {
		t.Errorf("GetTyreName incorrect, want %s, got %s", "Medium", telemetry.GetTyreName())
	}
}

func TestGetDriverName(t *testing.T) {
	var telemetry = TelemetryPacket{}

	telemetry.Era = 2017
	data1 := CarData{DriverID: 9}
	result1, err := telemetry.GetDriverName(data1)
	if err != nil {
		t.Errorf("GetTeamName returned err: %v", err)
	}
	if result1 != "Lewis Hamilton" {
		t.Errorf("GetDriverName failed, want %s, got %s", "Lewis Hamilton", result1)
	}

	telemetry.Era = 1980
	data2 := CarData{DriverID: 3}
	result2, err := telemetry.GetDriverName(data2)
	if err != nil {
		t.Errorf("GetTeamName returned err: %v", err)
	}
	if result2 != "Sophie Levasseur" {
		t.Errorf("GetDriverName failed, want %s, got %s", "Sophie Levasseur", result2)
	}
}

func TestGetTeamName(t *testing.T) {
	var telemetry = TelemetryPacket{}

	telemetry.Era = 2017
	data1 := CarData{DriverID: 4}
	result1, err := telemetry.GetTeamName(data1)
	if err != nil {
		t.Errorf("GetTeamName returned err: %v", err)
	}
	if result1 != "Mercedes" {
		t.Errorf("GetTeamName failed, want %s, got %s", "Mercedes", result1)
	}

	telemetry.Era = 1980
	data2 := CarData{DriverID: 10}
	result2, err := telemetry.GetTeamName(data2)
	if err != nil {
		t.Errorf("GetTeamName returned err: %v", err)
	}
	if result2 != "Ferrari 2002" {
		t.Errorf("GetTeamName failed, want %s, got %s", "Ferrari 2002", result2)
	}
}
