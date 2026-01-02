package ir

import (
	"encoding/json"
	"testing"
)

func TestQNameJSONRoundTrip(t *testing.T) {
	original := QNameFromParts(
		PathFromParts([]Name{NameFromParts([]string{"Fantastic"}), NameFromParts([]string{"Module"})}),
		NameFromParts([]string{"Amazing", "Local", "Name"}),
	)

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded QName
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if !original.Equal(decoded) {
		mp1, ln1 := original.Parts()
		mp2, ln2 := decoded.Parts()
		t.Fatalf("expected roundtrip equality; original=(%v,%v) decoded=(%v,%v)", mp1.Parts(), ln1.Parts(), mp2.Parts(), ln2.Parts())
	}
}

func TestQNameUnmarshalRejectsWrongLength(t *testing.T) {
	var q QName
	if err := json.Unmarshal([]byte(`[]`), &q); err == nil {
		t.Fatalf(expectedError)
	}
	if err := json.Unmarshal([]byte(`[[],[],[]]`), &q); err == nil {
		t.Fatalf(expectedError)
	}
}

func TestQNameUnmarshalRejectsInvalidParts(t *testing.T) {
	var q QName
	// modulePath must be a Path (array of Names), localName must be a Name (array of strings).
	if err := json.Unmarshal([]byte(`[{"bad":true}, ["ok"]]`), &q); err == nil {
		t.Fatalf(expectedError)
	}
	if err := json.Unmarshal([]byte(`[[["ok"]], {"bad":true}]`), &q); err == nil {
		t.Fatalf(expectedError)
	}
}
