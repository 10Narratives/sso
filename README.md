# SSO service

## 🚀 Installation

Follow these steps to set up sso:

### 1️⃣ Clone the Repository

``` shell
git clone https://github.com/10Narratives/sso.git
cd task-tracker
```

### 2️⃣ Install Dependencies

Run the following command to install required Go packages:

```shell
go get ./...
```

### 3️⃣ Set Up Environment Variables

Create a .env file and specify the configuration path:

```.env
CONFIG_PATH=./path/to/your/config.yaml
```

### 4️⃣ Create a Custom Configuration File

The project uses a YAML configuration file to manage environment and service settings. Below is a breakdown of available configuration parameters:

| Parameter       | Description                                      | Default/Required         |
|-----------------|--------------------------------------------------|--------------------------|
| `env`           | Specifies the environment (e.g., "local", "dev"). | Default: `"local"`      |
| `token_ttl`     | Token time-to-live duration.                     | Required                 |
| `storage.driver`| Database driver (e.g., "postgres", "mysql").     | Required                 |
| `storage.dsn`   | Data source name (connection string).            | Required                 |
| `grpc.port`     | Port for the gRPC server.                        | Default: `4000`          |
| `grpc.timeout`  | Timeout for gRPC operations.                     | Default: `10s`           |

This is an example of configuration file

```yaml
env: "local"
token_ttl: 4h
storage:
  driver: "postgres"
  dsn: "./storage/sso.db"
grpc:
  port: 4444
  timeout: 4s
```
