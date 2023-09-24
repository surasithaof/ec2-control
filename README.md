# EC2 Control

This repository is contain Go code and scripts to start and stop instances.
The purpose is for schedule start and stop EC2 instances to reduce cost of usage.
To start and stop instances have so many options exmaple aws-cli and lambda function.

## Option 1: aws-cli

[Stop and start your instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html)

Install AWS CLI
[AWS Command Line Interface](https://aws.amazon.com/cli/)

Config AWS credention

```sh
aws configure
```

Start instance

```sh
INSTANCE_ID=<instance_id>

aws ec2 start-instances \
--instance-ids $INSTANCE_ID
```

Stop instance

```sh
INSTANCE_ID=<instance_id>

aws ec2 start-instances \
--instance-ids $INSTANCE_ID
```

can put flag `--dry-run` to check permission without make request. [ec2 cli options](https://docs.aws.amazon.com/cli/latest/reference/ec2/start-instances.html#options).

## Option 2: lambda function

Can use lambda function with event bridge to schedule control EC2 instance.

1. Create an IAM policy and IAM role for your Lambda function.

   ```json
   {
     "Version": "2012-10-17",
     "Statement": [
       {
         "Effect": "Allow",
         "Action": [
           "logs:CreateLogGroup",
           "logs:CreateLogStream",
           "logs:PutLogEvents"
         ],
         "Resource": "arn:aws:logs:*:*:*"
       },
       {
         "Effect": "Allow",
         "Action": ["ec2:Start*", "ec2:Stop*"],
         "Resource": "*"
       }
     ]
   }
   ```

2. Build binary file follow the instructions of [aws-lambda-go](https://github.com/aws/aws-lambda-go).

   ```sh
   GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main *.go
   zip main.zip main
   ```

3. Create AWS lambda function with granted role.
4. upload binary zip file.
5. Runtime settings, set handler file name to build file example `main`.
6. Run test start and stop instances.
7. Can use EventBridge to trigger function for schdule start and stop instances.
