package main

import (
	"context"
	"os"
	"strings"
	"tw-go/awsgo"
	"tw-go/bd"
	"tw-go/handlers"
	"tw-go/models"
	"tw-go/secretmanager"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLamdda)
}

func EjecutoLamdda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InicializamosAWS()

	if !validoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en Vatiables de entorno",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en Vatiables de entorno",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["tw-go"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("bucketName"))
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)

	err = bd.ConectarBD(awsgo.Ctx)

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error en conexion a BD" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	respApi := handlers.Manejadores(awsgo.Ctx, request)

	if respApi.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       respApi.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
	} else {
		return respApi.CustomResp, nil
	}
}

func validoParametros() bool {
	_, traeParams := os.LookupEnv("SecretName")
	if !traeParams {
		return traeParams
	}
	_, traeParams = os.LookupEnv("BucketName")
	if !traeParams {
		return traeParams
	}
	_, traeParams = os.LookupEnv("UrlPrefix")
	if !traeParams {
		return traeParams
	}

	return traeParams

}
