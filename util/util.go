package util

import (
	"goWeb/model"
	"math/rand"
	"strings"
	"text/template"
	"time"
)

// Caracteres posibles
const letras = "abcdefghijklmnopqrstuvwxyz"
const numeros = "0123456789"
const simbolos = "!@#$%^&*()-_=+[]{}<>?/"

// GenerateRandomPassword generates a random password of the specified length.
// It can include letters (uppercase and lowercase), numbers, and symbols based on the parameters.
//
// Parameters:
//   - length: desired length of the generated password.
//   - useNumbers: if true, includes numbers in the password.
//   - useSymbols: if true, includes symbols in the password.
//
// Returns:
//   - A string representing the generated random password.
func GenerateRandomPassword(length int, useNumbers bool, useSymbols bool) string {
	// Base set: letters
	charset := letras                  // duplicate to include uppercase
	charset += strings.ToUpper(letras) // add uppercase letters

	// Add numbers and symbols according to options
	if useNumbers {
		charset += numeros
	}
	if useSymbols {
		charset += simbolos
	}

	// Generate password
	rand.NewSource(time.Now().UnixNano())
	randomPassword := make([]byte, length)
	for i := 0; i < length; i++ {
		randomPassword[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomPassword)
}

// PrintPasswordsTemplate renders a slice of Password objects using the provided template string.
// It parses the template, executes it for each Password in the slice, and concatenates the results,
// separating each rendered template with a newline. Returns the final string or an error if template
// parsing or execution fails.
//
// Parameters:
//   - passwords: A slice of pointers to model.Password objects to be rendered.
//   - tpl:       A string representing the Go template to use for rendering each Password.
//
// Returns:
//   - string: The concatenated rendered templates, separated by newlines.
//   - error:  An error if template parsing or execution fails.
func PrintPasswordsTemplate(passwords []*model.Password, tpl string) (string, error) {
	t, err := template.New("Passwords").Parse(tpl)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	for _, p := range passwords {
		if err := t.Execute(&sb, p); err != nil {
			return "", err
		}

		sb.WriteString("\n")
	}

	return sb.String(), nil
}