package handlers

import (
	"context"
	"fmt"

	"tw-go/models"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("voy a procesar" + ctx.Value(models.Key("oath")).(string) + " > " + ctx.Value(models.Key("method")).(string))
	var r models.RespApi
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("oath")).(string) {

		}
	case "GET":
		switch ctx.Value(models.Key("oath")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("oath")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("oath")).(string) {

		}
	}
	r.Message = "Method Invalid"
	return r
}
