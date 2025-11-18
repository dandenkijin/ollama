package format

import (
	"testing"
)

func TestTOONFormat(t *testing.T) {
	// Test IsTOONFormat function
	if !IsTOONFormat([]byte(`"toon"`)) {
		t.Error("IsTOONFormat should return true for \"toon\"")
	}
	
	if IsTOONFormat([]byte(`"json"`)) {
		t.Error("IsTOONFormat should return false for \"json\"")
	}
	
	// Test ConvertJSONToTOON function
	jsonData := []byte(`{"users": [{"id": 1, "name": "Alice"}]}`)
	toonData, err := ConvertJSONToTOON(jsonData)
	if err != nil {
		t.Errorf("ConvertJSONToTOON failed: %v", err)
	}
	
	if len(toonData) == 0 {
		t.Error("ConvertJSONToTOON should return non-empty data")
	}
	
	// Test ConvertTOONToJSON function
	jsonResult, err := ConvertTOONToJSON(toonData)
	if err != nil {
		t.Errorf("ConvertTOONToJSON failed: %v", err)
	}
	
	if len(jsonResult) == 0 {
		t.Error("ConvertTOONToJSON should return non-empty data")
	}
}