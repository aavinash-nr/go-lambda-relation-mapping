import boto3
import json

def lambda_handler(event, context):
    # Initialize a boto3 client for Lambda
    lambda_client = boto3.client('lambda')

    # Specify the name of the Lambda function you want to invoke
    target_function_name = 'arn:aws:lambda:us-east-1:466768951184:function:python-function-reciever-HelloWorldFunction-vmpC0rTcVBoR'

    # Prepare the payload to send to the target Lambda function
    payload = {
        'key1': 'value1',
        'key2': 'value2'
    }

    # Invoke the target Lambda function
    try:
        response = lambda_client.invoke(
            FunctionName=target_function_name,
            InvocationType='RequestResponse',  # Synchronous call
            Payload=json.dumps(payload)
        )

        # Decode the response
        response_payload = json.loads(response['Payload'].read())
        print("Response from invoked Lambda function:", response_payload)

        return {
            'statusCode': 200,
            'body': json.dumps('Invocation Successful!')
        }

    except Exception as e:
        print(f"Error invoking the target Lambda function: {e}")
        return {
            'statusCode': 500,
            'body': json.dumps('Invocation Failed!')
        }