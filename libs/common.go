package libs

// ControllerError is controller error info structer.
type ControllerError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	//DevInfo  string `json:"dev_info"`
	//MoreInfo string `json:"more_info"`
}

// Predefined controller error values.
var (
	Err404 = &ControllerError{404, "404", "page not found"}

	// 10000 ~ related on account or auth
	ErrInputData        = &ControllerError{400, "10001", "Data input error"}
	ErrDisplayname      = &ControllerError{400, "10002", "Displayname should have 4 ~ 16 letters."}
	ErrEmail            = &ControllerError{400, "10003", "Must be a valid email address"}
	ErrMaxEmail         = &ControllerError{400, "10004", "Displayname cannot have over 100 letters."}
	ErrPassword         = &ControllerError{400, "10005", "Password should have 8 ~ 16 letters with number and special characters"}
	ErrDupDisplayname   = &ControllerError{400, "10006", "Displayname already exists"}
	ErrDupEmail         = &ControllerError{400, "10007", "Email already exists"}
	ErrAlreadyConfirmed = &ControllerError{400, "10008", "Email already confirmed."}
	ErrWrongToken       = &ControllerError{400, "10009", "wrong token."}
	ErrExpiredToken     = &ControllerError{401, "10010", "The token was already expired or invalid token. try again."}
	ErrPass             = &ControllerError{400, "10011", "User information does not exist or the password is incorrect"}
	ErrTokenAbsent      = &ControllerError{400, "10012", "Token absent"}
	ErrTokenInvalid     = &ControllerError{400, "10013", "Token invalid"}
	ErrTokenOther       = &ControllerError{400, "10014", "Token other"}
	ErrNoUser           = &ControllerError{400, "10015", "User information does not exist"}
	ErrIDAbsent         = &ControllerError{400, "10016", "Id absent"}
	ErrLoginFacebook    = &ControllerError{400, "10017", "Your disaplayname is connected a facebook. use facebook login."}
	ErrLoginGoogle      = &ControllerError{400, "10018", "Your disaplayname is connected a Google. use Google login."}

	ErrNoUserPass   = &ControllerError{400, "10006", "User information does not exist or the password is incorrect"}
	ErrNoUserChange = &ControllerError{400, "10007", "User information does not exist or data has not changed"}
	ErrInvalidUser  = &ControllerError{400, "10008", "User information is incorrect"}
	ErrOpenFile     = &ControllerError{500, "10009", "Error opening file"}
	ErrWriteFile    = &ControllerError{500, "10010", "Error writing a file"}
	ErrSystem       = &ControllerError{500, "10011", "Operating system error"}
	ErrExpired      = &ControllerError{400, "10012", "Login has expired"}
	ErrPermission   = &ControllerError{400, "10013", "Permission denied"}

	// 90000 ~ related on system error
	ErrDatabase      = &ControllerError{500, "90001", "Database operation error"}
	ErrJSONUnmarshal = &ControllerError{500, "90002", "JSON Unmarshal error"}
)
