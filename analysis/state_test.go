package analysis

import "testing"

func TestNewState(t *testing.T) {
	state := NewState()
	if state.Documents == nil {
		t.Fatalf("Expected state.Documents to not be nil")
	}
}

func TestOpenDocument(t *testing.T) {
	state := NewState()
	state.OpenDocument("test.txt", "test")
	if state.Documents["test.txt"] != "test" {
		t.Fatalf("Expected state.Documents[\"test.txt\"] to be \"test\"")
	}
}

func TestUpdateDocument(t *testing.T) {
	state := NewState()
	state.Documents["test.txt"] = "test"
	state.UpdateDocument("test.txt", "test updated")
	if state.Documents["test.txt"] != "test updated" {
		t.Fatalf("Expected state.Documents[\"test.txt\"] to be \"test updated\"")
	}
}
