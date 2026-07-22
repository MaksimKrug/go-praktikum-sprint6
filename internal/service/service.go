package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyInput = errors.New("service: empty input")
var ErrUnsupportedChars = errors.New("service: unsupported characters in input")

// Проверяет содержит ли текст только DefaultMorse символы
// p.s. не уверен нужна ли она, т.к. morse и так всё отбрасывает, но при этом т.к. ошибку он не возвращает, решил добавить свою
func isConvertible(s string) bool {
	for _, r := range s {
		if r == ' ' {
			continue
		}
		if _, ok := morse.DefaultMorse[unicode.ToUpper(r)]; !ok {
			return false
		}
	}
	return true
}

// Проверяет содержит ли текст только "." / "-" / " "
func isMorse(s string) bool {
	notMorse := strings.ContainsFunc(s, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})
	return !notMorse
}

// Convert конвертирует вход в противоположный формат.
// Неподдерживаемые символы выбрасываются; в этом случае возвращается частичный результат И ошибка ErrUnsupportedChars одновременно.
// p.s. опять же с одной стороны по общему правилу "ошибка == не доверяй результату", но от дать пользователю пустую строку вместо условного "test!" тоже не очень
func Convert(input string) (string, error) {
	if input == "" {
		return "", ErrEmptyInput
	}
	var err error
	if !isConvertible(input) {
		err = ErrUnsupportedChars
	}

	if isMorse(input) {
		return morse.ToText(input), err
	}
	return morse.ToMorse(input), err
}
