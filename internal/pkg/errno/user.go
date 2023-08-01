package errno

var (
	// ErrUserAlreadyExist 代表用户已经存在.
	ErrUserAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "User already exist."}

	// ErrUserNotFound 表示未找到用户.
	ErrUserNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.UserNotFound", Message: "User was not found."}

	// ErrPasswordIncorrect 表示密码不正确.
	ErrPasswordIncorrect = &Errno{HTTP: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "Password is incorrect."}

	// ErrUserDeleted
	ErrUserDeleted = &Errno{HTTP: 400, Code: "InvalidParameter.ErrUserDeleted", Message: "User delete error."}
	// ErrUserList
	ErrUserList = &Errno{HTTP: 400, Code: "InvalidParameter.ErrUserList", Message: "User list find error."}

	ErrUserIdType = &Errno{HTTP: 400, Code: "InvalidParameter.ErrUserIdType", Message: "User id  was number."}
)
