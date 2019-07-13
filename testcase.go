package huffy

// TestCase is a storable data object, which plays role of container for generated test cases.
// Tester encoded
type TestCase struct {
	ID   int64       `json:"id"`
	Data interface{} `json:"data"`
}
