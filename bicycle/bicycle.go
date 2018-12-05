package bicycle

// Bicycle is a bicycle
type Bicycle struct {
	ID            string         `json:"id,omitempty"`
	Make          string         `json:"make,omitempty"`
	Model         string         `json:"model,omitempty"`
	Size          int            `json:"size,omitempty"`
	FrameMaterial *FrameMaterial `json:"lastname,omitempty"`
}
