package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

type AWSClient struct {
	ce *costexplorer.CostExplorer
}

func NewAWSClient() *AWSClient {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	return &AWSClient{
		ce: costexplorer.New(sess),
	}
}

func (c *AWSClient) GetCostAndUsage(start, end string) (*costexplorer.GetCostAndUsageOutput, error) {
	params := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Granularity: aws.String("DAILY"),
		Metrics:     []*string{aws.String("AmortizedCost")},
	}
	return c.ce.GetCostAndUsage(params)
}
