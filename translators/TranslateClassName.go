package translators

var classNameTranslationMap = map[string]string{
	"Cavaleira Negra": "Dark Knight",
	"Corsária":        "Corsair",
	"Domadora":        "Tamer",
	"Feiticeira":      "Sorceress",
	"Guardiã":         "Guardian",
	"Guerreiro":       "Warrior",
	"Lutador":         "Striker",
	"Maga":            "Witch",
	"Mago":            "Wizard",
	"Mística":         "Mystic",
	"Musah":           "Musa",
	"Sábio":           "Sage",
	"Sagitário":       "Archer",
	"Valquíria":       "Valkyrie",
}

func TranslateClassName(className *string) {
	if val, ok := classNameTranslationMap[*className]; ok {
		*className = val
	}
}
