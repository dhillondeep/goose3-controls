// Provides functionality for Serial communication

package communication

import (
    "github.com/tarm/serial"
    "github.com/waterloop/wcomms/wjson"
    "time"
)

type Config struct {
    Name        string
    Baud        int
    ReadTimeout time.Duration // Total timeout

    // Size is the number of data bits. If 0, DefaultSize is used.
    Size byte

    // Parity is the bit to use and defaults to ParityNone (no parity bit).
    Parity serial.Parity

    // Number of stop bits to use. Default is 1 (1 stop bit).
    StopBits serial.StopBits
}

type SerialBus struct {
    port *serial.Port
}

// Opens Serial Port based on the configuration file
func (sb *SerialBus) OpenPort(config *Config) (error) {
    serialConfig := serial.Config{
        Name: config.Name,
        Baud: config.Baud,
        ReadTimeout: config.ReadTimeout,
        Size: config.Size,
        Parity: config.Parity,
        StopBits: config.StopBits}

    port, err := serial.OpenPort(&serialConfig)
    sb.port = port

    return err
}

// Reads Data from Serial bus and convert's binary packets to JSON packets
func (sb *SerialBus) ReadAndParse() (*wjson.CommPacketJson) {
    return nil
}
