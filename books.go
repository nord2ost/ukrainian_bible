package english_bible

import (
	"regexp"
	"strings"
)

var BookNames = map[string]string{
	// Old Testament
	"gen":               "Бут.",
	"genesis":           "Бут.",
	"бут":               "Бут.",
	"буття":             "Бут.",
	"exod":              "Вих.",
	"exodus":            "Вих.",
	"вих":               "Вих.",
	"вихід":             "Вих.",
	"lev":               "Лев.",
	"leviticus":         "Лев.",
	"лев":               "Лев.",
	"левит":             "Лев.",
	"num":               "Чис.",
	"numbers":           "Чис.",
	"чис":               "Чис.",
	"числа":             "Чис.",
	"deut":              "Втор.",
	"deuteronomy":       "Втор.",
	"втор":              "Втор.",
	"второзаконня":      "Втор.",
	"josh":              "Нав.",
	"joshua":            "Нав.",
	"нав":               "Нав.",
	"навина":            "Нав.",
	"judges":            "Суд.",
	"judg":              "Суд.",
	"суд":               "Суд.",
	"суддів":            "Суд.",
	"ruth":              "Руф.",
	"руф":               "Руф.",
	"1 kgs":             "1Цар.",
	"1 kings":           "1Цар.",
	"1 царів":           "1Цар.",
	"1 цар":             "1Цар.",
	"2 kgs":             "2Цар.",
	"2 kings":           "2Цар.",
	"2 цар":             "2Цар.",
	"2 царів":           "2Цар.",
	"3 kgs":             "3Цар.",
	"3 kings":           "3Цар.",
	"3 цар":             "3Цар.",
	"3 царів":           "3Цар.",
	"4 kgs":             "4Цар.",
	"4 kings":           "4Цар.",
	"4 цар":             "4Цар.",
	"4 царів":           "4Цар.",
	"1 chr":             "1Пар.",
	"1 chronicles":      "1Пар.",
	"1 пар":             "1Пар.",
	"1 паралипоменон":   "1Пар.",
	"2 chr":             "2Пар.",
	"2 chronicles":      "2Пар.",
	"2 пар":             "2Пар.",
	"2 паралипоменон":   "2Пар.",
	"ezra":              "EZR",
	"1 езд":             "1Езд.",
	"1 ездри":           "1Езд.",
	"2 езд":             "2Езд.",
	"2 ездри":           "2Езд.",
	"3 езд":             "3Езд.",
	"3 ездри":           "3Езд.",
	"neh":               "Неєм.",
	"nehemiah":          "Неєм.",
	"неєм":              "Неєм.",
	"неємії":            "Неєм.",
	"esth":              "Есф.",
	"esther":            "Есф.",
	"есф":               "Есф.",
	"есфирі":            "Есф.",
	"job":               "Іов.",
	"іов":               "Іов.",
	"ps":                "Пс.",
	"psalm":             "Пс.",
	"psalms":            "Пс.",
	"пс":                "Пс.",
	"псалом":            "Пс.",
	"псалми":            "Пс.",
	"псалтирь":          "Пс.",
	"prov":              "Притч.",
	"proverbs":          "Притч.",
	"притч":             "Притч.",
	"притчі":            "Притч.",
	"eccl":              "Еккл.",
	"ecclesiastes":      "Еккл.",
	"еккл":              "Еккл.",
	"екклезіаст":        "Еккл.",
	"song":              "Пісн.",
	"song of solomon":   "Пісн.",
	"song of songs":     "Пісн.",
	"пісн":              "Пісн.",
	"пісня над піснями": "Пісн.",
	"isa":               "Іса.",
	"isaiah":            "Іса.",
	"іса":               "Іса.",
	"ісая":              "Іса.",
	"jer":               "Єр.",
	"jeremiah":          "Єр.",
	"єр":                "Єр.",
	"єремія":            "Єр.",
	"hos":               "Ос.",
	"hosea":             "Ос.",
	"ос":                "Ос.",
	"осія":              "Ос.",
	"joel":              "Іоїл.",
	"іоїл":              "Іоїл.",
	"amos":              "Ам.",
	"амоса":             "Ам.",
	"ам":                "Ам.",
	"obad":              "Авд.",
	"obadiah":           "Авд.",
	"авд":               "Авд.",
	"авдія":             "Авд.",
	"jonah":             "Іона.",
	"jon":               "Іона.",
	"іона":              "Іона.",
	"іон":               "Іона.",
	"mic":               "Мих.",
	"micah":             "Мих.",
	"мих":               "Мих.",
	"михея":             "Мих.",
	"nah":               "Наум.",
	"nahum":             "Наум.",
	"наума":             "Наум.",
	"наум":              "Наум.",
	"hab":               "Авв.",
	"habakkuk":          "Авв.",
	"авв":               "Авв.",
	"аввакума":          "Авв.",
	"zech":              "Зах.",
	"zechariah":         "Зах.",
	"зах":               "Зах.",
	"захарії":           "Зах.",
	"hag":               "Агг.",
	"hagai":             "Агг.",
	"агг":               "Агг.",
	"аггея":             "Агг.",
	"lam":               "Плач.",
	"lamentations":      "Плач.",
	"плач єремії":       "Плач.",
	"плач":              "Плач.",
	"ezek":              "Єз.",
	"ezekiel":           "Єз.",
	"єзекиїля":          "Єз.",
	"єз":                "Єз.",
	"dan":               "Дан.",
	"daniel":            "Дан.",
	"дан":               "Дан.",
	"даниїла":           "Дан.",
	"zeph":              "Соф.",
	"zephaniah":         "Соф.",
	"соф":               "Соф.",
	"софонії":           "Соф.",
	"mal":               "Мал.",
	"malachi":           "Мал.",
	"мал":               "Мал.",
	"малахії":           "Мал.",

	// Deuterocanonical
	"tobit":               "Тов.",
	"tob":                 "Тов.",
	"тов":                 "Тов.",
	"товит":               "Тов.",
	"judith":              "Юдиф.",
	"юдиф":                "Юдиф.",
	"юдифі":               "Юдиф.",
	"additions to esther": "ESG",
	"wis":                 "Прем.",
	"wisdom":              "Прем.",
	"wisdom of solomon":   "Прем.",
	"прем":                "Прем.",
	"премудрості":         "Прем.",
	"премудрості соломона": "Прем.",
	"sirach":           "Сир.",
	"wisdom of sirach": "Сир.",
	"сир":              "Сир.",
	"сираха":           "Сир.",
	"премудрости ісуса, сина сирахового": "Сир.",
	"вар":                    "Вар.",
	"варуха":                 "Вар.",
	"bareuch":                "Вар.",
	"letter of jeremiah":     "Посл. Єр.",
	"послання єремії":        "Посл. Єр.",
	"посл єр":                "Посл. Єр.",
	"song of the three":      "S3Y",
	"prayer of azariah":      "S3Y",
	"susanna":                "SUS",
	"bel and the dragon":     "BEL",
	"1 maccabees":            "1Мак.",
	"2 maccabees":            "2Мак.",
	"3 maccabees":            "3Мак.",
	"1 маккавеїв":            "1Мак.",
	"2 маккавеїв":            "2Мак.",
	"3 маккавеїв":            "3Мак.",
	"4 maccabees":            "4MA",
	"manasseh":               "MAN",
	"the prayer of manasseh": "MAN",

	// New Testament
	"matt":            "Мф.",
	"matthew":         "Мф.",
	"mt":              "Мф.",
	"мф":              "Мф.",
	"матфея":          "Мф.",
	"mark":            "Мк.",
	"mk":              "Мк.",
	"марка":           "Мк.",
	"мк":              "Мк.",
	"luke":            "Лк.",
	"lk":              "Лк.",
	"луки":            "Лк.",
	"лк":              "Лк.",
	"john":            "Ін.",
	"jn":              "Ін.",
	"Іоана":           "Ін.",
	"ін":              "Ін.",
	"acts":            "Діян.",
	"діян":            "Діян.",
	"діяння":          "Діян.",
	"rom":             "Рим.",
	"римлян":          "Рим.",
	"рим":             "Рим.",
	"1 cor":           "1Кор.",
	"1 corinthians":   "1Кор.",
	"1 кор":           "1Кор.",
	"1 коринф'ян":     "1Кор.",
	"2 cor":           "2Кор.",
	"2 corinthians":   "2Кор.",
	"2 кор":           "2Кор.",
	"2 коринф'ян":     "2Кор.",
	"1 thess":         "1Сол.",
	"1 thessalonians": "1Сол.",
	"1 сол":           "1Сол.",
	"1 солунян":       "1Сол.",
	"2 thess":         "2Сол.",
	"2 thessalonians": "2Сол.",
	"2 сол":           "2Сол.",
	"2 солунян":       "2Сол.",
	"gal":             "Гал.",
	"galatians":       "Гал.",
	"гал":             "Гал.",
	"галатів":         "Гал.",
	"eph":             "Еф.",
	"ephesians":       "Еф.",
	"еф":              "Еф.",
	"ефесян":          "Еф.",
	"phil":            "Флп.",
	"philippians":     "Флп.",
	"флп":             "Флп.",
	"филип'ян":        "Флп.",
	"col":             "Кол.",
	"colosians":       "Кол.",
	"кол":             "Кол.",
	"колосян":         "Кол.",
	"1 tim":           "1Тим.",
	"1 timothy":       "1Тим.",
	"1 тим":           "1Тим.",
	"1 тимофія":       "1Тим.",
	"2 tim":           "2Тим.",
	"2 timothy":       "2Тим.",
	"2 тим":           "2Тим.",
	"2 тимофія":       "2Тим.",
	"1 john":          "1Ін.",
	"1 jn":            "1Ін.",
	"1 ін":            "1Ін.",
	"1 іоана":         "1Ін.",
	"2 john":          "2Ін.",
	"2 jn":            "2Ін.",
	"2 ін":            "2Ін.",
	"2 іоана":         "2Ін.",
	"3 john":          "3Ін.",
	"3 jn":            "3Ін.",
	"3 ін":            "3Ін.",
	"3 іона":          "3Ін.",
	"1 pet":           "1Пет.",
	"1 peter":         "1Пет.",
	"1 пет":           "1Пет.",
	"1 петра":         "1Пет.",
	"2 pet":           "2Пет.",
	"2 peter":         "2Пет.",
	"2 петра":         "2Пет.",
	"2 пет":           "2Пет.",
	"heb":             "Євр.",
	"hebrews":         "Євр.",
	"євреїв":          "Євр.",
	"євр":             "Євр.",
	"тита":            "Тит.",
	"тит":             "Тит.",
	"philemon":        "Фил.",
	"phlm":            "Фил.",
	"фил":             "Фил.",
	"филимона":        "Фил.",
	"jas":             "Як.",
	"james":           "Як.",
	"як":              "Як.",
	"якова":           "Як.",
	"jude":            "Іуд.",
	"іуди":            "Іуд.",
	"іуд":             "Іуд.",
	"rev":             "Одкр.",
	"revelation":      "Одкр.",
	"одкр":            "Одкр.",
	"одкровення":      "Одкр.",
}

func NormalizeBookName(name string) string {
	name = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(name, " ")
	name = strings.Replace(name, ".", "", -1)
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	if normalized, ok := BookNames[name]; ok {
		return normalized
	} else {
		return ""
	}
}
