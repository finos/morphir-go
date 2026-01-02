package ir

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Path represents a Morphir IR Path.
//
// In Morphir IR JSON (morphir-elm `Morphir.IR.Path.Codec`), a path is encoded as a JSON
// array of Names, and each Name is encoded as an array of strings.
//
// This type implements encoding/json's hook methods MarshalJSON and UnmarshalJSON
// to provide custom marshalling/unmarshalling that matches the Morphir IR schemas.
//
// Example JSON:
//
//	[["morphir"],["s","d","k"]]
//
// Note: We keep the underlying slice unexported to preserve immutability/value semantics.
// Use PathFromParts to construct and Parts() to retrieve a defensive copy.
type Path struct {
	parts []Name
}

// PathFromParts constructs a Path from its component Name parts.
// The input slice is defensively copied.
func PathFromParts(parts []Name) Path {
	if len(parts) == 0 {
		return Path{parts: nil}
	}
	copyParts := make([]Name, len(parts))
	copy(copyParts, parts)
	return Path{parts: copyParts}
}

// Parts returns a defensive copy of the path parts.
func (p Path) Parts() []Name {
	if len(p.parts) == 0 {
		return nil
	}
	copyParts := make([]Name, len(p.parts))
	copy(copyParts, p.parts)
	return copyParts
}

// Equal performs structural equality.
func (p Path) Equal(other Path) bool {
	if len(p.parts) != len(other.parts) {
		return false
	}
	for i := range p.parts {
		if !p.parts[i].Equal(other.parts[i]) {
			return false
		}
	}
	return true
}

// MarshalJSON implements encoding/json.Marshaler.
// It encodes the path as a JSON array of names.
func (p Path) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.parts)
}

// UnmarshalJSON implements encoding/json.Unmarshaler.
// It decodes the path from a JSON array of names.
func (p *Path) UnmarshalJSON(data []byte) error {
	if p == nil {
		return fmt.Errorf("ir.Path: UnmarshalJSON on nil receiver")
	}

	trimmed := bytes.TrimSpace(data)
	if bytes.Equal(trimmed, []byte("null")) {
		return fmt.Errorf("ir.Path: expected array of names, got null")
	}

	var parts []Name
	if err := json.Unmarshal(trimmed, &parts); err != nil {
		return fmt.Errorf("ir.Path: expected array of names: %w", err)
	}

	*p = PathFromParts(parts)
	return nil
}
