package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	fmt.Println(session, err)
	s3Obj := s3.New(session)
	// s3Obj.ListBuckets()
	// s3Obj.CreateBucket()
	//s3Obj.DeleteObject() //delete object in bucket
	// s3Manager.NewUploader(session)   //upload file

	//can define bucket policies, ACL(Access control list)

}
