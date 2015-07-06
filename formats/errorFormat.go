package formats

type ErrorFormat struct {
	ErrorCode int64 `json:",string"`
	ErrorMessage string
}
