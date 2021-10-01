package connectors

import (
	"context"
	"os"

	"github.com/aws/aws-dax-go/dax"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoClient DynamoService
var secureEndpoint string

type DynamoService struct {
	client     *dynamodb.Client
	dax_client *dax.Dax
}

func init() {
	awsRegion := os.Getenv("AWS_REGION")
	daxEndpoint := os.Getenv("DAX_ENDPOINT")

	dax_cfg := dax.DefaultConfig()
	dax_cfg.HostPorts = []string{daxEndpoint}
	dax_cfg.Region = "us-west-2"
	dax_client, err := dax.New(dax_cfg)

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(lo *config.LoadOptions) error {
		lo.Region = awsRegion
		return nil
	})

	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(cfg)
	DynamoClient = DynamoService{
		client:     client,
		dax_client: dax_client,
	}
}

type ExampleKey struct {
	key string `dynamodbav:"key" json:"key"`
}

type ExampleValue struct {
	value string
}

func (c *DynamoService) Get(tableName string, id string, isDax bool) (string, error) {
	key := ExampleKey{key: id}

	avs, err := attributevalue.MarshalMap(key)
	if err != nil {
		return "", err
	}

	var itemInput = &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       avs,
	}
	var out *dynamodb.GetItemOutput
	if isDax {
		out, err = c.client.GetItem(context.TODO(), itemInput)
	} else {
		out, err = c.dax_client.GetItem(context.TODO(), itemInput)
	}

	value := ExampleValue{}
	err = attributevalue.UnmarshalMap(out.Item, &value)
	return value.value, err
}
