# Docker Bake Setup

This project demonstrates how to use **Docker Buildx Bake** to build and run multiple Docker images using a single configuration file (`docker-bake.hcl`).

## 📁 Project Structure

```

docker-bake/
├── api/
│   ├── dockerfile
│   └── index.js
├── web/
│   ├── dockerfile
│   └── index.html
└── docker-bake.hcl

````

## 🚀 Build All Images Using Docker Buildx Bake

Run the following command to build both `api` and `web` images:

```bash
docker buildx bake
````

> This will use the `docker-bake.hcl` file to build both services in parallel.

## 🐳 Run Containers

### 1. Run the API container

```bash
docker run -d --name api-container -p 3000:3000 anshuman/api:latest
```

> Access the API at: [http://localhost:3000](http://localhost:3000)

### 2. Run the Web container

```bash
docker run -d --name web-container -p 8080:80 anshuman/web:latest
```

> Access the web page at: [http://localhost:8080](http://localhost:8080)

## 🧹 Stop & Clean Up

To stop and remove the containers:

```bash
docker stop api-container web-container
docker rm api-container web-container
```

---

## 🧱 Requirements

* Docker version with Buildx support
* Linux/macOS/WSL environment
* Docker daemon running

---

## 📝 Notes

* Make sure the files `index.js` and `index.html` are present in their respective folders (`api/` and `web/`) before building.
* Port numbers (`3000` and `8080`) can be adjusted as needed.

```