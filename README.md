# AWS Lambda Hello World

This is a simple AWS Lambda function that logs "Hello, World!" to the console.

## Prerequisites

Before running this project, make sure you have the following installed:

- Node.js
- npm
- Serverless framework
- Docker
- Docker Compose

## Getting Started

1. Clone this repository.

2. Install the dependencies by running the following command:

   ```
   npm install
   ```

3. Deploy the AWS Lambda function using the Serverless framework:

   ```
   serverless deploy
   ```

4. Once the deployment is complete, you can test the function by invoking it:

   ```
   serverless invoke -f helloWorld
   ```

   The function will log "Hello, World!" to the console.

## Local Development with LocalStack

To run the project locally using LocalStack, follow these steps:

1. Start LocalStack using Docker Compose:

   ```
   docker-compose up
   ```

2. Update the `serverless.yml` file with the appropriate endpoint URLs for the AWS services you want to use with LocalStack.

3. Deploy the AWS Lambda function to LocalStack:

   ```
   serverless deploy --stage local
   ```

4. Test the function by invoking it:

   ```
   serverless invoke -f helloWorld --stage local
   ```

   The function will log "Hello, World!" to the console.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more information.# aws-lambda-localstack
