## Getting Started

### 1. Setting up the Environment

First, start the Docker containers for MySQL and RabbitMQ:

```bash
docker compose up -d
```

This will start:
- MySQL database on port 3306
- RabbitMQ on ports 5672 (AMQP) and 15672 (Management UI)

### 2. Running Database Migrations

Run the following command to apply database migrations:

```bash
migrate -path=sql/migrations --database="mysql://root:root@tcp(localhost:3306)/orders" up
```

### 3. Starting the Application

Navigate to the application directory and run:

```bash
cd cmd/ordersystem
go run .
```

This will start:
- REST API on port 8000
- GraphQL API on port 8080
- gRPC service on port 50051

## API Endpoints

### REST API

- Create Order: `POST http://localhost:8000/order`
- List Orders: `GET http://localhost:8000/order`

### GraphQL API

- Endpoint: `http://localhost:8080/query`
- Playground: `http://localhost:8080/`
- Operations:
  - Create Order: Mutation
  - List Orders: Query

### gRPC API

- Service: `OrderService`
- Endpoint: `localhost:50051`
- Methods:
  - CreateOrder
  - ListOrders