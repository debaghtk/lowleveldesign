# Interface Segregation Principle Problem Statement

## Background
A smart home system manages various devices like lights, thermostats, and security cameras. Each device can have different capabilities, such as turning on/off, adjusting settings, or streaming video.

## Problem
The current implementation violates ISP:

```go
type SmartDevice interface {
    TurnOn()
    TurnOff()
    SetTemperature(temp float64)
    StreamVideo() []byte
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

// Light doesn't need these methods, violating ISP
func (l *Light) SetTemperature(temp float64) {
    // No-op or panic
}

func (l *Light) StreamVideo() []byte {
    // No-op or panic
    return nil
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

// Thermostat doesn't need this method, violating ISP
func (t *Thermostat) StreamVideo() []byte {
    // No-op or panic
    return nil
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

// SecurityCamera doesn't need this method, violating ISP
func (c *SecurityCamera) SetTemperature(temp float64) {
    // No-op or panic
}
```

## Issues
1. The `SmartDevice` interface forces devices to implement methods they don't need.
2. Devices like `Light` and `SecurityCamera` have to provide meaningless implementations for `SetTemperature`.
3. Devices like `Thermostat` have to provide meaningless implementations for `StreamVideo`.
4. The interface is not cohesive, leading to poor design and potential runtime errors.

## Challenge
Refactor the smart home system to follow the Interface Segregation Principle so that:
1. Devices only implement methods relevant to their functionality.
2. Interfaces are cohesive and meaningful.
3. Client code can work with interfaces without knowing concrete types.
4. No methods should panic or return meaningless data.

## Requirements
1. Support existing device types: Light, Thermostat, Security Camera.
2. Ensure each device only implements relevant methods.
3. Maintain clean separation of concerns.
4. Make the code testable and maintainable.

## Bonus Challenge
Add a new device type (e.g., Smart Speaker) to demonstrate how your solution supports extension without violating ISP.

## Expected Usage
The solution should allow code like this to work naturally:

```go
func ControlDevice(device SmartDevice) {
    device.TurnOn()
    device.TurnOff()
}

// Specific operations should be safe and meaningful
func AdjustTemperature(thermostat TemperatureControl) {
    thermostat.SetTemperature(22.5)
}

func StreamFromCamera(camera VideoStreaming) {
    stream := camera.StreamVideo()
    fmt.Println("Streaming video:", stream)
}
```
