package pkg

// validatePassword will return true if the password is valid, false if not, the string value contains the reason
func validatePassword(password string) (bool, string) {
	if len(password) < 12 {
		return false, "password must be at least 12 characters"
	}
	// Make sure password contains at least one number
	if !containsNumber(password) {
		return false, "password must contain at least one number"
	}
	// Make sure password contains at least one uppercase letter
	if !containsUpper(password) {
		return false, "password must contain at least one uppercase letter"
	}
	// Make sure password contains at least one lowercase letter
	if !containsLower(password) {
		return false, "password must contain at least one lowercase letter"
	}
	// Make sure password contains at least one special character
	if !containsSpecial(password) {
		return false, "password must contain at least one special character"
	}
	return true, ""
}

// containsNumber will return true if the string contains a number, false if not
func containsNumber(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

// containsUpper will return true if the string contains an uppercase letter, false if not
func containsUpper(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

// containsLower will return true if the string contains a lowercase letter, false if not
func containsLower(s string) bool {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			return true
		}
	}
	return false
}

// containsSpecial will return true if the string contains a special character, false if not
func containsSpecial(s string) bool {
	for _, r := range s {
		if r >= '!' && r <= '/' {
			return true
		}
		if r >= ':' && r <= '@' {
			return true
		}
		if r >= '[' && r <= '`' {
			return true
		}
		if r >= '{' && r <= '~' {
			return true
		}
	}
	return false
}
