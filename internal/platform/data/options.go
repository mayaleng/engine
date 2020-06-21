package data

// FindOptions defines a find multiple documents options
type FindOptions struct {
	Filter map[string]interface{}
	Limit  int64
	Skip   int64
	Order  map[string]int
}
