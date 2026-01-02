package ir

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Name represents a Morphir IR Name.
//
// JSON encoding (Morphir-compatible): a Name is encoded as a JSON array of strings,
// for example: ["local","name"].
//
// This type implements encoding/json's hook methods MarshalJSON and UnmarshalJSON
// to provide custom marshalling/unmarshalling that matches the Morphir IR schemas.
//
// Note: We keep the underlying slice unexported to preserve immutability/value semantics.
// Use NameFromParts to construct and Parts() to retrieve a defensive copy.
type Name struct {
	parts []string
}

// NameFromParts constructs a Name from its component parts.
// The input slice is defensively copied.
func NameFromParts(parts []string) Name {
	if len(parts) == 0 {
		return Name{parts: nil}
	}
	copyParts := make([]string, len(parts))
	copy(copyParts, parts)
	return Name{parts: copyParts}
}

// Parts returns a defensive copy of the name parts.
func (n Name) Parts() []string {
	if len(n.parts) == 0 {
		return nil
	}
	copyParts := make([]string, len(n.parts))
	copy(copyParts, n.parts)
	return copyParts
}

// Equal performs structural equality.
func (n Name) Equal(other Name) bool {
	if len(n.parts) != len(other.parts) {
		return false
	}
	for i := range n.parts {
		if n.parts[i] != other.parts[i] {
			return false
		}
	}
	return true
}

// MarshalJSON implements encoding/json.Marshaler.
// It encodes the name as a JSON array of strings.
func (n Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.parts)
}

// UnmarshalJSON implements encoding/json.Unmarshaler.
// It decodes the name from a JSON array of strings.
func (n *Name) UnmarshalJSON(data []byte) error {
	if n == nil {
		return fmt.Errorf("ir.Name: UnmarshalJSON on nil receiver")
	}

	trimmed := bytes.TrimSpace(data)
	if bytes.Equal(trimmed, []byte("null")) {
		return fmt.Errorf("ir.Name: expected array of strings, got null")
	}

	var parts []string
	if err := json.Unmarshal(trimmed, &parts); err != nil {
		return fmt.Errorf("ir.Name: expected array of strings: %w", err)
	}

	// Defensive copy to prevent retaining references from json internals.
	*n = NameFromParts(parts)
	return nil
}
