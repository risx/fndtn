package fndtn

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

//FeatureFlag checks AWS Param store for values, just a test for future reference
func FeatureFlag(flag string) string {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	ssmsvc := ssm.New(sess, aws.NewConfig().WithRegion("us-west-2"))
	keyname := flag
	decryption := false
	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &decryption,
	})

	if err != nil {
		//Later, it may be worth syncing configs in another region if failed to
		//pick up in us-west-2
		log.Fatal(err)
		return "Error"
	}

	value := *param.Parameter.Value
	return value
}
