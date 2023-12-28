# hashicorp-vault

HashiCorp Vault is a powerful tool for managing secrets and protecting sensitive data. This documentation provides a deep dive into various aspects of HashiCorp Vault, from setup to advanced features.

## Table of Contents

1. [**Vault Setup & Basic Commands**](#1-vault-setup--basic-commands)
2. [**Vault Secret Engine & Dynamic Secret Generation**](#2-vault-secret-engine--dynamic-secret-generation)
3. [**Vault Policy & Token Authentication**](#3-vault-policy--token-authentication)
4. [**Manage Go Application Secrets Using Vault**](#4-manage-go-application-secrets-using-vault)

---

## 1. Vault Setup & Basic Commands

### 1.1 Installation

Follow these steps to install HashiCorp Vault:

```bash
# Example commands for Linux
sudo apt-get update
sudo apt-get install vault
```

### 1.2 Initialization & Unsealing

Initialize Vault and unseal it using the following commands:

```bash
vault operator init
vault operator unseal [UNSEAL_KEY]
```

...

---

## 2. Vault Secret Engine & Dynamic Secret Generation

### 2.1 Enable Secret Engine

Enable a secret engine to manage secrets:

```bash
vault secrets enable -path=secret/ kv
```

### 2.2 Dynamic Secret Generation

Generate dynamic secrets for a database:

```bash
vault write database/config/mydb \
    plugin_name=mysql-database-plugin \
    connection_url="{{username}}:{{password}}@tcp(localhost:3306)/" \
    allowed_roles="my-role" \
    username="root" \
    password="root"
```

...

---

## 3. Vault Policy & Token Authentication

### 3.1 Create Policy

Define a policy to manage access:

```bash
vault policy write my-policy - <<EOF
path "secret/data/myapp/*" {
  capabilities = ["read"]
}
EOF
```

### 3.2 Token Authentication

Authenticate and get a token:

```bash
vault login -method=userpass username=myuser password=mypassword
```

...

---

## 4. Manage Go Application Secrets Using Vault

### 4.1 Vault Integration in Go

Integrate Vault into a Go application:

```go
import (
	"github.com/hashicorp/vault/api"
	"log"
)

func main() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	// Use the client to interact with Vault in your Go application
}
```

### 4.2 Retrieve Secrets

Retrieve secrets from Vault in your Go application:

```go
secret, err := client.Logical().Read("secret/data/myapp/credentials")
if err != nil {
	log.Fatal(err)
}

```

...

---

