# GoCache
Golang Caching Microservices

## Overview
GoCache is a caching microservice built in Go (Golang) that provides a simple and efficient way to store and retrieve data in memory. It is designed to be lightweight and easy to integrate into other applications.

## Features
- In-memory caching for fast data retrieval
- Simple API for setting, getting, and deleting cache entries
- Configurable settings for easy deployment
- Extensible architecture for future enhancements

## Project Structure
```
gocache
├── cmd
│   └── gocache
│       └── main.go         # Entry point of the application
├── internal
│   ├── cache
│   │   ├── cache.go        # Cache interface and implementation
│   │   └── memory.go       # In-memory cache implementation
│   ├── config
│   │   └── config.go       # Configuration handling
│   ├── server
│       └── server.go       # HTTP server setup and routes
├── pkg
│   └── utils
│       └── utils.go        # Utility functions
├── go.mod                   # Module definition and dependencies
└── README.md                # Project documentation
```

## Installation
To install GoCache, clone the repository and run the following commands:

```bash
go mod tidy
```

## Usage
To start the caching microservice, navigate to the `cmd/gocache` directory and run:

```bash
go run main.go
```

The server will start and listen for incoming requests.

## API Endpoints
- `POST /cache` - Set a cache entry
- `GET /cache/{key}` - Get a cache entry
- `DELETE /cache/{key}` - Delete a cache entry

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.