package service

import (
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(text string) string {
	if text == "" || len(text) <= 0 {
		log.Println("Передана пустая строка. Повторите ввод.")
		return ""
	}

	test := strings.ContainsAny(strings.ToLower(text), "абвгдеёжзийклмнопрстуфхцчшщъыьэюя1234567890") //
	// true = standart text,
	//  false = morse

	if test == true {
		strToMorse := morse.ToMorse(text)
		return strToMorse
	} else {
		var test2 bool
		for _, v := range text { // вторая проверка, проверяет на "точно морзе",
			//  если есть символы вне поиска - выводим ошибку и просим повторить ввод.
			if v == '.' || v == '-' || v == ' ' || v == '/' {
				test2 = true
			} else {
				test2 = false
				break
			}
		}
		if test2 == true {
			strToText := morse.ToText(text)
			return strToText
		} else {
			log.Println("Передана некорректная строка. Повторите ввод.")
			return ""
		}
	}

}
