package apperrors

type ErrCode string

const (
	Unknown          ErrCode = "U000"
	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NoData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
	ReqBodyEncodeFailed ErrCode = "R003"
)
