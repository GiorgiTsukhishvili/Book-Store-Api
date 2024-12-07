package translations

type EmailTranslations map[string]map[string]string

var EmailTranslationData = EmailTranslations{
	"en": {
		"account-verification":    "Account Verification",
		"name":                    "BookShelf",
		"joining-text":            "Thanks for joining BookShelf! We really appreciate it. Please use the 4-digit code below to verify your account:",
		"verify-button":           "Verify account",
		"clicking":                "If you encounter any issues using the code, please contact us:",
		"any-problem":             "If you have any problems, please contact us:",
		"crew":                    "BookQuotes Crew",
		"hello":                   "Hello",
		"password-reset":          "Password Reset",
		"reset-text":              "Please use the 4-digit code below to reset your password:",
		"reset-button":            "Reset password",
		"email-not-verified":      "Email is not verified",
		"email-verification":      "Email Verification",
		"email-verification-text": "Thanks for adding your email! We really appreciate it. Please use the 4-digit code below to verify your email:",
		"email-verify-button":     "Verify email",
	},
	"ka": {
		"account-verification":    "ანგარიშის დადასტურება",
		"name":                    "წიგნების თარო",
		"joining-text":            "გმადლობთ, რომ შეუერთდით წიგნების თაროსს! ჩვენ ნამდვილად ვაფასებთ მას. გთხოვთ, გამოიყენოთ ქვემოთ მოცემული 4-ნიშნა კოდი თქვენი ანგარიშის დასადასტურებლად:",
		"verify-button":           "გააქტიურე ანგარიში",
		"clicking":                "თუ კოდის გამოყენებისას რაიმე პრობლემა შეგექმნებათ, დაგვიკავშირდით:",
		"any-problem":             "თუ რაიმე პრობლემა გაქვთ, გთხოვთ დაგვიკავშირდეთ:",
		"crew":                    "წიგნების თაროს გუნდი",
		"hello":                   "გამარჯობა",
		"password-reset":          "პაროლის რედაქტირება",
		"reset-text":              "გთხოვთ, გამოიყენოთ ქვემოთ მოცემული 4-ნიშნა კოდი თქვენი პაროლის შესაცვლელად:",
		"reset-button":            "პაროლის რედაქტირება",
		"email-not-verified":      "მეილი არ არის ვერიფიცირებული",
		"email-verification":      "მეილის ვერიფიკაცია",
		"email-verification-text": "გმადლობთ, რომ დაამატეთ დამატებითი მეილი. გთხოვთ, გამოიყენოთ ქვემოთ მოცემული 4-ნიშნა კოდი თქვენი მეილის დასადასტურებლად:",
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
