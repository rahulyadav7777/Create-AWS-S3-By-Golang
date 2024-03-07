package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Set your AWS credentials (replace these with your own credentials or use environment variables).
	awsAccessKey := "AKIAZQ3DTBCIU3SQK4FR"
	awsSecretKey := "JSFWFqa5amEqcWa59xmNsGptKA1Zf2WGcmq4eFMo"
	awsRegion := "us-east-1" // Replace with your preferred AWS region.

	// Create a new AWS session.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create an S3 service client.
	s3Client := s3.New(sess)

	// Specify the name of the S3 bucket you want to create.
	bucketName := "yrahulawss3bucket"

	// Create the S3 bucket.
	_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("Error creating S3 bucket:", err)
		return
	}

	fmt.Printf("S3 bucket '%s' created successfully.\n", bucketName)
}
