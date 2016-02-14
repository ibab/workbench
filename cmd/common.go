package cmd

import (
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetSpotZonesAndPrices(instanceType string) ([]*string, []*string) {
	// TODO Rewrite so that we only make one API call (don't fix the AvailabilityZone)
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	zones, err := svc.DescribeAvailabilityZones(nil)
	if err != nil {
		panic(err)
	}
	names := []*string{}
	for _, zone := range zones.AvailabilityZones {
		names = append(names, zone.ZoneName)
	}

	prices := []*string{}
	for _, name := range names {
		params := &ec2.DescribeSpotPriceHistoryInput{
			AvailabilityZone: aws.String(*name),
			MaxResults: aws.Int64(10),
			InstanceTypes: []*string{aws.String(instanceType)},
			ProductDescriptions: []*string{aws.String("Linux/UNIX")},
		}
		resp, err := svc.DescribeSpotPriceHistory(params)
		if err != nil {
			panic(err)
		}
		prices = append(prices, resp.SpotPriceHistory[0].SpotPrice)
	}

	return names, prices
}

func LaunchSpotInstance(price string, zone string, image string, instanceType string) {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	params := &ec2.RequestSpotInstancesInput{
		SpotPrice:             aws.String(price),
		InstanceCount:         aws.Int64(1),
		LaunchSpecification: &ec2.RequestSpotLaunchSpecification{
			ImageId:      aws.String(image),
			InstanceType: aws.String(instanceType),
			Placement: &ec2.SpotPlacement{
				AvailabilityZone: aws.String(zone),
			},
		},
		ValidFrom:  aws.Time(time.Now().Add(10 * time.Second)),
		ValidUntil: aws.Time(time.Now().Add(10 * time.Hour)),
	}
	_, err := svc.RequestSpotInstances(params)
	if err != nil {
		panic(err)
	}
}

func GetInstances() []*ec2.Instance {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	filter := ec2.Filter{
		Name: aws.String("instance-state-name"),
		Values: []*string{aws.String("running"), aws.String("pending")},
	}
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

func TerminateInstance(instance *ec2.Instance) {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	params := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{aws.String(*instance.InstanceId)},
	}
	_, err := svc.TerminateInstances(params)
	if err != nil {
		panic(err)
	}
}
