package viewmodel

// APIResult
type APIResult struct {
	Success bool
	Code    uint
	Data    interface{}
	Message string
	Count   uint
}

// Pager
type Pager struct {
	Page int
	Per  int
}
