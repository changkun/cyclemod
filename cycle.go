package cyclemod

import (
	"fmt"

	"changkun.de/x/cyclemod2/v2"
)

// Call prints package version
func Call() {
	fmt.Printf("cyclemod@%s depends cyclemod2@%s\n",
		Version, cyclemod2.Version)
}

// Version is the version of the module
var Version = "v1.0.0"
