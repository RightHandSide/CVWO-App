package users

const (
	ListUsers     = "users.HandleList"
	RegisterUsers = "users.HandleRegister"

	SuccessfulListUsersMessage     = "Successfully Listed Users"
	SuccessfulRegisterUsersMessage = "Successfully Registered User"
	NameRepeated                   = "Name Repeated. Please Try Again"
	ErrRetrieveDatabase            = "Failed to Retrieve Database in %s"
	ErrRetrieveUsers               = "Failed to Retrieve Users in %s"
	ErrRegisterUsers               = "Failed to Register User in %s"
	ErrEncodeView                  = "Failed to Retrieve Users in %s"
)
