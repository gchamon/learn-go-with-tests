package helloworld

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func getLanguagePrefix(language string) string {
	switch language {
	case "French":
		return frenchHelloPrefix
	case "Spanish":
		return spanishHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getLanguagePrefix(language) + name
}
