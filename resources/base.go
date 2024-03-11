package resources

var lang map[string]string
var commands map[string]string
var errors map[string]string

func Lang(key string) string {
	if name, ok := lang[key]; ok {
		return name
	}

	return key // todo: мб возвращать здесь что-то другое вместо ключа
}

func Command(key string) string {
	if name, ok := commands[key]; ok {
		return name
	}

	return key // todo: мб возвращать здесь что-то другое вместо ключа
}

func Error(key string) string {
	if name, ok := errors[key]; ok {
		return name
	}

	return key // todo: мб возвращать здесь что-то другое вместо ключа
}
