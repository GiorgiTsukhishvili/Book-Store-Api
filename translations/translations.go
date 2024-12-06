package translations

type EmailTranslations map[string]map[string]string

var EmailTranslationData = EmailTranslations{
	"en": {
		"account-verification":    "Account Verification",
		"name":                    "BookShelf",
		"joining-text":            "Thanks for joining BookShelf! We really appreciate it. Please click the button below to verify your account:",
		"verify-button":           "Verify account",
		"clicking":                "If clicking doesn`t work, you can try copying and pasting it to your browser:",
		"any-problem":             "If you have any problems, please contact us:",
		"crew":                    "BookQuotes Crew",
		"hello":                   "Hello",
		"password-reset":          "Password Reset",
		"reset-text":              "Please click the button below to reset your password:",
		"reset-button":            "Reset password",
		"email-not-verified":      "Email is not verified",
		"email-verification":      "Email Verification",
		"email-verification-text": "Thanks for adding email! We really appreciate it. Please click the button below to verify your email:",
		"email-verify-button":     "Verify email",
	},
	"ka": {
		"account-verification":    "ანგარიშის დადასტურება",
		"name":                    "წიგნების თარო",
		"joining-text":            "გმადლობთ, რომ შეუერთდით წიგნების თაროსს! ჩვენ ნამდვილად ვაფასებთ მას. გთხოვთ, დააჭიროთ ქვემოთ მოცემულ ღილაკს თქვენი ანგარიშის დასადასტურებლად:",
		"verify-button":           "გააქტიურე ანგარიში",
		"clicking":                "თუ დაწკაპუნება არ მუშაობს, შეგიძლიათ სცადოთ მისი კოპირება და ჩასმა თქვენს ბრაუზერში:",
		"any-problem":             "თუ რაიმე პრობლემა გაქვთ, გთხოვთ დაგვიკავშირდეთ:",
		"crew":                    "წიგნების თაროს გუნდი",
		"hello":                   "გამარჯობა",
		"password-reset":          "პაროლის რედაქტირება",
		"reset-text":              "გთხოვთ, დააჭიროთ ქვემოთ მოცემულ ღილაკს თქვენი პაროლის შესაცვლელად:",
		"reset-button":            "პაროლის რედაქტირება",
		"email-not-verified":      "მეილი არ არის ვერიფიცირებული",
		"email-verification":      "მეილის ვერიფიკაცია",
		"email-verification-text": "გმადლობთ, რომ დაამატეთ დამატებითი მეილი. გთხოვთ, დააჭიროთ ქვემოთ მოცემულ ღილაკს თქვენი მეილის დასადასტურებლად:",
		"email-verify-button":     "დაადასტურე მეილი",
	},
}

func GetTranslation(lang, key string) string {
	if translations, exists := EmailTranslationData[lang]; exists {
		if value, exists := translations[key]; exists {
			return value
		}
	}

	if value, exists := EmailTranslationData["en"][key]; exists {
		return value
	}

	return ""
}
