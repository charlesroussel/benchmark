package connectors

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-dax-go/dax"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	dax_cfg.Region = awsRegion
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
	_id string `dynamodbav:_id`
}

type ExampleValue struct {
	value  string `dynamodbav:"value"`
	value2 int    `dynamodbav:"value2"`
}

func (c *DynamoService) Get(tableName string, id string, isDax bool) (string, error) {
	var itemInput = &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"_id": &types.AttributeValueMemberS{Value: id},
		},
	}
	var out *dynamodb.GetItemOutput
	var err error
	if isDax {
		out, err = c.dax_client.GetItem(context.TODO(), itemInput)
	} else {
		out, err = c.client.GetItem(context.TODO(), itemInput)
	}

	if err != nil {
		return "", err
	}

	if out.Item == nil {
		return "", errors.New("not found")
	}

	v := ExampleValue{}
	attributevalue.Unmarshal(out.Item["value"], &v.value)
	attributevalue.Unmarshal(out.Item["value2"], &v.value2)
	if err != nil {
		return "", err
	}
	return v.value, err
}
