package enum

type MarkType string

const (
	MARK   MarkType = "mark"
	UNMARK MarkType = "unmark"
)

func (self MarkType) Equals(code string) bool {
	return string(self) == code
}
