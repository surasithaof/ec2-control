# EC2 Control

This repository is contain Go code and scripts to start and stop instances.
The purpose is for schedule start and stop EC2 instances to reduce cost of usage.
To start and stop instances have so many options but for this will have 2 options: one is aws-cli and other one is Go code.

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

can put flag `--dry-run` to check permission without make request. [ec2 cli options](https://docs.aws.amazon.com/cli/latest/reference/ec2/start-instances.html#options)

## Option 2: Go sdk

[Managing Amazon EC2 Instances](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ec2-example-manage-instances.html)

For Go SDK you can run

```sh
go run main.go <state> <instance_id>
```

for `state` option is can be `START` and `STOP` only.
