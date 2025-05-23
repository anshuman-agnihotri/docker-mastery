````markdown
# 🕵️‍♂️ Fing Agent via Docker

Easily run Fing Agent in a Docker container for real-time network monitoring, device discovery, and diagnostics.

> 📌 Supports: ARM64 and AMD64  
> ⚠️ Linux-only | macOS & Windows not supported via Docker  
> 🐳 Docker & Docker Compose compatible  

---

## 📋 Table of Contents

- # Prerequisites
- # Limitations
- # Install via Docker Compose
- # Manual Installation via Docker
- # Migrating from Snap
- # Uninstallation
- # Resources

---

## ✅ Prerequisites

Make sure you have the following ready:

- Docker installed and running.
- Docker Compose installed (`docker compose` or `docker-compose`).
- ARM64 or AMD64 architecture.
- An active internet connection.

---

## ⚠️ Limitations

- ❌ **Not supported on Windows or macOS** via Docker.
- 🚫 Only **one Fing Agent per device**.
- 🔄 **Migration from Snap to Docker is not supported**. You'll need to uninstall the Snap version first.

---

## 📦 Install via Docker Compose (Recommended)

This is the easiest and most automated way to run the Fing Agent.

### 🧰 Step-by-Step:

```bash
# Download the Docker Compose YAML
curl https://get.fing.com/fing-docker/compose.yaml -o compose.yaml

# Start the container in detached mode
docker compose up -d    # For Docker v2+
# OR
docker-compose up -d    # For older Docker versions
````

✅ This will:

* Download the latest Fing Agent image
* Run it as a background service
* Enable persistent volumes
* Use host network mode for full visibility

---

## 🛠 Manual Installation via Docker

Prefer the CLI? Here’s how:

```bash
sudo docker run --network=host --cap-add NET_ADMIN \
  --publish 44444:44444 -d --restart=always \
  --name FingAgent \
  -v fing-data-volume:/app/fingdata \
  fing/fing-agent:latest
```

📌 This setup:

* Uses host networking (needed for network scanning)
* Adds `NET_ADMIN` capabilities
* Publishes port `44444` (for UPnP status)
* Mounts volume for persistent data

---

## 🔄 Migrating from Snap (Optional)

If you have an existing Snap install:

1. **Deactivate** the Snap version in the Fing mobile app (long press > Deactivate)
2. Run the following to uninstall:

```bash
sudo snap remove fing-agent
```

3. Proceed with Docker installation as described above.

---

## ❌ Uninstallation

To stop and remove the Docker container:

```bash
docker stop FingAgent
docker rm FingAgent
```

To remove the volume:

```bash
docker volume rm fing-data-volume
```

---

## 📚 Resources

* [Fing Agent Docs](https://help.fing.com/hc/en-us/articles/16874649602588-How-to-install-the-Fing-Agent-on-Linux-Docker)
* [Fing Official Docker Hub](https://hub.docker.com/r/fing/fing-agent)
* [Docker Installation Guide](https://docs.docker.com/get-docker/)

---

## 📬 Feedback or Questions?

Feel free to open an [Issue](https://github.com/anshuman-agnihotri/docker-mastery) or contribute with a Pull Request.

---

> 🧠 Happy Monitoring!
> Built with 💙 by [Anshuman](https://www.linkedin.com/in/anshuman-agnihotri)

```
