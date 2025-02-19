package model

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Hardcoded admin credentials (use hashing for production)
var HardcodedAdmin = Admin{
	Username: "admin",
	Password: "securepassword123", // Change this to a secure password
}
