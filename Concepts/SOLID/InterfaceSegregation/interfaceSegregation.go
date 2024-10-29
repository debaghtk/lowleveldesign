package interfacesegregation

type SmartDevice interface {
	TurnOn()
	TurnOff()
}

type SmartLight interface {
	SmartDevice
}

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
}

func (l *Light) TurnOff() {
	l.isOn = false
}

type SmartThermostat interface {
	SmartDevice
	SetTemperature(temp float64)
}

type Thermostat struct {
	temperature float64
}

func (t *Thermostat) TurnOn() {
	// Implementation
}

func (t *Thermostat) TurnOff() {
	// Implementation
}

func (t *Thermostat) SetTemperature(temp float64) {
	t.temperature = temp
}

type SmartSecurityCamera interface {
	SmartDevice
	StreamVideo() []byte
}

type SecurityCamera struct {
	isOn bool
}

func (c *SecurityCamera) TurnOn() {
	c.isOn = true
}

func (c *SecurityCamera) TurnOff() {
	c.isOn = false
}

func (c *SecurityCamera) StreamVideo() []byte {
	// Implementation
	return []byte("video stream")
}

type SmartSpeaker interface {
	SmartDevice
	PlayMusic() []byte
}

type Speaker struct {
	isOn bool
}

func (c *Speaker) TurnOn() {
	c.isOn = true
}

func (c *Speaker) TurnOff() {
	c.isOn = false
}

func (c *Speaker) PlayMusic() []byte {
	// Implementation
	return []byte("video stream")
}
