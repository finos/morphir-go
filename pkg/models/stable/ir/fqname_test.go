package ir

import (
	"encoding/json"
	"testing"
)

func TestFQNameJSONRoundTrip(t *testing.T) {
	original := FQNameFromParts(
		PathFromParts([]Name{NameFromParts([]string{"Excellent"}), NameFromParts([]string{"Package"})}),
		PathFromParts([]Name{NameFromParts([]string{"Fantastic"}), NameFromParts([]string{"Module"})}),
		NameFromParts([]string{"Amazing", "Local", "Name"}),
	)

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded FQName
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if !original.Equal(decoded) {
		pp1, mp1, ln1 := original.Parts()
		pp2, mp2, ln2 := decoded.Parts()
		t.Fatalf(
			"expected roundtrip equality; original=(%v,%v,%v) decoded=(%v,%v,%v)",
			pp1.Parts(), mp1.Parts(), ln1.Parts(),
			pp2.Parts(), mp2.Parts(), ln2.Parts(),
		)
	}
}

func TestFQNameUnmarshalRejectsWrongLength(t *testing.T) {
	var f FQName
	if err := json.Unmarshal([]byte(`[]`), &f); err == nil {
		t.Fatalf(expectedError)
	}
	if err := json.Unmarshal([]byte(`[[],[]]`), &f); err == nil {
		t.Fatalf(expectedError)
	}
	if err := json.Unmarshal([]byte(`[[],[],[],[]]`), &f); err == nil {
		t.Fatalf(expectedError)
	}
}

func TestFQNameUnmarshalRejectsInvalidParts(t *testing.T) {
	var f FQName
	if err := json.Unmarshal([]byte(`[{"bad":true}, [], []]`), &f); err == nil {
		t.Fatalf(expectedError)
	}
	if err := json.Unmarshal([]byte(`[[], {"bad":true}, []]`), &f); err == nil {
		t.Fatalf(expectedError)
	}
	if err := json.Unmarshal([]byte(`[[], [], {"bad":true}]`), &f); err == nil {
		t.Fatalf(expectedError)
	}
}
