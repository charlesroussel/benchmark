module github.com/Ogury/profiling

require (
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/aws/aws-dax-go v1.2.8
	github.com/aws/aws-sdk-go-v2 v1.9.1
	github.com/aws/aws-sdk-go-v2/config v1.8.2
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.2.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.5.1
	github.com/gofiber/fiber/v2 v2.18.0
	github.com/klauspost/compress v1.13.5 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
)

replace github.com/aws/aws-dax-go => github.com/kochie/aws-dax-go v1.2.8-0.20210606022114-926ae2149af1

go 1.16
