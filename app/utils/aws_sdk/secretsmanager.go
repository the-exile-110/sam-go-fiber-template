package aws_sdk

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type RdsSecretData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"dbname"`
}

func (s *AWSService) GetSecret() (*RdsSecretData, error) {

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(os.Getenv("RDS_SECRET_NAME")),
		VersionStage: aws.String("AWSCURRENT"), // default value is AWSCURRENT
	}

	result, err := s.SecretManagerSvc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	var secretData RdsSecretData
	err = json.Unmarshal([]byte(secretString), &secretData)
	if err != nil {
		return nil, err
	}

	return &secretData, nil
}
