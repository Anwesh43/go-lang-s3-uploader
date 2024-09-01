package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileUploader struct {
	filename string
}

func (fu *FileUploader) Upload() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Args[4]),
	}))
	uploader := s3manager.NewUploader(sess)
	myBucket := os.Args[1]
	myKey := os.Args[2]
	f, err := os.Open(fu.filename)
	if err != nil {
		fmt.Println("no such file exists")
	}
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myKey),
		Body:   f,
	})
	if err != nil {
		fmt.Println("Error uploading file", err)
	} else {
		fmt.Println(aws.StringValue((&result.Location)))
	}
}

func main() {
	fmt.Println(os.Args[1:])
	fu := FileUploader{
		filename: os.Args[3],
	}
	fu.Upload()
}
