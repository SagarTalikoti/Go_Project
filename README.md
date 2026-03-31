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
1. **Create a User** → get a `user_id`
2. **Create an Order** → provide a valid `user_id`, get back an `order_id`
3. **Update Tracking Status** → provide a valid `order_id` to sync status across both services
4. **List All Orders** → see the order status updated to match the tracking status

---

### 👤 User CRUD

#### Create a User
```graphql
mutation {
  createUser(username: "sagar", email: "sagar@example.com") {
    id
    username
    email
  }
}
```

#### Get a Single User
```graphql
query {
  getUser(id: "user-1") {
    id
    username
    email
  }
}
```

#### List All Users
```graphql
query {
  listUsers {
    id
    username
    email
  }
}
```

#### Update a User
```graphql
mutation {
  updateUser(id: "user-1", username: "sagar_updated", email: "sagar2@example.com") {
    id
    username
    email
  }
}
```

#### Delete a User
```graphql
mutation {
  deleteUser(id: "user-1")
}
```

---

### 📦 Order Operations

#### Create an Order *(valid user_id required)*
```graphql
mutation {
  createOrder(
    user_id: "user-1"
    item_name: "Wireless Mouse"
    quantity: 2
    total_price: 39.98
  ) {
    id
    user_id
    item_name
    quantity
    total_price
    status
  }
}
```

#### Get a Single Order
```graphql
query {
  getOrder(id: "order-1") {
    id
    user_id
    item_name
    quantity
    total_price
    status
  }
}
```

#### List All Orders
> **Result:** After updating tracking status, `status` here will automatically reflect the latest tracking value (e.g. `"Delivered"`).
```graphql
query {
  listOrders {
    id
    user_id
    item_name
    quantity
    total_price
    status
  }
}
```

---

### 🚚 Track Order Operations

#### Update Tracking Status *(valid order_id required)*
> **Note:** This simultaneously updates the Order's `status` field in `order-service`.
```graphql
mutation {
  updateTrackingStatus(
    order_id: "order-1"
    shipping_status: "Delivered"
  ) {
    id
    order_id
    shipping_status
  }
}
```

#### Get Tracking Info for an Order
```graphql
query {
  getTrackingInfo(order_id: "order-1") {
    id
    order_id
    shipping_status
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