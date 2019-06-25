package git

import (
	"log"
	"o2clock/api-proto/webhooks/git"
)

func SaveGithubPushWebhookInfo(req *githubpb.GithubPushWebhookRequest) error {
	log.Println("Data From Github:", req)
	return nil
}
