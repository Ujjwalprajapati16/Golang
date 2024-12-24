# Students API

Students API is a RESTful API built with Golang to manage student records. It provides endpoints to create, read, update, and delete student information.

## Features

- Create a new student
- Get a student by ID
- Get a list of students
- Update a student
- Delete a student

## Project Structure

```
25_Project/
├── cmd/
│   └── students_API/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go
│   ├── storage/
│   │   ├── sqlite/
│   │   │   └── sqlite.go
│   │   └── storage.go
│   ├── types/
│   │   └── types.go
│   └── utils/
│       └── response/
│           └── response.go
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.23.4 or later
- SQLite

### Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/students-api.git
cd students-api/25_Project
```

2. Install dependencies:

```sh
go mod tidy
```

3. Create a configuration file:

Create a `config.yaml` file with the following content:

```yaml
env: dev
storage_path: "./students.db"
http_server:
  address: "localhost:8082"
```

### Running the Application

1. Start the server:

```sh
go run cmd/students_API/main.go --config=config.yaml
```

2. The server will start on `localhost:8082`.

### API Endpoints

- **POST /api/students**: Create a new student
- **GET /api/students/{id}**: Get a student by ID
- **GET /api/students**: Get a list of students
- **PUT /api/students/{id}**: Update a student
- **DELETE /api/students/{id}**: Delete a student

### Example Requests

- **Create a new student**

```sh
curl -X POST http://localhost:8082/api/students -d '{"name": "John Doe", "email": "john@example.com", "age": 20}' -H "Content-Type: application/json"
```

- **Get a student by ID**

```sh
curl http://localhost:8082/api/students/1
```

- **Get a list of students**

```sh
curl http://localhost:8082/api/students
```

- **Update a student**

```sh
curl -X PUT http://localhost:8082/api/students/1 -d '{"name": "John Smith", "email": "johnsmith@example.com", "age": 21}' -H "Content-Type: application/json"
```

- **Delete a student**

```sh
curl -X DELETE http://localhost:8082/api/students/1
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.