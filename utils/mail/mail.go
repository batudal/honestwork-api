package mail

import (
	"fmt"

	"github.com/mailazy/mailazy-go"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
)

type MailClient struct {
	Key    string
	Secret string
}

func NewMailClient(key string, secret string) *MailClient {
	return &MailClient{
		Key:    key,
		Secret: secret,
	}
}

func (mc *MailClient) sendNewApplicantMail(recruiter_address string, slot int, applicant_address string) {
	job_controller := controller.NewJobController(recruiter_address, slot)
	job, err := job_controller.GetJob()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
	}

	senderClient := mailazy.NewSenderClient(mc.Key, mc.Secret)
	from := "Job Alert! <no-reply@honestwork.app>"
	subject := fmt.Sprintf("New Applicant for '%s'", job.Title)
	textContent := "You've received a new application for your job listing."
	htmlContent := fmt.Sprintf("<p>Hello <b>%s</b>!</p><br/><br/><p>You've received a new application for your <a href='https://honestwork.app/job/%s/%v'>job listing</a> on HonestWork.<br /><br />It might be the candidate you're looking for!</p><br/><br/><p><a href='https://honestwork.app/creator/%s'>Check their profile now</a></p>", job.Username, recruiter_address, slot, applicant_address)
	req := mailazy.NewSendMailRequest(job.Email, from, subject, textContent, htmlContent)

	senderClient.Send(req)
}
