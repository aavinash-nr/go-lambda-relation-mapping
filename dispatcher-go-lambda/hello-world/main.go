// package main

// import (
// 	"context"
// 	"fmt"
// 	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
// 	"github.com/newrelic/go-agent/v3/newrelic"
// )

// func handler(ctx context.Context) (string, error) {
// 	// At this point, we're handling an invocation. Cold start is over; this code runs for each invocation.
// 	// We'd like to add a custom event, and a custom attribute. For that, we need the transaction.
// 	if txn := newrelic.FromContext(ctx); nil != txn {
// 		// This is an example of a custom event. `FROM MyGoEvent SELECT *` in New Relic will find this event.
// 		txn.Application().RecordCustomEvent("MyGoEvent", map[string]interface{}{
// 			"zip": "zap",
// 		})

// 		// This attribute gets added to the normal AwsLambdaInvocation event
// 		txn.AddAttribute("customAttribute", "customAttributeValue")
// 	}

// 	// As normal, anything you write to stdout ends up in CloudWatch
// 	fmt.Println("hello world!, from dispatcher!!!")

// 	return "Success!", nil
// }

// func main() {
// 	// Here we are in cold start. Anything you do in main happens once.
// 	// In main, we initialize the agent.
// 	app, err := newrelic.NewApplication(nrlambda.ConfigOption())
// 	if nil != err {
// 		fmt.Println("error creating app (invalid config):", err)
// 	}
// 	// Then we start the lambda handler using `nrlambda` rather than `lambda`
// 	nrlambda.Start(handler, app)
// }

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func handler(ctx context.Context) (string, error) {

	if txn := newrelic.FromContext(ctx); nil != txn {
		txn.Application().RecordCustomEvent("MyGoEvent", map[string]interface{}{
			"zip": "zap",
		})
		txn.AddAttribute("customAttribute", "customAttributeValue")
	}

	
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
		return "", err
	}

	
	svc := lambda.NewFromConfig(cfg)

	
	functionName := "arn:aws:lambda:us-east-1:466768951184:function:converter-go-lambda-HelloWorldFunction-IUVyQYeYJLuN" 
	payload := []byte(`{"key": "value"}`)      

	input := &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
	}

	result, err := svc.Invoke(ctx, input)
	if err != nil {
		log.Printf("Failed to invoke lambda function: %v", err)
		return "", err
	}

	fmt.Printf("Function response: %s\n", string(result.Payload))
	fmt.Println("hello world!, from dispatcher!!!")

	return "Success!", nil
}

func main() {
	app, err := newrelic.NewApplication(nrlambda.ConfigOption())
	if nil != err {
		fmt.Println("error creating app (invalid config):", err)
	}

	nrlambda.Start(handler, app)
}