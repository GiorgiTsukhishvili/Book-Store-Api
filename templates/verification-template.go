package templates

import (
	"fmt"
	"os"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/translations"
)

func VerificationEmailTemplate(
	lang string,
	code string,
	name string,
	mainText string,
	buttonText string,
) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Email Verification</title>
    <style>
        body {
            margin: 0;
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            color: #333;
            padding: 0;
        }

        .email-container {
            max-width: 600px;
            margin: 0 auto;
            background: #ffffff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }

        .email-header {
            text-align: center;
            margin-bottom: 20px;
        }

        .email-header img {
            height: 50px;
        }

        .email-body {
            text-align: center;
        }

        .email-body h1 {
            color: #333;
            font-size: 24px;
            margin-bottom: 10px;
        }

        .email-body p {
            font-size: 16px;
            margin-bottom: 20px;
            line-height: 1.5;
        }

        .email-button {
            display: inline-block;
            background-color: #007BFF;
            color: #ffffff;
            padding: 10px 20px;
            text-decoration: none;
            border-radius: 5px;
            font-size: 16px;
        }

        .email-footer {
            text-align: center;
            margin-top: 30px;
            font-size: 14px;
            color: #888;
        }

        .email-footer a {
            color: #007BFF;
            text-decoration: none;
        }

        @media (max-width: 600px) {
            .email-container {
                padding: 15px;
            }

            .email-body p {
                font-size: 14px;
            }

            .email-button {
                padding: 8px 15px;
                font-size: 14px;
            }
        }
    </style>
</head>

<body>
    <div class="email-container">
        <div class="email-header">
            <img src="%s/imgs/logo.png" alt="BookShelf Logo">
        </div>

        <div class="email-body">
            <h1>%s %s,</h1>
            <p>%s</p>
            <a class="email-button">%s</a>
        </div>

        <div class="email-footer">
            <p>%s <a href="mailto:support@bookquotes.ge">support@bookquotes.ge</a></p>
            <p>%s</p>
        </div>
    </div>
</body>

</html>`,
		os.Getenv("APP_URL"),
		translations.GetTranslation(lang, "hello"), name,
		mainText,
		code,
		translations.GetTranslation(lang, "any-problem"),
		translations.GetTranslation(lang, "crew"),
	)
}
