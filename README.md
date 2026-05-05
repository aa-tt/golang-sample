# Java Spring Boot vs. Golang for Microservices

This document outlines the differences between Java Spring Boot and Golang (Go) for microservice architectures, specifically in the context of a supply chain technology organization like **Makro**.

## Comparison Table

| Feature | Java Spring Boot | Golang |
| :--- | :--- | :--- |
| **Startup Time** | Slow (JVM overhead, classpath scanning) | Instant (Compiled binary) |
| **Memory Footprint** | High (JVM heap, stack, Metaspace) | Very Low (No VM runtime overhead) |
| **Concurrency** | Threads (Heavyweight, now Loom virtual threads) | Goroutines (Lightweight, millions per process) |
| **Binary Size** | Large (Fat JARs + JRE) | Small (Static executable) |
| **Developer Productivity** | High (Extensive ecosystem, "Magic") | High (Simple syntax, fast builds) |
| **Scalability** | Vertical focus (large heaps) | Horizontal focus (efficient packing) |

## Why Golang for Makro?

Makro, as a supply chain tech organization, deals with high-frequency events, real-time routing, and massive ingestion pipelines. Golang fits this ecosystem perfectly due to:

1.  **Lower Latency**: Essential for real-time routing and warehouse automation signals.
2.  **Efficient Resource Utilization**: Running hundreds of microservices on Kubernetes is significantly cheaper with Go's low memory consumption.
3.  **Fast Iteration**: Near-instant compile times allow developers to test and deploy faster.
4.  **Static Binaries**: Simplifies containerization and deployment, making supply chain logic portable.

---

## Supply Chain Scenarios in Golang (Makro Examples)

### 1. Ingestion Service (Real-time IoT/Telematics)
Go's `net/http` and `encoding/json` are highly optimized for high-throughput ingestion.
```go
func IngestTelematics(w http.ResponseWriter, r *http.Request) {
    var data TelematicsData
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Async process to downstream Kafka/PubSub
    go processIngestion(data)
    w.WriteHeader(http.StatusAccepted)
}
```

### 2. Event-Driven Services (Kafka Consumers)
Processing warehouse events using channels.
```go
func ConsumeWarehouseEvents(reader *kafka.Reader) {
    for {
        m, err := reader.ReadMessage(context.Background())
        if err != nil {
            break
        }
        // Concurrent processing with worker pool
        jobChannel <- m.Value
    }
}
```

### 3. Routing Engine (Path Optimization)
Goroutines can parallelize complex route computations across multiple warehouses.
```go
func OptimizeRoutes(destinations []string) {
    var wg sync.WaitGroup
    for _, dest := range destinations {
        wg.Add(1)
        go func(d string) {
            defer wg.Done()
            calculateBestRoute(d)
        }(dest)
    }
    wg.Wait()
}
```

### 4. Updating Warehouses (Database Interactions)
Using `sqlx` or `gorm` for efficient, typed updates.
```go
func UpdateInventory(db *sql.DB, itemID string, qty int) error {
    _, err := db.Exec("UPDATE inventory SET quantity = quantity - $1 WHERE id = $2", qty, itemID)
    return err
}
```

### 5. Scheduling (Cron/Timed Tasks)
Lightweight tickers for inventory reconciliation.
```go
ticker := time.NewTicker(1 * time.Hour)
for {
    select {
    case <-ticker.C:
        reconcileStock()
    }
}
```

---

## Production Strategy (AWS Focus)

AWS is recommended for Makro due to its mature serverless and container ecosystem.

### 1. Infrastructure as Code (IaC)
- Use **Terraform** or **AWS CDK** to define VPCs, EKS clusters, and RDS instances.

### 2. Deployment: Amazon EKS
- **Reason**: Go binaries are extremely small (10-20MB). We can pack more Go containers per node compared to Java.
- **Auto-scaling**: Use **Karpenter** for rapid node scaling during peak supply chain hours.

### 3. Observability
- **AWS X-Ray**: For tracing requests across the routing and ingestion pipeline.
- **CloudWatch Metrics**: Custom metrics for "Order to Warehouse Acknowledgement" latency.

### 4. CI/CD Pipeline
- **GitHub Actions**: Build the Go binary once, package in a `scratch` or `distroless` image, and push to **Amazon ECR**.
- **ArgoCD**: Use GitOps for automated deployments to the cluster.

### 5. Security
- **IAM Roles for Service Accounts (IRSA)**: Ensure the Ingestion service only has rights to write to specific S3 buckets or Kinesis streams.
- **Static Analysis**: Integrate `golangci-lint` and `gosec` in the pipeline.

----

### Features Implemented:
* **REST API**: Built using the Gin framework.
* **Consumer/Producer**: Uses Go channels to simulate an asynchronous messaging pattern where a producer sends data to a background consumer.
* **Persistence**: SQLite in-memory database for fast, relational storage.
* **API Gateway/Routing**: Centralized routing with grouped handlers.
* **Rate Limiter**: Token bucket implementation via golang.org/x/time/rate.
* **Authentication**: JWT interceptor for protecting API routes.
* **Exception Handling**: Global recovery middleware and custom error handlers.
* **Logging**: Custom middleware to log request methods, paths, and durations.
### Project Structure:
* `main.go`: Application entry point and router configuration.
* `db.go`: Database initialization and schema.
* `middleware.go`: JWT, Rate Limiter, and Logging logic.
* `handlers.go`: API request handlers and error management.
* `product_service.go`: Business logic, producer, and consumer implementations.

**How to Run:**
1. Ensure Go is installed (the setup attempted to install it via Homebrew).
2. Run the application:
```go run main.go```
3. Login to get a token:
`POST /login` with body `{"username": "admin", "password": "any"}`
4. Create a product (Protected):
`POST /api/products` with `Authorization: Bearer {token}`
5. List products:
`GET /api/products` with `Authorization: Bearer {token}`
The service is configured to run on `http://localhost:8080`.