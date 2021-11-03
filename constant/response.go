package constant

type Code string

const (
	// ErrorGlobal for global message error
	ErrorGlobal = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"

	// SuccessGetProfile when success get data profile user from header/authorization
	SuccessGetProfile Code = "201010"
	// FailedGetProfile failed when try to get data profile
	FailedGetProfile Code = "201040"

	// UserNotFound is when try to find data from database
	// and not found
	UserNotFound Code = "204540"
	// ErrorQuery is error when find data from database
	ErrorQuery Code = "209091"
	// ErrorParseDate error when parsing date
	ErrorParseDate Code = "209992"
)

var (
	Message = map[Code]string{
		SuccessGetProfile: "user.get.successfully",
		FailedGetProfile:  "user.get.failed",

		ErrorParseDate: "system.error.parse_date",
	}

	Reason = map[Code]string{
		UserNotFound:   "User tidak ditemukan",
		ErrorParseDate: "Tanggal tidak sesuai, harus berupa (YYYY-mm-dd)",
	}
)
