package f1telemetry

// Various maps for data, used with various fields in the structs (driverID, etc)
var (
	// Tracks is a look up map for the track names
	Tracks = map[int]string{
		0:  "Melbourne",
		1:  "Sepang",
		2:  "Shanghai",
		3:  "Sakhir (Bahrain)",
		4:  "Catalunya",
		5:  "Monaco",
		6:  "Montreal",
		7:  "Silverstone",
		8:  "Hockenheim",
		9:  "Hungaroring",
		10: "Spa",
		11: "Monza",
		12: "Singapore",
		13: "Suzuka",
		14: "Abu Dhabi",
		15: "Texas",
		16: "Brazil",
		17: "Austria",
		18: "Sochi",
		19: "Mexico",
		20: "Baku (Azerbaijan)",
		21: "Sakhir Short",
		22: "Silverstone Short",
		23: "Texas Short",
		24: "Suzuka Short",
	}
	// Teams is a look up map for the normal team names
	Teams = map[int]string{
		0:  "Redbull",
		1:  "Ferrari",
		2:  "McLaren",
		3:  "Renault",
		4:  "Mercedes",
		5:  "Sauber",
		6:  "Force India",
		7:  "Williams",
		8:  "Toro Rosso",
		11: "Haas",
	}
	// ClassicTeams is a look up map for the classic team names
	ClassicTeams = map[int]string{
		0:  "Williams 1992",
		1:  "McLaren 1988",
		2:  "McLaren 2008",
		3:  "Ferrari 2004",
		4:  "Ferrari 1995",
		5:  "Ferrari 2007",
		6:  "McLaren 1998",
		7:  "Williams 1996",
		8:  "Renault 2006",
		10: "Ferrari 2002",
		11: "Redbull 2010",
		12: "McLaren 1991",
	}
	// Drivers is a look up map for the normal driver names
	Drivers = map[int]string{
		0:  "Sebastian Vettel",
		1:  "Daniil Kvyat",
		2:  "Fernando Alonso",
		3:  "Felipe Massa",
		5:  "Sergio Perez",
		6:  "Kimi Räikkönen",
		7:  "Romain Grosjean",
		9:  "Lewis Hamilton",
		10: "Nico Hulkenberg",
		14: "Kevin Magnussen",
		15: "Valtteri Bottas",
		16: "Daniel Ricciardo",
		18: "Marcus Ericsson",
		20: "Jolyon Palmer",
		22: "Max Verstappen",
		23: "Carlos Sainz Jr.",
		31: "Pascal Wehrlein",
		33: "Esteban Ocon",
		34: "Stoffel Vandoorne",
		35: "Lance Stroll",
	}
	// ClassicDrivers is a look up map for the classic driver names
	ClassicDrivers = map[int]string{
		0:  "Klimek Michalski",
		1:  "Martin Giles",
		2:  "Igor Correia",
		3:  "Sophie Levasseur",
		4:  "Alain Forest",
		5:  "Santiago Moreno",
		6:  "Esto Saari",
		7:  "Peter Belousov",
		8:  "Lars Kaufmann",
		9:  "Yasar Atiyeh",
		10: "Howard Clarke",
		14: "Marie Laursen",
		15: "Benjamin Coppens",
		16: "Alex Murray",
		18: "Callisto Calabresi",
		20: "Jay Letourneau",
		22: "Naota Izum",
		23: "Arron Barnes",
		24: "Jonas Schiffer",
		31: "Flavio Nieves",
		32: "Noah Visser",
		33: "Gert Waldmuller",
		34: "Julian Quesada",
		68: "Lucas Roth",
	}
	// Tyres is a look up map for the names of each tyre compound
	Tyres = map[int]string{
		0: "Ultra Soft",
		1: "Super Soft",
		2: "Soft",
		3: "Medium",
		4: "Hard",
		5: "Intermediate",
		6: "Wet",
	}
)

// TelemetryPacket is a Go representation of the UDP data received from the F1 2017 game
type TelemetryPacket struct {
	Time                 float32
	LapTime              float32
	LapDistance          float32
	TotalDistance        float32
	X                    float32    // World space position
	Y                    float32    // World space position
	Z                    float32    // World space position
	Speed                float32    // Speed of car in m/s (not MPH!)
	XV                   float32    // Velocity in world space
	YV                   float32    // Velocity in world space
	ZV                   float32    // Velocity in world space
	XR                   float32    // World space right direction
	YR                   float32    // World space right direction
	ZR                   float32    // World space right direction
	XD                   float32    // World space forward direction
	YD                   float32    // World space forward direction
	ZD                   float32    // World space forward direction
	SuspPos              [4]float32 // Wheel array order: RL, RR, FL, FR
	SuspVel              [4]float32
	WheelSpeed           [4]float32
	Throttle             float32
	Steer                float32
	Brake                float32
	Clutch               float32
	Gear                 float32
	GForceLat            float32
	GForceLon            float32
	Lap                  float32
	EngineRate           float32
	SliProNativeSupport  float32    // SLI Pro support
	CarPosition          float32    // car race position
	KersLevel            float32    // kers energy left
	KersMaxLevel         float32    // kers maximum energy
	DRS                  float32    // 0 off, 1 on
	TractionControl      float32    // 0 off - 2 high
	AntiLockBrakes       float32    // 0 off, 1 on
	FuelInTank           float32    // current fuel mass
	FuelCapacity         float32    // fuel capacity
	InPits               float32    // 0 = none, 1 = pitting, 2 = in pit area
	Sector               float32    // 0 = sector1, 1 = sector2, 2 = sector3
	Sector1Time          float32    // time of sector1 (or 0)
	Sector2Time          float32    // time of sector2 (or 0)
	BrakesTemp           [4]float32 // brakes temperature (centigrade)
	TyresPressure        [4]float32 // tyres pressure PSI
	TeamInfo             float32    // team ID
	TotalLaps            float32    // total number of laps in this race
	TrackSize            float32    // track size meters
	LastLapTime          float32    // last lap time
	MaxRpm               float32    // cars max RPM, at which point the rev limiter will kick in
	IdleRpm              float32    // cars idle RPM
	MaxGears             float32    // maximum number of gears
	SessionType          float32    // 0 = unknown, 1 = practice, 2 = qualifying, 3 = race
	DrsAllowed           float32    // 0 = not allowed, 1 = allowed, -1 = invalid / unknown
	TrackNumber          float32    // -1 for unknown, 0-21 for tracks
	VehicleFIAFlags      float32    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	Era                  float32    // era, 2017 (modern) or 1980 (classic)
	EngineTemperature    float32    // engine temperature (centigrade)
	GforceVert           float32    // vertical g-force component
	AngVelX              float32    // angular velocity x-component
	AngVelY              float32    // angular velocity y-component
	AngVelZ              float32    // angular velocity z-component
	TyresTemperature     [4]byte    // tyres temperature (centigrade)
	TyresWear            [4]byte    // tyre wear percentage
	TyreCompound         byte       // compound of tyre – 0 = ultra soft, 1 = super soft, 2 = soft, 3 = medium, 4 = hard, 5 = inter, 6 = wet
	FrontBrakeBias       byte       // front brake bias (percentage)
	FuelMix              byte       // fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	CurrentLapInvalid    byte       // current lap invalid - 0 = valid, 1 = invalid
	TyresDamage          [4]byte    // tyre damage (percentage)
	FrontLeftWingDamage  byte       // front left wing damage (percentage)
	FrontRightWingDamage byte       // front right wing damage (percentage)
	RearWingDamage       byte       // rear wing damage (percentage)
	EngineDamage         byte       // engine damage (percentage)
	GearBoxDamage        byte       // gear box damage (percentage)
	ExhaustDamage        byte       // exhaust damage (percentage)
	PitLimiterStatus     byte       // pit limiter status – 0 = off, 1 = on
	PitSpeedLimit        byte       // pit speed limit in m/s (not MPH!)
	SessionTimeLeft      float32    // NEW: time left in session in seconds
	RevLightsPercent     byte       // NEW: rev lights indicator (percentage)
	IsSpectating         byte       // NEW: whether the player is spectating
	SpectatorCarIndex    byte       // NEW: index of the car being spectated

	// Car data
	NumCars        byte        // number of cars in data
	PlayerCarIndex byte        // index of player's car in the array
	CarData        [20]CarData // data for all cars on track

	// New (v1.8)
	Yaw              float32
	Pitch            float32
	Roll             float32
	XLocalVelocity   float32    // Velocity in local space
	YLocalVelocity   float32    // Velocity in local space
	ZLocalVelocity   float32    // Velocity in local space
	SuspAcceleration [4]float32 // RL, RR, FL, FR
	AngAccX          float32    // angular acceleration x-component
	AngAccY          float32    // angular acceleration y-component
	AngAccZ          float32    // angular acceleration z-component
}

// CarData represents a secondary struct for specific car data
type CarData struct {
	WorldPosition     [3]float32 // world co-ordinates of vehicle
	LastLapTime       float32
	CurrentLapTime    float32
	BestLapTime       float32
	Sector1Time       float32
	Sector2Time       float32
	LapDistance       float32
	DriverID          byte
	TeamID            byte
	CarPosition       byte // UPDATED: track positions of vehicle
	CurrentLapNum     byte
	TyreCompound      byte // compound of tyre – 0 = ultra soft, 1 = super soft, 2 = soft, 3 = medium, 4 = hard, 5 = inter, 6 = wet
	InPits            byte // 0 = none, 1 = pitting, 2 = in pit area
	Sector            byte // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid byte // current lap invalid - 0 = valid, 1 = invalid
	Penalties         byte // NEW: accumulated time penalties in seconds to be added
}
