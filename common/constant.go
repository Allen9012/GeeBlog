package common

import "time"

const (
	RoleUser  = 0
	RoleAdmin = 1

	KeyUser     = "user"
	KeyUserId   = "user_id"
	KeyUserRole = "user_role"

	KeyPostId = "post_id"

	TokenHeader  = "Authorization"
	TokenEmpty   = ""
	TokenPrefix  = "login:token:"
	TokenTimeout = time.Hour * 24
)
