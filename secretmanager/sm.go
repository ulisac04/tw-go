package secretmanager

import (
	"encoding/json"
	"fmt"
	"tw-go/awsgo"
	"tw-go/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("> pido secreto " + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Ctg)
	claves, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}
	json.Unmarshal([]byte(*claves.SecretString), &datosSecret)
	fmt.Println("lectura secreto ok " + secretName)
	return datosSecret, nil
}
