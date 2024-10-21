package passcheck_test

import (
	"errors"
	"testing"

	"github.com/MatthiasSchild/passcheck"
)

func TestCheck_MinLength(t *testing.T) {
	err := passcheck.Check("hello", []passcheck.PasswordCheck{
		passcheck.MinLengthCheck(4),
	})
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	err = passcheck.Check("hello", []passcheck.PasswordCheck{
		passcheck.MinLengthCheck(6),
	})
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
	if !errors.Is(err, passcheck.ErrMinLength) {
		t.Errorf("Returns the wrong error: %v", err)
	}
}

func TestCheck_MaxLength(t *testing.T) {
	err := passcheck.Check("hello world", []passcheck.PasswordCheck{
		passcheck.MaxLengthCheck(32),
	})
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	err = passcheck.Check("hello world", []passcheck.PasswordCheck{
		passcheck.MaxLengthCheck(8),
	})
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
	if !errors.Is(err, passcheck.ErrMaxLength) {
		t.Errorf("Returns the wrong error: %v", err)
	}
}

func TestCheck_MinLowerLetterCount(t *testing.T) {
	err := passcheck.Check("!$Hello_123", []passcheck.PasswordCheck{
		passcheck.MinLowerLetterCount(4),
	})
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	err = passcheck.Check("!$Hello_123", []passcheck.PasswordCheck{
		passcheck.MinLowerLetterCount(5),
	})
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
	if !errors.Is(err, passcheck.ErrMinLowerLetter) {
		t.Errorf("Returns the wrong error: %v", err)
	}
}

func TestCheck_MinUpperLetterCount(t *testing.T) {
	err := passcheck.Check("!$hELLO_123", []passcheck.PasswordCheck{
		passcheck.MinUpperLetterCount(4),
	})
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	err = passcheck.Check("!$hELLO_123", []passcheck.PasswordCheck{
		passcheck.MinUpperLetterCount(5),
	})
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
	if !errors.Is(err, passcheck.ErrMinUpperLetter) {
		t.Errorf("Returns the wrong error: %v", err)
	}
}

func TestCheck_MinDigitCount(t *testing.T) {
	err := passcheck.Check("!$Hello_123", []passcheck.PasswordCheck{
		passcheck.MinDigitCount(3),
	})
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	err = passcheck.Check("!$Hello_123", []passcheck.PasswordCheck{
		passcheck.MinDigitCount(4),
	})
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
	if !errors.Is(err, passcheck.ErrMinDigits) {
		t.Errorf("Returns the wrong error: %v", err)
	}
}

func TestCheck_MinSpecialCharCount(t *testing.T) {
	err := passcheck.Check("!$Hello_123", []passcheck.PasswordCheck{
		passcheck.MinSpecialCharCount(3),
	})
	if err != nil {
		t.Errorf("Should not return an error: %v", err)
	}

	err = passcheck.Check("!$Hello_123", []passcheck.PasswordCheck{
		passcheck.MinSpecialCharCount(4),
	})
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
	if !errors.Is(err, passcheck.ErrMinSpecialChars) {
		t.Errorf("Returns the wrong error: %v", err)
	}
}
