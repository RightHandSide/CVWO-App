package users

const (
	ListUsers    = "users.HandleList"
	RegisterUser = "users.HandleRegister"

	SuccessfulListUsersMessage    = "Successfully Listed Users"
	SuccessfulRegisterUserMessage = "Successfully Registered User"
	ErrRetrieveDatabase           = "Failed to Retrieve Database in %s"
	ErrRetrieveUsers              = "Failed to Retrieve Users in %s"
	ErrRegisterUser               = "Failed to Register User in %s"
	ErrEncodeView                 = "Failed to Retrieve Users in %s"
)
