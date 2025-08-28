package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {

	input = strings.TrimSpace(input)
	if input == "" {

		return "", errors.New("Пустая строка")
	}

	isMorse := true

	for _, ch := range input {
		if ch != '.' && ch != '-' && !unicode.IsSpace(ch) {
			isMorse = false
			break
		}
	}

	if isMorse {
		result := morse.ToText(input)
		if strings.TrimSpace(result) == "" {
			return "", errors.New("Ошибка конвертации из Морзе в Текст")
		}
		return result, nil
	}
	for _, ch := range input {
		if _, ok := morse.DefaultMorse[unicode.ToUpper(ch)]; !ok {
			return input, nil
		}
	}
	result := morse.ToMorse(input)
	if strings.TrimSpace(result) == "" {
		return "", errors.New("Ошибка конвертации из Текста в Морзе")
	}
	return result, nil
}
