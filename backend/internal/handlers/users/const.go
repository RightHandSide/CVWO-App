package users

const (
	ListUsers     = "users.HandleList"
	RegisterUsers = "users.HandleRegister"
	LoginUsers    = "users.HandleLogin"

	SuccessfulListUsersMessage     = "Successfully Listed Users"
	SuccessfulRegisterUsersMessage = "Successfully Registered User"
	SuccessfulLoginUsersMessage    = "User Login Successfully"

	NameRepeated   = "Name Repeated. Please Try Again"
	PleaseRegister = "Name not Found. Please Register"

	ErrRetrieveDatabase = "Failed to Retrieve Database in %s"

	ErrRetrieveUsers = "Failed to Retrieve Users in %s"
	ErrRegisterUsers = "Failed to Register User in %s"
	ErrLoginUser     = "Failed to Login User in %s"

	ErrEncodeView = "Failed to Retrieve in %s"
)
