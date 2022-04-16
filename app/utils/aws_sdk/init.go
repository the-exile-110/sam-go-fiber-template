package aws_sdk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"os"
)

var (
	region = os.Getenv("REGION")
)

func CreateAWSSDKFactory() *AWSService {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	secretsmanagerSvc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	return &AWSService{
		SecretManagerSvc: secretsmanagerSvc,
	}
}

type AWSService struct {
	SecretManagerSvc *secretsmanager.SecretsManager
}
