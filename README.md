# E-commerce Microservices Project

A modern microservices-based e-commerce backend system built with Go, featuring gRPC communication and GraphQL API.

## 🏗️ Architecture Overview

This project implements a **microservices architecture** with three main backend services interfacing through a single gateway:

```text
┌─────────────────┐    gRPC     ┌─────────────────┐
│   API Gateway   │◄────────────┤  Order Service  │
│    (GraphQL)    │             │    (gRPC)       │
│                 │             └─────────────────┘
│                 │             ┌─────────────────┐
│ - GraphQL API   │    gRPC     │Track Order Serv │
│ - Request Route │◄────────────┤    (gRPC)       │
│ - Schema Valid  │             └─────────────────┘
│                 │             ┌─────────────────┐
│                 │    gRPC     │  User Service   │
└────────┬────────┘◄────────────┤    (gRPC)       │
         │                      └─────────────────┘
         │ HTTP/GraphQL
         ▼
┌─────────────────┐
│   Clients       │
│ (Frontend Apps) │
└─────────────────┘
```

### Service Responsibilities

**API Gateway (Port 8080):**
- Provides GraphQL API for external clients
- Acts as a facade for the microservices
- Handles GraphQL schema and resolvers
- Routes requests to appropriate services

**Order Service (Port 50053):**
- Implements business logic for order management
- Provides gRPC endpoints for Order CRUD operations
- Manages order data in memory

**Track Order Service (Port 50054):**
- Implements business logic for tracking order shipments
- Provides gRPC endpoints to get and update shipping status
- Manages tracking data in memory

**User Service (Port 50052):**
- Implements business logic for user management
- Provides gRPC endpoints for User CRUD operations
- Manages user data in memory

## 🛠️ Technologies Used

### Core Technologies
- **Go 1.25.1** - Primary programming language
- **gRPC** - Inter-service communication
- **GraphQL** - API query language
- **Protocol Buffers** - Interface definition and serialization

### Libraries & Frameworks
- **gqlgen** - GraphQL server library for Go
- **Google Protocol Buffers** - Data serialization
- **Go modules** - Dependency management

### Development Tools
- **protoc** - Protocol buffer compiler
- **gqlgen** - GraphQL code generation
- **Go workspace** - Multi-module project management

## 📁 Project Structure

```text
mini-project/
├── go.work                 # Go workspace configuration
├── go.work.sum             # Workspace checksums
├── api-gateway/            # GraphQL API Gateway Service
│   ├── go.mod
│   ├── server.go           # Main server entry point
│   └── graph/
│       ├── schema.graphqls        # GraphQL schema definition
│       ├── schema.resolvers.go    # GraphQL resolvers
│       ├── resolver.go            # Dependency injection
│       ├── generated.go           # Auto-generated GraphQL code
│       └── model/
│           └── models_gen.go      # Generated GraphQL models
├── order-service/          # gRPC Order Service
│   ├── go.mod
│   ├── main.go             # Main service entry point
│   └── proto/
│       ├── order.proto            # Protocol buffer definitions
│       └── ...                    # Generated Go code
├── track-order-service/    # gRPC Track Order Service
│   ├── go.mod
│   ├── main.go             # Main service entry point
│   └── proto/
│       ├── track_order.proto      # Protocol buffer definitions
│       └── ...                    # Generated Go code
└── user-service/           # gRPC User Service
    ├── go.mod
    ├── main.go             # Main service entry point
    └── proto/
        ├── user.proto             # Protocol buffer definitions
        └── ...                    # Generated Go code
```

## 🚀 Getting Started

### Prerequisites
- Go 1.25.1 or later
- Protocol Buffer Compiler

### Installation & Setup

1. **Clone and navigate to the project:**
   ```bash
   cd mini-project
   ```

2. **Initialize Go modules:**
   ```bash
   go work sync
   ```

3. **Build all services:**
   ```bash
   go build ./...
   ```

### Running the Services

You must run all 4 services concurrently. Open multiple terminal windows and run:

1. **Start the Order Service (gRPC):**
   ```bash
   cd order-service
   go run main.go
   ```

2. **Start the Track Order Service (gRPC):**
   ```bash
   cd track-order-service
   go run main.go
   ```

3. **Start the User Service (gRPC):**
   ```bash
   cd user-service
   go run main.go
   ```

4. **Start the API Gateway (GraphQL):**
   ```bash
   cd api-gateway
   go run server.go
   ```
   The GraphQL playground will be available at `http://localhost:8080/`

## 📡 API Documentation

### GraphQL Schema

#### Types
```graphql
type Order {
  id: ID!
  user_id: String!
  item_name: String!
  quantity: Int!
  total_price: Float!
  status: String!
}

type TrackOrder {
  id: ID!
  order_id: String!
  shipping_status: String!
}

type User {
  id: ID!
  username: String!
  email: String!
}
```

#### Queries
```graphql
type Query {
  getOrder(id: ID!): Order!
  listOrders: [Order!]!
  getTrackingInfo(order_id: String!): TrackOrder!
  getUser(id: ID!): User!
  listUsers: [User!]!
}
```

#### Mutations
```graphql
type Mutation {
  createOrder(user_id: String!, item_name: String!, quantity: Int!, total_price: Float!): Order!
  updateTrackingStatus(order_id: String!, shipping_status: String!): TrackOrder!
  createUser(username: String!, email: String!): User!
  updateUser(id: ID!, username: String!, email: String!): User!
  deleteUser(id: ID!): Boolean!
}
```

### Example GraphQL Operations

#### Workflow Example: End-to-End
The API Gateway enforces business logic across microservices. The strict workflow is:
1. **Create a User** (Get a `user_id`)
2. **Create an Order** (Requires a valid `user_id`)
3. **Track an Order** (Requires a valid `order_id`)

#### 1. Create a User
```graphql
mutation {
  createUser(username: "sagar", email: "sagar@example.com") {
    id
    username
  }
}
```

#### 2. Create an Order (Using the User ID)
```graphql
mutation {
  createOrder(
    user_id: "user-123"
    item_name: "Wireless Mouse"
    quantity: 2
    total_price: 39.98
  ) {
    id
    item_name
    status
  }
}
```

#### 3. Update Tracking Status (Using the Order ID)
```graphql
mutation {
  updateTrackingStatus(
    order_id: "order-1"
    shipping_status: "SHIPPED"
  ) {
    id
    order_id
    shipping_status
  }
}
```

#### List All Orders
```graphql
query {
  listOrders {
    id
    item_name
    quantity
    total_price
    status
  }
}
```

## 🔧 Development

### Code Generation

**Generate GraphQL code:**
```bash
cd api-gateway
go run github.com/99designs/gqlgen generate
```

**Generate Protocol Buffer code:**
```bash
cd order-service
protoc --go_out=. --go-grpc_out=. proto/order.proto

cd track-order-service
protoc --go_out=. --go-grpc_out=. proto/track_order.proto
```

## 🏛️ Design Decisions

### Why Microservices?
- **Scalability**: Services can be scaled independently
- **Technology Diversity**: Different services can use different tech stacks
- **Team Autonomy**: Teams can work on different services independently
- **Fault Isolation**: Failure in one service doesn't affect others

### Why gRPC for Inter-Service Communication?
- **Performance**: Binary protocol, faster than REST/JSON
- **Type Safety**: Strongly typed interfaces
- **Streaming**: Support for bidirectional streaming
- **Code Generation**: Automatic client/server code generation

### Why GraphQL for External API?
- **Flexible Queries**: Clients request exactly what they need
- **Single Endpoint**: One URL for all operations
- **Strongly Typed**: Schema provides type safety

### Why In-Memory Storage?
- **Simplicity**: No database setup required for demo
- **Speed**: Fast read/write operations

## 🚀 Future Enhancements
- **Database Integration**: Replace in-memory storage with PostgreSQL/MongoDB
- **Authentication**: Add JWT-based authentication
- **Logging/Monitoring**: Add structured logging and metrics
- **Docker**: Containerize services

---

**Built with ❤️ using Go, gRPC, and GraphQL**