package core

import (
	"fmt"

	"github.com/ThotaGopichandThota/gn-core2/core/check"
)

// EnableEpochFlag defines a flag specific to the enableEpochs.toml
type EnableEpochFlag string

// CheckHandlerCompatibility checks if the provided handler is compatible with this gn-core2-go version
func CheckHandlerCompatibility(handler EnableEpochsHandler, requiredFlags []EnableEpochFlag) error {
	if check.IfNil(handler) {
		return ErrNilEnableEpochsHandler
	}

	for _, flag := range requiredFlags {
		if !handler.IsFlagDefined(flag) {
			return fmt.Errorf("%w for flag %s", ErrInvalidEnableEpochsHandler, flag)
		}
	}

	return nil
}
