package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func sendSubmittMessage(name string, email string, code string) {

	Sender := "givaway@support-pp.de"
	Recipient := email
	Subject := "[Support++] Gewinnspiel Teilnahmebestätigung!"
	HtmlBody := "Hey " + name + ", <br><br> mit dieser E-Mail möchten wir deine Teilname bestätigen! <br> Wir drücken dir die Daumen, am ??.??.2018 erfährst du ob due einer der glückelichen bist.<br><br> Ein kleiner Schritt fehlt jedoch noch! <br>Bitte rufe die komplett kostenfreie Telefonummer <b>800 4030172</b> anrufen und diesen Code eingeben: <code>" + code + "</code> !<br><br> Bitte beachte hierzu auch die<a href='https://givaway.support-pp.de/agb'>Teilnahmebedinngungen</a> <br><br> Dein Pirat John und das Support++ Team."
	CharSet := "UTF-8"

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("aws_id"), os.Getenv("aws_secert"), ""),
	},
	)

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Email Sent to address: " + Recipient)
	fmt.Println(result)
}
