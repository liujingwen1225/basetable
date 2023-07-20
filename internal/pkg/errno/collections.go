package errno

var (

	//// ErrUserAlreadyExist 代表用户已经存在.
	//ErrUserAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "User already exist."}
	//
	//// ErrUserNotFound 表示未找到用户.
	//ErrUserNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.UserNotFound", Message: "User was not found."}
	//
	//// ErrPasswordIncorrect 表示密码不正确.
	//ErrPasswordIncorrect = &Errno{HTTP: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "Password is incorrect."}

	// ErrCollectionsCreate 集合创建失败.
	ErrCollectionsCreate = &Errno{HTTP: 400, Code: "FailedOperation.ErrCollectionsCreate", Message: "Collections create error."}
	ErrCollectionsExist  = &Errno{HTTP: 400, Code: "FailedOperation.ErrCollectionsExist", Message: "Collection name already exist (case insensitive)."}

	ErrCollectionsTypeNotFound = &Errno{HTTP: 400, Code: "FailedOperation.ErrCollectionsTypeNotFound", Message: "Collections type not found."}

	ErrCollectionsFieldsTypeNotFound = &Errno{HTTP: 400, Code: "FailedOperation.ErrCollectionsFieldsTypeNotFound", Message: "Collections fields type not found."}
)
