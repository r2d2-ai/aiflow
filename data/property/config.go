package property

import (
	"os"
	"strings"
)

var EnvAppPropertySnapshotEnabled = "AIFLOW_APP_PROP_SNAPSHOTS"

func IsPropertySnapshotEnabled() bool {
	appPropertySnapshotEnabled := os.Getenv(EnvAppPropertySnapshotEnabled)
	if strings.EqualFold(appPropertySnapshotEnabled, "true") {
		return true
	}
	return false
}
