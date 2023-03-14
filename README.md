go-clean-architecture
├── config
│   └── config.go
├── docs
│   ├── app.env.example
│   └── swagger.yaml
├── internal
│   └── student
│       ├── handler
│       │   ├── handler.go
│       │   └── payload.go
│       ├── repository
│       │   ├── model.go
│       │   ├── repository.go
│       │   └── repository_mock.go
│       ├── service
│       │   ├── service.go
│       │   └── service_test.go
│       └── entity.go
├── pkg
│   ├── database
│   │   └── mysql.go
│   ├── third_party_api
│   │   └── third_party_api.go
│   └── util
│       └── copy.go
├── README.md
├── app.env
├── go.mod
├── go.sum
└── main.go