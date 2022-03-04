package main

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	goTextExamples()
	goI18NExamples()
}

func goTextExamples() {
	fmt.Println("/TEXT EXAMPLE RESULTS")
	fmt.Println("")
	// PRINTER EXAMPLE
	p := message.NewPrinter(language.BritishEnglish)
	p.Printf("There are %v flowers in our garden.\n", 1500) // retorna "There are 1,500 flowers in our garden"

	p = message.NewPrinter(language.BrazilianPortuguese)
	p.Printf("There are %v flowers in our garden.\n", 1500) // retorna "There are 1.500 flowers in our garden"

	// LANGUAGE TAG EXAMPLE

	ja, _ := language.ParseBase("ja")
	jp, _ := language.ParseRegion("JP")
	jpLngTag, _ := language.Compose(ja, jp)
	fmt.Println(jpLngTag)          // retorna "ja-JP"
	fmt.Println(language.Japanese) // retorna "ja-JP"

	// CATALOGS EXAMPLE
	message.SetString(language.Spanish, "Este é um exemplo\n", "esto es un ejemplo\n")
	message.SetString(language.English, "Este é um exemplo\n", "This is an example\n")

	p = message.NewPrinter(language.English)
	p.Printf("Este é um exemplo\n") // retorna This is an example
	p = message.NewPrinter(language.Spanish)
	p.Printf("Este é um exemplo\n") // retorna esto es un ejemplo
	fmt.Println("")
}

func goI18NExamples() {
	fmt.Println("GOI18N EXAMPLE RESULTS")
	fmt.Println("")
	// MESSAGE EXAMPLE
	exampleMsg := i18n.Message{
		ID:          "hello",
		Description: "descrição opcional da mensagem",
		Zero:        "Hello! you don't have any messages…",
		One:         "Hello! you have a message!",
		Two:         "Hello! you have two messages!",
		Few:         "Hello! you have a few messages!",
		Many:        "Hello! you have many messages!",
		Other:       "Hello! this is a default message",
	}
	fmt.Println(exampleMsg.Other)

	// BUNDLE EXAMPLE
	messageEnglish := i18n.Message{
		ID:    "hello",
		Other: "Hello!",
	}
	messageFrench := i18n.Message{
		ID:    "hello",
		Other: "Bonjour!",
	}
	bundle := i18n.NewBundle(language.English)
	bundle.AddMessages(language.English, &messageEnglish)
	bundle.AddMessages(language.French, &messageFrench)

	// LOCALIZER EXAMPLE
	localizer := i18n.NewLocalizer(bundle,
		language.French.String(),
		language.English.String())

	localizeConfig := i18n.LocalizeConfig{
		MessageID: "hello",
	}
	localization, _ := localizer.Localize(&localizeConfig)

	fmt.Println(localization) // retorna "Bonjour!"

	// MULTIPLE TAGS LOCALIZER EXAMPLE
	localizer = i18n.NewLocalizer(bundle, language.French.String(), language.English.String())
	loc, _ := localizer.Localize(&localizeConfig)
	fmt.Println(loc) // retorna "Bonjour!"

	localizer = i18n.NewLocalizer(bundle, language.English.String(), language.French.String())
	loc, _ = localizer.Localize(&localizeConfig)
	fmt.Println(loc) // retorna "Hello!"

	// PLURALIZATION EXAMPLE
	msg := i18n.Message{
		ID:    "pluralMSG",
		One:   "You have {{.PluralCount}} message!",
		Other: "You have {{.PluralCount}} messages!",
	}
	bundle = i18n.NewBundle(language.English)
	bundle.AddMessages(language.English, &msg)

	localizer = i18n.NewLocalizer(bundle, language.English.String())

	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   "pluralMSG",
		PluralCount: 1,
	})) // retorna "You have 1 message!"

	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   "pluralMSG",
		PluralCount: 2,
	})) // retorna "You have 2 messages!"

	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   "pluralMSG",
		PluralCount: "3",
	})) // retorna "You have 3 messages!"

	// DEFAULT MSG LOCALIZE CONFIG EXAMPLE
	defaultMSG := i18n.Message{
		ID:    "default",
		Other: "THIS IS A DEFAULT MESSAGE",
	}

	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:      "default",
		DefaultMessage: &defaultMSG,
	})) // retorna "THIS IS A DEFAULT MESSAGE"

}
