package factory

import (
	"github.com/ThotaGopichandThota/gn-core2/hashing"
	"github.com/ThotaGopichandThota/gn-core2/hashing/blake2b"
	"github.com/ThotaGopichandThota/gn-core2/hashing/keccak"
	"github.com/ThotaGopichandThota/gn-core2/hashing/sha256"
)

// NewHasher will return a new instance of hasher based on the value stored in config
func NewHasher(name string) (hashing.Hasher, error) {
	switch name {
	case "sha256":
		return sha256.NewSha256(), nil
	case "keccak":
		return keccak.NewKeccak(), nil
	case "blake2b":
		return blake2b.NewBlake2b(), nil
	}

	return nil, ErrNoHasherInConfig
}
