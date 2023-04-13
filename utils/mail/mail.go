package mail

import (
	"fmt"

	"github.com/mailazy/mailazy-go"
	"github.com/takez0o/honestwork-api/api/controller"
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
		fmt.Println(err)
	}

	senderClient := mailazy.NewSenderClient(mc.Key, mc.Secret)
	from := "Job Alert! <no-reply@honestwork.app>"
	subject := fmt.Sprintf("New Applicant for '%s'", job.Title)
	textContent := "You've received a new application for your job listing."
	htmlContent := fmt.Sprintf("<p>Hello <b>%s</b>!</p><br/><br/><p>You've received a new application for your <a href='https://honestwork.app/job/%s/%v'>job listing</a> on HonestWork.<br /><br />It might be the candidate you're looking for!</p><br/><br/><p><a href='https://honestwork.app/creator/%s'>Check their profile now</a></p>", job.Username, recruiter_address, slot, applicant_address)
	req := mailazy.NewSendMailRequest(job.Email, from, subject, textContent, htmlContent)

	senderClient.Send(req)
}

func (mc *MailClient) SendWhitelistMail(recipient_mail string, recipient_address string) {
	senderClient := mailazy.NewSenderClient(mc.Key, mc.Secret)
	from := "Mint Alert! <no-reply@honestwork.app>"
	subject := fmt.Sprintf("Whitelist for '%s'", recipient_address)
	textContent := "You've been whitelisted for HonestWork."
	htmlContent := fmt.Sprintf("<p>Hello <b>%s</b>!</p><br/><br/><p>Congratulations on joining the future of the workforce! Your wallet address(%s) is now on the whitelist and you can mint HonestWork Genesis NFT for free. </p><br/><br/>All you have to do now is head over to <a href='https://www.honestwork.app/mint'>minting page</a> and mint your NFT. Then you can access our platform without any restrictions or limitations. It would be dope if you could also fill out the relevant information on your <a href='https://honestwork.app/profile'>profile</a>. Plus, <a href='https://honestwork.app/profile/skills'>add a skill</a> to join fellow freelancers on HonestWork!<br/><br/>Please take a minute to learn all the cool stuff our platform can do. If you have any questions, ask HNST-4 (AI on our platform) or contact our team on <a href='https://discord.gg/kF6TsVke'>Discord</a>.<br/><br/>Note that this is the alpha version of HonestWork, so it is possible that you may encounter some issues. If you experience any technical difficulties or find that something is not working properly, please inform the team on <a href='https://discord.gg/kF6TsVke'>Discord</a> so we can promptly fix it. And if you have any other feedback, we would love to hear it!<br/><br/>It ain’t much, but it’s HonestWork!<br/><br/>HonestWork Team", recipient_address, recipient_address)
	req := mailazy.NewSendMailRequest(recipient_mail, from, subject, textContent, htmlContent)
	senderClient.Send(req)
}
