package models

import "regexp"

// Helper function to validate alphanumeric strings
func isValidAlphanumeric(s string) bool {
	match, _ := regexp.MatchString("^[A-Za-z0-9 ]+$", s)
	return match
}

// Helper function to validate phone numbers
func isValidPhoneNumber(s string) bool {
	match, _ := regexp.MatchString("^[0-9\\-]+$", s)
	return match
}
