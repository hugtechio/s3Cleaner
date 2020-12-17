package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var filePath string = "buckets.txt"
var region string = "us-east-1"

func deleteBucket(bucket *string) (ret bool) {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := s3.New(sess)

	objects, error := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: bucket,
	})
	fmt.Println(objects, error)

	for i, v := range objects.Contents {
		fmt.Println(i, v)
		res, error := svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: bucket,
			Key:    v.Key,
		})
		fmt.Println(res, error)
	}

	versions, _ := svc.ListObjectVersions(&s3.ListObjectVersionsInput{
		Bucket: bucket,
	})

	for i, v := range versions.Versions {
		fmt.Println(i, v)
		res, error := svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket:    bucket,
			Key:       v.Key,
			VersionId: v.VersionId,
		})
		fmt.Println(res, error)
	}

	for i, v := range versions.DeleteMarkers {
		fmt.Println(i, v)
		res, error := svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket:    bucket,
			Key:       v.Key,
			VersionId: v.VersionId,
		})
		fmt.Println(res, error)
	}

	res, error := svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: bucket,
	})
	fmt.Println(res, error)

	return true
}

func main() {
	fmt.Println(os.Args)

	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}

	if len(os.Args) == 3 {
		filePath = os.Args[1]
		region = os.Args[2]
	}

	fmt.Println("[File of bucket list]: " + filePath)
	fmt.Println("[S3 region]: " + region)
	f, error := os.Open(filePath)
	if error != nil {
		fmt.Println("File open error")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		deleteBucket(&text)
	}
}
