# passcheck go library

This go library allows you to define password policies and check if a given password
matches the requirements.

## Example

    yourPassword := "helloWorld123!"
    err := passcheck.Check(yourPassword, []passcheck.PasswordCheck{
        passcheck.MinLengthCheck(8),
        passcheck.MaxLengthCheck(64),
        passcheck.MinLowerLetterCount(4),
        passcheck.MinUpperLetterCount(2),
        passcheck.MinDigitCount(2),
        passcheck.MinSpecialCharCount(3),
    })

## Special characters

For checking the special characters, the library has the variable `passcheck.SpecialChars`.
You can change the variable to specify the set of characters, which should count as
special characters.

    passcheck.SpecialChars = "+-!"

## Customizing errors

The errors, which are returned, when a case fails, are variables, which can be modified.
Set the error variables with the error you want, e.g. to customize the error messages.

    passcheck.ErrMinLength = errors.New("das passwort ist zu kurz")
