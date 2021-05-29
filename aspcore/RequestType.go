package aspcore

type RequestType int

const (
	GET RequestType = iota
	POST
	PUT
	DELETE
)
