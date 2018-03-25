// Provides functionality for CAN communication

package communication

import (
    "github.com/waterloop/wcomms/wjson"
)

type CanBus struct {
}

// Reads Data from CAN bus and convert's binary packets to JSON packets
func (cb *CanBus) ReadAndParse() (*wjson.CommPacketJson) {
    return nil
}
