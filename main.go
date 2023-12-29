package main

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

const password string = "admin"

func main() {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	secretData := map[string]interface{}{
		"password": password,
	}

	ctx := context.Background()
	_, err = client.KVv2("secret").Put(ctx, "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("Unable to write secret: %v to the vault", err)
	}
	log.Println("Super secret password written successfully to the vault.")

	secret, err := client.KVv2("secret").Get(ctx, "my-secret-password")
	if err != nil {
		log.Fatalf(
			"Unable to read the super secret password from the vault: %v",
			err,
		)
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf(
			"value type assertion failed: %T %#v",
			secret.Data["password"],
			secret.Data["password"],
		)
	}

	log.Printf("Super secret password [%s] was retrieved.\n", value)
}
