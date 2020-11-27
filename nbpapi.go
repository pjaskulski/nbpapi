package nbpapi

func init() {
	SetLang("en")
}

// SetLang function
func SetLang(lang string) {
	if lang == "pl" {
		l = langTexts["pl"]
	} else if lang == "en" {
		l = langTexts["en"]
	}
}
