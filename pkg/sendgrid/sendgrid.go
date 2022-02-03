package sendgrid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/go-playground/validator.v9"
)

// SendEmail is the struct so for the email to be sent
type SendEmail struct {
	SenderName       string `json:"senderName,omitempty" bson:"senderName,omitempty" validate:"required,min=3,max=128"`
	SenderEmail      string `json:"senderEmail" bson:"senderEmail" validate:"required,min=10,max=128"`
	Subject          string `json:"subject" bson:"subject" validate:"required,min=10,max=128"`
	RecipientName    string `json:"recipientName" bson:"recipientName" validate:"required,min=3,max=128"`
	RecipientEmail   string `json:"recipientEmail" bson:"recipientEmail" validate:"required,min=10,max=128"`
	PlainTextContent string `json:"plainTextContent" bson:"plainTextContent" validate:"required,min=10,max=512"`
	HTMLContent      string `json:"htmlContent" bson:"htmlContent" validate:"min=10,max=512"`
	Code             string `json:"code,omitempty" bson:"code,omitempty" validate:"omitempty,min=1,max=64"`
	Template         string `json:"template" bson:"template"`
}

// EmailValidator - This will validate the user using the structs annotations
type EmailValidator struct {
	validator *validator.Validate
}

// Validate - This validates the request body
func (u *EmailValidator) Validate(i interface{}) error {
	return u.validator.Struct(i)
}

// ROUTES --------------------------------------------------------

func configureCustomRoutes() {
	e.POST("/send-email", sendEmail)
	e.POST("/auth/confirm-account", sendConfirmAccount)
	e.POST("/auth/reset-password", sendResetPassword)

}

// FUNCTIONS ------------------------------------------------------

func sendEmail(c echo.Context) error {

	var sendEmail SendEmail

	c.Echo().Validator = &EmailValidator{validator: v}

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	fmt.Println(string(bodyBytes))

	err := json.Unmarshal(bodyBytes, &sendEmail)
	if err != nil {
		fmt.Println("error:", err)
	}

	from := mail.NewEmail(sendEmail.SenderName, sendEmail.SenderEmail)
	subject := sendEmail.Subject
	to := mail.NewEmail(sendEmail.RecipientName, sendEmail.RecipientEmail)
	plainTextContent := sendEmail.PlainTextContent
	htmlContent := sendEmail.HTMLContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SendGridAPI)

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusNotFound, "error")

	}

	return c.JSON(http.StatusAccepted, response)

}

func sendResetPassword(c echo.Context) error {

	var sendEmail SendEmail

	c.Echo().Validator = &EmailValidator{validator: v}

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	fmt.Println(string(bodyBytes))

	err := json.Unmarshal(bodyBytes, &sendEmail)
	if err != nil {
		fmt.Println("error:", err)
	}

	from := mail.NewEmail(sendEmail.SenderName, sendEmail.SenderEmail)
	subject := sendEmail.Subject
	to := mail.NewEmail(sendEmail.RecipientName, sendEmail.RecipientEmail)
	plainTextContent := ""

	htmlContent := resetPasswordTemplateStart + "http://127.0.0.1:4200/auth/change-password/" + sendEmail.RecipientEmail + "/" + sendEmail.Code + "/" + resetPasswordTemplateEnd
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SendGridAPI)

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusNotFound, "error")

	}

	return c.JSON(http.StatusAccepted, response)

}

func sendConfirmAccount(c echo.Context) error {

	fmt.Println("STARTING confirm account email")

	var sendEmail SendEmail

	c.Echo().Validator = &EmailValidator{validator: v}

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	fmt.Println(string(bodyBytes))

	err := json.Unmarshal(bodyBytes, &sendEmail)
	if err != nil {
		fmt.Println("error:", err)
	}

	from := mail.NewEmail(sendEmail.SenderName, sendEmail.SenderEmail)
	subject := sendEmail.Subject
	to := mail.NewEmail(sendEmail.RecipientName, sendEmail.RecipientEmail)
	plainTextContent := ""

	htmlContent := welcomeTemplateStart + "http://127.0.0.1:8080/auth/confirm-email/" + sendEmail.RecipientEmail + "/" + sendEmail.Code + "/" + welcomeTemplateEnd
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SendGridAPI)

	fmt.Println(message)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusNotFound, "error")

	}

	return c.JSON(http.StatusAccepted, response)
}
