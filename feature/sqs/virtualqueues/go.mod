module github.com/aws/aws-sdk-go-v2/feature/sqs/virtualqueues

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.21.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.24.7 // indirect
)

replace github.com/aws/aws-sdk-go-v2 => ../../../
replace github.com/aws/aws-sdk-go-v2/service/sqs => ../../../service/sqs
