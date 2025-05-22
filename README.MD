## Getting Started

### Running the Application

Simply run the following command to start the application:

```bash
docker compose up -d
```



## API Endpoints

### REST API

- Create Order: `POST http://localhost:8000/order`
- List Orders: `GET http://localhost:8000/order`

### GraphQL API

- Endpoint: `http://localhost:8080/graphql`
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