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

	ErrInputData                    = &ControllerError{400, "10001", "Data input error"}
	ErrDisplayname                  = &ControllerError{400, "10002", "Displayname should have 4 ~ 16 letters."}
	ErrEmail                        = &ControllerError{400, "10003", "Must be a valid email address"}
	ErrMaxEmail                     = &ControllerError{400, "10004", "Displayname cannot have over 100 letters."}
	ErrPassword                     = &ControllerError{400, "10005", "Password should have 8 ~ 16 letters with number and special characters"}
	ErrDupDisplayname               = &ControllerError{400, "10006", "Displayname already exists"}
	ErrDupEmail                     = &ControllerError{400, "10007", "Email already exists"}
	ErrAlreadyConfirmedOrWrongToken = &ControllerError{400, "10008", "Email already confirmed or wrong token. try again."}
	ErrExpiredToken                 = &ControllerError{400, "10009", "The token was already expired. try again."}

	ErrNoUser       = &ControllerError{400, "10004", "User information does not exist"}
	ErrPass         = &ControllerError{400, "10005", "User information does not exist or the password is incorrect"}
	ErrNoUserPass   = &ControllerError{400, "10006", "User information does not exist or the password is incorrect"}
	ErrNoUserChange = &ControllerError{400, "10007", "User information does not exist or data has not changed"}
	ErrInvalidUser  = &ControllerError{400, "10008", "User information is incorrect"}
	ErrOpenFile     = &ControllerError{500, "10009", "Error opening file"}
	ErrWriteFile    = &ControllerError{500, "10010", "Error writing a file"}
	ErrSystem       = &ControllerError{500, "10011", "Operating system error"}
	ErrExpired      = &ControllerError{400, "10012", "Login has expired"}
	ErrPermission   = &ControllerError{400, "10013", "Permission denied"}

	ErrDatabase = &ControllerError{500, "90001", "Database operation error"}
)
