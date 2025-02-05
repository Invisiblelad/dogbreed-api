# Dog Breed API - Kubernetes Deployment

This project consists of a **Go-based Dog Breed API** running on **Kubernetes** with a MongoDB backend. This README provides the steps to deploy the application, access the service, and interact with the API using **curl**.

## Project Setup

### Prerequisites

- Kubernetes (MicroK8s, Minikube, or any other cluster)
- Docker (for building and pushing images)
- Go (for building the Go application)
- `kubectl` (for interacting with Kubernetes)
- `curl` (for making HTTP requests)

### Kubernetes Resources

- **Go API** running as a service (`dogbreed-service`) in the `sharath` namespace.
- **MongoDB** running as a service (`mongo-service`).

## Deployment Steps

### 1. Deploy MongoDB
Make sure MongoDB is running in your Kubernetes cluster as a service (`mongo-service`).

```
kubectl apply -f manifests/deployment-mongo.yaml 
```

### 2. Deploy Application

```
kubectl apply -f manifests/deployment.yaml -f manifests/service.yaml

```

### 3. Port Forward for external access

```
kubectl port-forward svc/dogbreed-service 8080:8080

```

## API ENDPOINS

**GET /dogbreeds**
*Get a list of all dog breeds.*
```
curl -X GET http://dogbreed-service:8080/dogbreeds

```

**POST /dogbreeds**
*Create a new dog breed.*

```
curl -X POST http://localhost:8080/dogbreeds \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Golden Retriever",
        "description": "Large",
        "origin": "United States"
      }'
```

**PUT /dogbreeds/{id}**
*Update an existing dog breed by its ID.*

```
curl -X PUT http://localhost:8080/dogbreeds/1 \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Golden Retriever",
        "description": "Large",
        "origin": "Canada"
      }'

```
**DELETE /dogbreeds/{id}**
*Delete a dog breed by its ID.*

```
curl -X DELETE http://localhost:8080/dogbreeds/1

```

