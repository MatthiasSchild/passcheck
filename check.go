package passcheck

import (
	"errors"
	"slices"
)

var (
	// SpecialChars is a set of characters used for the special chars checks
	SpecialChars = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	ErrMinLength       = errors.New("the password is not long enough")
	ErrMaxLength       = errors.New("the password is too long")
	ErrMinLowerLetter  = errors.New("the password has not enough lower case letter")
	ErrMinUpperLetter  = errors.New("the password has not enough upper case letter")
	ErrMinDigits       = errors.New("the password has not enough digits")
	ErrMinSpecialChars = errors.New("the password has not enough special chars")
)

// PasswordCheck is the function definition for all checks usable for the Check function.
type PasswordCheck func(password string) error

// MinLengthCheck creates a check for the minimum length of the password.
// minLength is the minimum amount of characters the password should have
func MinLengthCheck(minLength int) PasswordCheck {
	return func(password string) error {
		if len(password) < minLength {
			return ErrMinLength
		}
		return nil
	}
}

// MaxLengthCheck creates a check for the maximum length of the password.
// maxLength is the maximum amount of characters the password should have
func MaxLengthCheck(maxLength int) PasswordCheck {
	return func(password string) error {
		if len(password) > maxLength {
			return ErrMaxLength
		}
		return nil
	}
}

// MinLowerLetterCount creates a check for the minimum amount of lower case letters (a-z)
// the password should have.
func MinLowerLetterCount(minCount int) PasswordCheck {
	return func(password string) error {
		counter := 0
		for _, c := range password {
			if c >= 'a' && c <= 'z' {
				counter++
				if counter == minCount {
					return nil
				}
			}
		}

		return ErrMinLowerLetter
	}
}

// MinUpperLetterCount creates a check for the minimum amount of upper case letters (A-Z)
// the password should have.
func MinUpperLetterCount(minCount int) PasswordCheck {
	return func(password string) error {
		counter := 0
		for _, c := range password {
			if c >= 'A' && c <= 'Z' {
				counter++
				if counter == minCount {
					return nil
				}
			}
		}

		return ErrMinUpperLetter
	}
}

// MinDigitCount creates a check for the minimum amount of digits (0-9)
// the password should have.
func MinDigitCount(minCount int) PasswordCheck {
	return func(password string) error {
		counter := 0
		for _, c := range password {
			if c >= '0' && c <= '9' {
				counter++
				if counter == minCount {
					return nil
				}
			}
		}

		return ErrMinDigits
	}
}

// MinSpecialCharCount creates a check for the minimum amount of special characters (see SpecialChars)
// the password should have.
func MinSpecialCharCount(minCount int) PasswordCheck {
	return func(password string) error {
		counter := 0
		for _, c := range password {
			if slices.Contains([]rune(SpecialChars), c) {
				counter++
				if counter == minCount {
					return nil
				}
			}
		}

		return ErrMinSpecialChars
	}
}

// Check executes a check of a given password and a list of checks.
// If the password fits the policy, nil is returned.
// Otherwise, the check will return its corresponding error.
func Check(password string, checks []PasswordCheck) error {
	var err error
	for _, check := range checks {
		err = check(password)
		if err != nil {
			return err
		}
	}

	return nil
}
