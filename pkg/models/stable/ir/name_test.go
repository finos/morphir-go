package ir

import (
	"encoding/json"
	"testing"
)

func TestNameFromPartsCopiesInput(t *testing.T) {
	parts := []string{"local", "name"}
	name := NameFromParts(parts)

	parts[0] = "mutated"
	got := name.Parts()
	if len(got) != 2 || got[0] != "local" || got[1] != "name" {
		t.Fatalf("expected Name to be independent of input slice; got %#v", got)
	}

	got[0] = "mutated2"
	again := name.Parts()
	if again[0] != "local" {
		t.Fatalf("expected Parts() to return a copy; got %#v", again)
	}
}

func TestNameJSONRoundTrip(t *testing.T) {
	original := NameFromParts([]string{"Amazing", "Local", "Name"})

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded Name
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if !original.Equal(decoded) {
		t.Fatalf("expected roundtrip equality; original=%v decoded=%v", original.Parts(), decoded.Parts())
	}
}

func TestNameUnmarshalRejectsNonStringElements(t *testing.T) {
	var n Name
	err := json.Unmarshal([]byte(`["ok", 1]`), &n)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestNameUnmarshalRejectsNonArray(t *testing.T) {
	var n Name
	err := json.Unmarshal([]byte(`{"not":"an array"}`), &n)
	if err == nil {
		t.Fatalf("expected error")
	}
}
