// Interface for a bus that is used to communicate data among devices

package communication

import (
    "github.com/waterloop/wcomms/wjson"
)

type Bus interface {
    // Reads the bus and parses the binary data into JSON
    ReadAndParse() (*wjson.CommPacketJson)
}
