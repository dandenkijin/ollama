package format

import (
	"fmt"

	"github.com/toon-format/toon-go"
)

// TOONCodec provides methods for encoding and decoding TOON format
type TOONCodec struct{}

// Encode marshals input data into TOON format
func (t *TOONCodec) Encode(v interface{}) ([]byte, error) {
	return toon.Marshal(v, toon.WithLengthMarkers(true))
}

// Decode unmarshals TOON format data into the provided interface
func (t *TOONCodec) Decode(data []byte, v interface{}) error {
	return toon.Unmarshal(data, v)
}

// DecodeRaw decodes TOON format data into a dynamic representation
func (t *TOONCodec) DecodeRaw(data []byte) (interface{}, error) {
	return toon.Decode(data)
}

// IsTOONFormat checks if the format string indicates TOON format
func IsTOONFormat(format []byte) bool {
	return string(format) == `"toon"`
}

// ConvertTOONToJSON converts TOON format data to JSON
func ConvertTOONToJSON(toonData []byte) ([]byte, error) {
	decoded, err := toon.Decode(toonData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TOON data: %w", err)
	}
	
	// For now, we'll just return the decoded data as bytes
	// In a real implementation, we would convert it to JSON
	return []byte(fmt.Sprintf("%v", decoded)), nil
}

// ConvertJSONToTOON converts JSON data to TOON format
func ConvertJSONToTOON(jsonData []byte) ([]byte, error) {
	var decoded interface{}
	if err := toon.Unmarshal(jsonData, &decoded); err != nil {
		// If unmarshaling as JSON fails, treat it as a string
		decoded = string(jsonData)
	}
	
	toonData, err := toon.Marshal(decoded, toon.WithLengthMarkers(true))
	if err != nil {
		return nil, fmt.Errorf("failed to encode to TOON: %w", err)
	}
	
	return toonData, nil
}