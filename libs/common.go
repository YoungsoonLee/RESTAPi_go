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

	ErrInputData        = &ControllerError{400, "10001", "Data input error"}
	ErrInputDisplayname = &ControllerError{400, "10002", "Displayname should 4 ~ 16 letters."}

	ErrDatabase     = &ControllerError{500, "10002", "Database operation error"}
	ErrDupUser      = &ControllerError{400, "10003", "User information already exists"}
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
)
