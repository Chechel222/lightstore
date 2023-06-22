package main

import (
	"chikenShop/email"
	"chikenShop/email/smtp"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

const FROM = "chikenshopferm@gmail.com"
const PASSWORD = "mzwtyiggyvkpcrwy"
const HOST = "smtp.gmail.com"
const SMTPPORT = 587

const EMAILTO = "ttopyc33@gmail.com"
const SUBJECT = "Новое бронирование курочки"
const EMAILTEMPLATEPATH = "email_template.html"

func main() {

	InitRoutes(PORT)
}
func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func InitRoutes(port string) {
	r := gin.Default()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)
	r.POST("/booking", Booking)

	http.ListenAndServe(":"+port, r)
}

type BookingInput struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Chiken      string `json:"chiken"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	PhoneNumber string `json:"phoneNumber"`
}

func Booking(c *gin.Context) {
	var bookingInput BookingInput

	if err := c.BindJSON(&bookingInput); err != nil {
		log.Print(err)
	}
	fmt.Print(bookingInput)

	SendEmail(context.Background(), bookingInput)
}

func SendEmail(ctx context.Context, input BookingInput) error {
	emailSender, err := smtp.NewSMTPSender(FROM, PASSWORD, HOST, SMTPPORT)
	subject := fmt.Sprintf(SUBJECT)
	sendInput := email.SendEmailInput{To: EMAILTO, Subject: subject}
	templateInput := BookingInput{Email: input.Email, Name: input.Name, Chiken: input.Chiken, Date: input.Date, Time: input.Time, PhoneNumber: input.PhoneNumber}

	err = sendInput.GenerateBodyFromHTML(EMAILTEMPLATEPATH, templateInput)
	if err != nil {
		log.Print(err)
		return err
	}

	err = emailSender.Send(sendInput)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
