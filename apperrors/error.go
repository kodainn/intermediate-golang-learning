package apperrors

type ErrCode string

const (
	Unkown ErrCode = "U000"

	InsertDataField ErrCode = "S001"
	GetDataFailed   ErrCode = "S002"
	NAData          ErrCode = "S003"
	NoTargetData    ErrCode = "S004"
	UpdateDataFaild ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}
