# EC2 Control

This repository contains Go code and scripts to start and stop instances.
The purpose is to schedule the start and stop of EC2 instances to reduce the cost of usage.
To start and stop instances, there are so many options, such as aws-cli and lambda functions.

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

can put flag `--dry-run` to check permission without making a request. [ec2 cli options](https://docs.aws.amazon.com/cli/latest/reference/ec2/start-instances.html#options).

## Option 2: lambda function

You can use the lambda function with the EventBridge to schedule the control of the EC2 instances.

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

2. Build a binary file by following the instructions of [aws-lambda-go](https://github.com/aws/aws-lambda-go).

   ```sh
   GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main *.go
   zip main.zip main
   ```

3. Create an AWS lambda function with the granted role.
4. Upload a binary zip file.
5. Runtime settings: set handler file name to build file example `main`.
6. Run test start and stop instances.
   example payload

   ```json
   {
     "state": "start",
     "instanceIds": [
       "i-0XXXXX"
     ]
   }
   ```
   
8. Add EventBridge as a function trigger or create the EventBridge schedule from the EventBridge console.
9. Set cron-based schedules; from my example, I will set up two schedules, one for the start instance at 8:00 a.m. and the other for the stop instance at 8:00 p.m. every day.
10. Select target as the AWS Lambda.
11. Select the function to control instances and put a payload for start or stop instances, like when you test the functions.
12. Finish!
