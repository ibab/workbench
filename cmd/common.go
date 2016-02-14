package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetInstances() []*ec2.Instance {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	name := "instance-state-name"
	status := "running"
	filter := ec2.Filter{Name: &name, Values: []*string{&status}}
	resp, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{Filters: []*ec2.Filter{&filter}})
	if err != nil {
		panic(err)
	}
	instances := []*ec2.Instance{}
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			instances = append(instances, inst)
		}
	}
	return instances
}

func GetImages() []*ec2.Image {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	identifier := "self"
	resp, err := svc.DescribeImages(&ec2.DescribeImagesInput{Owners: []*string{&identifier}})
	if err != nil {
		panic(err)
	}
	return resp.Images
}
