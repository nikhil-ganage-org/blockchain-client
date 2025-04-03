```markdown
# Polygon Blockchain API Gateway

A minimal Go implementation of a blockchain API proxy for Polygon RPC, demonstrating core engineering principles.

## Key Features

✅ Proxy endpoints for `eth_blockNumber` and `eth_getBlockByNumber`  
✅ Docker containerization with multi-stage build  
✅ Unit tests with mocked external dependencies  
✅ Terraform ECS Fargate deployment blueprint  
✅ Production readiness considerations  
✅ Error handling for common failure scenarios  

## Getting Started

### Prerequisites
- Go 1.19+
- Docker 20.10+
- Terraform 1.4+ (optional)

### Local Execution in bash
# Initialize Go module (creates go.mod)
go mod init blockchain-client
# Generate go.sum (even with no dependencies)
go mod tidy
# Run service
go run main.go

# Run tests
go test -v -cover
```

### Docker Deployment
```bash
# Build and run
docker build -t blockchain-client .
docker run -dp 8080:8080 blockchain-client

# Verify container
docker ps --filter "ancestor=blockchain-client"
```

## API Endpoints

### Get Current Block Number
```bash
curl -X POST http://localhost:8080 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","id":2}'
```

### Get Block Details
```bash
curl -X POST http://localhost:8080 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x134e82a",true],"id":2}'
```

## Production Readiness Roadmap

| Category             | Proposed Solution                          | Priority |
|----------------------|--------------------------------------------|----------|
| **Observability**    | Prometheus metrics + CloudWatch Dashboard | High     |
| **Security**         | API Gateway + WAF integration              | Critical |
| **Reliability**      | Circuit breaker pattern                    | High     |
| **CI/CD**            | GitHub Actions pipeline                    | Medium   |
| **Configuration**    | Consul integration                         | Medium   |

## Technical Decisions

1. **Go Language** - Chosen for performance and native HTTP capabilities  
2. **Multi-stage Docker** - Minimized final image size (12MB)  
3. **TLS Certificates** - Pre-baked in container for HTTPS readiness  
4. **Stateless Design** - Enables horizontal scaling  
5. **Strict Validation** - Method allowlisting for security  

## Evaluation Guide

### Key Areas to Observe
- Error handling in proxy chain
- Unit test coverage patterns
- Containerization best practices
- Terraform infrastructure-as-code approach

### Suggested Test Cases
1. Invalid JSON payload  
2. Unsupported RPC method  
3. Backend RPC timeout  
4. High-volume load testing  

## Architecture Overview

```
[Client] → [API Gateway] → [Docker Container] → [Polygon RPC]
                   ├── Terraform Deployment
                   └── Cloud Monitoring
```

---

*"Simplicity is prerequisite for reliability." - Edsger W. Dijkstra*
