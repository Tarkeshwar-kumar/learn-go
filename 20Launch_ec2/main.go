package main

import (
	"context"
	"fmt"
	"strings"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)
func main() {
	ctx := context.Background()
	InstanceId, err := CreateEC2(ctx, "ap-south-1")
	if err != nil {
		fmt.Println("Failed to create ec2 instance ")
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("Intsance id of created instance is ", InstanceId)
}

func CreateEC2(ctx context.Context, region string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
    if err != nil {
        return "", fmt.Errorf("unable to load SDK config, %v", err)
    }
	ec2Client := ec2.NewFromConfig(cfg)
	existingKeyPairs, err := ec2Client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		KeyNames: []string{"go-aws-sdk"},
	})
	if err != nil && !strings.Contains(err.Error(), "InvalidKeyPair.NotFound") {
		return "", fmt.Errorf("DescribeKeyPairs error: %s", err)
	}
	if existingKeyPairs == nil || len(existingKeyPairs.KeyPairs) == 0 {
		_, err = ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
			KeyName: aws.String("go-aws-sdk"),
		})
		if err != nil {
			return "", fmt.Errorf("CreateKeyPair error, %v", err)
		}
	}
	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name: aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"},
			},
			{
				Name: aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})
	if err != nil {
        return "", fmt.Errorf("DescribeImages error, %v", err)
    }
	if len(imageOutput.Images) == 0 {
		return "", fmt.Errorf("noo image found")
	}

	fmt.Println("Image id is: ", imageOutput.Images[0].ImageId)

	runInstance, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      imageOutput.Images[0].ImageId,
		InstanceType: types.InstanceTypeT2Micro,
		KeyName:      aws.String("go-aws-sdk"),
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})
	if err != nil {
		return "", fmt.Errorf("RunInstance error: %s", err)
	}

	if len(runInstance.Instances) == 0 {
		return "", fmt.Errorf("RunInstance has empty length (%d)", len(runInstance.Instances))
	}

	return *runInstance.Instances[0].InstanceId, nil
}