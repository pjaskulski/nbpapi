package nbpapi

// SetLang function (language for output functions)
func setLang(lang string) {
	if lang == "pl" {
		l = langTexts["pl"]
	} else if lang == "en" {
		l = langTexts["en"]
	}
}
