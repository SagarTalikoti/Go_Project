# Task Management Microservices Project

A modern microservices-based task management system built with Go, featuring gRPC communication and GraphQL API.

## 🏗️ Architecture Overview

This project implements a **microservices architecture** with two main services:

```
┌─────────────────┐    gRPC     ┌─────────────────┐
│   API Gateway   │◄────────────┤  Task Service   │
│    (GraphQL)    │             │    (gRPC)       │
│                 │             └─────────────────┘
│ - GraphQL API   │
│ - Request Routing│    gRPC     ┌─────────────────┐
│ - Schema Validation│◄────────────┤  User Service   │
└─────────────────┘             │    (gRPC)       │
                                └─────────────────┘
         │
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

**Task Service (Port 50051):**
- Implements business logic for task management
- Provides gRPC endpoints for CRUD operations
- Manages task data in memory

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

```
mini-project/
├── go.work                 # Go workspace configuration
├── go.work.sum            # Workspace checksums
├── api-gateway/           # GraphQL API Gateway Service
│   ├── go.mod
│   ├── go.sum
│   ├── server.go          # Main server entry point
│   └── graph/
│       ├── schema.graphqls        # GraphQL schema definition
│       ├── schema.resolvers.go    # GraphQL resolvers
│       ├── resolver.go           # Dependency injection
│       ├── generated.go          # Auto-generated GraphQL code
│       └── model/
│           └── models_gen.go     # Generated GraphQL models
└── task-service/         # gRPC Task Management Service
    ├── go.mod
    ├── go.sum
    ├── main.go           # Main service entry point
    └── proto/
        ├── task.proto            # Protocol buffer definitions
        ├── task.pb.go           # Generated protobuf code
        ├── task_grpc.pb.go      # Generated gRPC code
        └── protoc-bin/          # Protocol buffer compiler
            ├── bin/
            ├── include/
            └── readme.txt
└── user-service/         # gRPC User Management Service
    ├── go.mod
    ├── go.sum
    ├── main.go           # Main service entry point
    └── proto/
        ├── user.proto            # Protocol buffer definitions
        ├── user.pb.go           # Generated protobuf code
        └── user_grpc.pb.go      # Generated gRPC code
```

## 🚀 Getting Started

### Prerequisites
- Go 1.25.1 or later
- Protocol Buffer Compiler (included in project)

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

1. **Start the Task Service (gRPC):**
   ```bash
   cd task-service
   go run main.go
   ```
   The service will start on `localhost:50051`

2. **Start the User Service (gRPC):**
   ```bash
   cd user-service
   go run main.go
   ```
   The service will start on `localhost:50052`

3. **Start the API Gateway (GraphQL):**
   ```bash
   cd api-gateway
   go run server.go
   ```
   The GraphQL playground will be available at `http://localhost:8080/`

## 📡 API Documentation

### GraphQL Schema

#### Types
```graphql
type Task {
  id: ID!           # Unique identifier
  title: String!    # Task title
  description: String!  # Task description
  completed: Boolean!   # Completion status
}

type User {
  id: ID!
  username: String!
  email: String!
}
```

#### Queries
```graphql
# Get a single task by ID
getTask(id: ID!): Task!

# Get all tasks
listTasks: [Task!]!

# Get a single user by ID
getUser(id: ID!): User!

# Get all users
listUsers: [User!]!
```

#### Mutations
```graphql
# Create a new task
createTask(title: String!, description: String!): Task!

# Update an existing task
updateTask(id: ID!, title: String!, description: String!, completed: Boolean!): Task!

# Delete a task
deleteTask(id: ID!): Boolean!

# User operations
createUser(username: String!, email: String!): User!
updateUser(id: ID!, username: String!, email: String!): User!
deleteUser(id: ID!): Boolean!
```

### Example GraphQL Operations

#### Create a Task
```graphql
mutation {
  createTask(
    title: "Learn Go Microservices"
    description: "Study gRPC and GraphQL integration"
  ) {
    id
    title
    description
    completed
  }
}
```

#### List All Tasks
```graphql
query {
  listTasks {
    id
    title
    description
    completed
  }
}
```

#### Update a Task
```graphql
mutation {
  updateTask(
    id: "task-id-here"
    title: "Updated Title"
    description: "Updated description"
    completed: true
  ) {
    id
    title
    completed
  }
}
```

#### Delete a Task
```graphql
mutation {
  deleteTask(id: "task-id-here")
}
```

#### Create a User
```graphql
mutation {
  createUser(
    username: "johndoe"
    email: "john@example.com"
  ) {
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
  updateUser(
    id: "user-id-here"
    username: "johndoe_updated"
    email: "john_new@example.com"
  ) {
    id
    username
    email
  }
}
```

#### Delete a User
```graphql
mutation {
  deleteUser(id: "user-id-here")
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
cd task-service
protoc --go_out=. --go-grpc_out=. proto/task.proto
```

### Testing the APIs

**Using GraphQL Playground:**
1. Open `http://localhost:8080/` in your browser
2. Use the interactive playground to test queries and mutations

**Using curl:**
```bash
# List tasks
curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{"query": "query { listTasks { id title description completed } }"}'
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
- **Introspection**: API can be explored and documented automatically

### Why In-Memory Storage?
- **Simplicity**: No database setup required for demo
- **Speed**: Fast read/write operations
- **Ephemeral**: Data resets on service restart (acceptable for demo)

## 🔄 Data Flow

1. **Client Request**: GraphQL query/mutation sent to API Gateway
2. **Request Processing**: API Gateway validates and routes request
3. **gRPC Call**: API Gateway calls Task Service via gRPC
4. **Business Logic**: Task Service processes the request
5. **Data Storage**: Task data stored in memory
6. **Response**: Result flows back through gRPC → GraphQL → Client

## 🚀 Future Enhancements

- **Database Integration**: Replace in-memory storage with PostgreSQL/MongoDB
- **Authentication**: Add JWT-based authentication
- **Logging**: Implement structured logging
- **Monitoring**: Add metrics and health checks
- **Docker**: Containerize services
- **Testing**: Add unit and integration tests
- **Caching**: Implement Redis for performance
- **Load Balancing**: Add service discovery and load balancing

## 📝 Notes

- This is a demonstration project showcasing microservices architecture
- Data is stored in memory and will be lost on service restart
- Services communicate via gRPC for type-safe, high-performance inter-service communication
- GraphQL provides a flexible, client-driven API
- The project demonstrates modern Go development practices with modules and workspaces

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

---

**Built with ❤️ using Go, gRPC, and GraphQL**