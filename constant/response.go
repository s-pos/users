package constant

type Code string

const (
	// ErrorParseDate error when parsing date
	ErrorParseDate Code = "209992"
)

var (
	Message = map[Code]string{
		ErrorParseDate: "system.error.parse_date",
	}

	Reason = map[Code]string{}
)
