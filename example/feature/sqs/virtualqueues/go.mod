module github.com/aws/aws-sdk-go-v2/example/feature/sqs/virtualqueues

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.21.2
	github.com/aws/aws-sdk-go-v2/config v1.21.2 
) 

replace github.com/aws/aws-sdk-go-v2 => ../../../../

replace github.com/aws/aws-sdk-go-v2/config => ../../../../config/

replace github.com/aws/aws-sdk-go-v2/credentials => ../../../../credentials/

