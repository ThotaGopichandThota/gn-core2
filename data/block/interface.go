package block

import "github.com/ThotaGopichandThota/gn-core2/data"

// EmptyBlockCreator is able to create empty block instances
type EmptyBlockCreator interface {
	CreateNewHeader() data.HeaderHandler
	IsInterfaceNil() bool
}
