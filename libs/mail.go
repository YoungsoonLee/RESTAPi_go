package libs

import (
	"github.com/astaxie/beego"

	"github.com/matcornic/hermes"

	gomail "gopkg.in/gomail.v2"
)

func MakeMail(email string, emailType string, token string) {
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		//Theme: new(flat),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	var rEmail hermes.Email

	switch emailType {
	case "confirm":
		host := beego.AppConfig.String("frontHost") + "/confirmEmail/" + token

		rEmail = hermes.Email{
			Body: hermes.Body{
				Name: "Jon Snow",
				Intros: []string{
					"Welcome to Hermes! We're very excited to have you on board.",
				},
				Actions: []hermes.Action{
					{
						Instructions: "To get started with Hermes, please click here:",
						Button: hermes.Button{
							Text: "Confirm your account",
							Link: host,
						},
					},
				},
				Outros: []string{
					"Need help, or have questions? Just reply to this email, we'd love to help.",
				},
			},
		}
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(rEmail)
	if err != nil {
		//panic(err) // Tip: Handle error with something else than a panic ;)
		beego.Error("Error MakeMail: ", err)
	}

	/*
		// Optionally, preview the generated HTML e-mail by writing it to a local file
		err = ioutil.WriteFile("preview.html", []byte(emailBody), 0644)
		if err != nil {
			//panic(err) // Tip: Handle error with something else than a panic ;)
			beego.Error("Error MakeMail: ", err)
		}
	*/

	//send email
	sendEmail(email, emailBody)
}

func sendEmail(email string, emailBody string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "youngtip@gmail.com")
	m.SetHeader("To", email)
	//m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", "golang test")
	m.SetBody("text/html", emailBody)
	//m.Attach("template.html") // attach whatever you want

	d := gomail.NewDialer("smtp.gmail.com", 587, beego.AppConfig.String("mail::gmail"), beego.AppConfig.String("mail::gpass"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		//panic(err)
		beego.Error("send email error: ", err)
	}

	beego.Info("success send email to ", email)
}
