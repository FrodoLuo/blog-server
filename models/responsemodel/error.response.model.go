package responsemodel

type ErrorMessage struct {
	Message       string        `json:"message"`
	NestedMessage *ErrorMessage `json:"nestedMessage"`
}
