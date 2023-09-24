package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func startInstance(svc *ec2.EC2, instanceIDs []string) error {
	instances := []*string{}
	for _, id := range instanceIDs {
		instances = append(instances, aws.String(id))
	}
	// We set DryRun to true to check to see if the instance exists and we have the
	// necessary permissions to monitor the instance.
	input := &ec2.StartInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(true),
	}
	_, err := svc.StartInstances(input)
	awsErr, ok := err.(awserr.Error)

	// If the error code is `DryRunOperation` it means we have the necessary
	// permissions to Start this instance
	if ok && awsErr.Code() == "DryRunOperation" {
		// Let's now set dry run to be false. This will allow us to start the instances
		input.DryRun = aws.Bool(false)
		result, err := svc.StartInstances(input)
		if err != nil {
			log.Println("Error", err)
			return err
		} else {
			log.Println("Success", result.StartingInstances)
			return nil
		}
	} else { // This could be due to a lack of permissions
		log.Println("Error", err)
		return err
	}
}

func stopInstance(svc *ec2.EC2, instanceIDs []string) error {
	instances := []*string{}
	for _, id := range instanceIDs {
		instances = append(instances, aws.String(id))
	}
	input := &ec2.StopInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(true),
	}
	_, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err := svc.StopInstances(input)
		if err != nil {
			log.Println("Error", err)
			return err
		} else {
			log.Println("Success", result.StoppingInstances)
			return nil
		}
	} else {
		log.Println("Error", err)
		return err
	}
}
