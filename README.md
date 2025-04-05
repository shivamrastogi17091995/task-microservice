# Task Management Microservice

## 📌 Overview
This is a **Task Management System** built using **Go** and the **Gin framework**. It follows **microservices architecture principles**, enabling scalability and maintainability. The service allows users to **create, read, update, and delete (CRUD) tasks** while supporting **pagination and filtering**.

---

## 📌 Problem Breakdown & Design Decisions
### 🔹 **Problem Statement**
We need a **task management system** that allows users to **manage tasks efficiently** while ensuring **scalability** and **extensibility**.

### 🔹 **Design Decisions**
- **Microservices Architecture**: Allows independent scaling and updates.
- **Gin Framework**: A fast, lightweight web framework for API development.
- **PostgreSQL**: Chosen for its reliability and support for structured data.
- **Docker & Kubernetes**: Enables containerization and horizontal scaling.
- **gRPC (Future Consideration)**: Used for inter-service communication in a multi-service setup.
- **RESTful API**: Ensures clear and consistent communication.

---

## 📌 🚀 Running the Service
### **1️⃣ Clone the repository**
```sh
git clone https://github.com/shivamrastogi17091995/task-microservice.git
cd task-microservice
```

### **2️⃣ Run the application**
```sh
docker-compose up
```

### **✅ API available at http://localhost:8080/tasks**

### **✅ Docs available at http://localhost:8080/swagger/index.html**

---

## 📌 API Endpoints
- `POST /tasks` - Create a task
- `GET /tasks/{id}` - Get single task
- `PUT /tasks/{id}` - Update task
- `DELETE /tasks/{id}` - Delete task
- `GET /tasks` - List tasks (with pagination & filtering)

---

## 📌 Example Requests
### **1️⃣ Create a task**
```sh
curl -X 'POST' \
  'http://localhost:8080/tasks' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{"Title":"Test","Description":"Test"}'
```

### **2️⃣ Get single task**
```sh
curl -X 'GET' \
  'http://localhost:8080/tasks/1' \
  -H 'accept: application/json'
```

### **3️⃣ Update task**
```sh
curl -X 'PUT' \
  'http://localhost:8080/tasks/1' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{"Title":"Test","Description":"Test","Status":"Completed"}'
```

### **4️⃣ Delete task**
```sh
curl -X 'DELETE' \
  'http://localhost:8080/tasks/1' \
  -H 'accept: application/json'
```

### **5️⃣ List tasks (with pagination & filtering)**
```sh
curl -X 'GET' \
  'http://localhost:8080/tasks?status=Pending&page=1' \
  -H 'accept: application/json'
```

---

## 📌 Microservices Concepts Demonstrated
### 🔹 **Separation of Concerns:**
- Business logic (services/)
- Database operations (repository/)
- API handlers (handlers/)

### 🔹 **Scalability:**
- Containerized using Docker.
- Can be horizontally scaled with Kubernetes.

### 🔹 **Inter-Service Communication:**
- Future extension with gRPC or message queues (Kafka, RabbitMQ).

---

## 📌 Scaling the Service Horizontally
- **Kubernetes Deployment**: Deploy multiple instances of this service behind a load balancer.
- **Database Connection Pooling**: Optimized connections with PostgreSQL.
- **Rate Limiting & Caching**: Using Redis to handle high loads.

---

## 📌 Future Microservices Communication
### 🔹 **If we add a User Service, we can use:**
- **REST API**: Simple but adds overhead.
- **gRPC**: High performance, low latency.
- **Message Queues (Kafka, RabbitMQ)**: Best for asynchronous processing.
