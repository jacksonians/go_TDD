package main 
import "fmt"

const (
	spanish = "Spanish"
	french = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func determinePrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func Hello(name string, language string) string {
	if (name == "") {
		name = "World"
	}
	prefix := determinePrefix(language)
	return prefix + name

}

func main() {
	fmt.Println(Hello("world", "English"))
}