# sam-go-fiber-template
go fiber template on AWS SAM

## What is it?

This is the go fiber Serverless template based on AWS SAM.

Libraries:
- go fiber
- aws-lambda-go
- aws-sdk-go
- aws-lambda-go-api-proxy
- go-json
- gorm

Features:
- Deployment with Zip package
- Dev/Prod environment deployment
- Domain name support
- Vpc, Subnet, SecurityGroup support
- Get RDS authentication data with AWS Secrets Manager

## Installation

1. Install AWS CLI:

```bash
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

2. Install AWS SAM CLI:

```bash
brew tap aws/tap
brew install aws-sam-cli
```

## Setup

Edit the samconfig.toml file and write the required settings for the project.

| Key             | Description                                                                                                                                                    |
|-----------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| stack_name      | The name of the stack                                                                                                                                          |
| s3_bucket       | Specify S3 bucket for deployment                                                                                                                               |
| Region          | Specify AWS region                                                                                                                                             |
| Environment     | Specify the environment                                                                                                                                        |
| ProjectName     | Specify the project name                                                                                                                                       |
| DevDomain       | Specify the development domain                                                                                                                                 |
| ProdDomain      | Specify the production domain                                                                                                                                  |
| Certifi0cateArn | Find and copy the arn of the domain certificate in AWS Certificate Manager                                                                                     |
| HostedZoneId    | Find and copy the HostedZoneId of the domain in AWS Route53                                                                                                    |
| VpcId           | Find and copy the VpcId                                                                                                                                        |
| SubnetIds       | SubnetIds, split by comma                                                                                                                                      |
| RDS_SECRET_NAME | Name of the RDS key created in AWS Secrets Manager, refer to the [document](https://docs.aws.amazon.com/secretsmanager/latest/userguide/managing-secrets.html) |


## Local start 

Edit and use the local-start.sh script to start the project locally


## Deploy

Edit and use the deploy.sh script to deploy the project to AWS