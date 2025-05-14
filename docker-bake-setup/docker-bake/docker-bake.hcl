target "api" {
  context = "."
  dockerfile = "./api/dockerfile"
  tags = ["anshuman/api:latest"]
}

target "web" {
  context = "."
  dockerfile = "./web/dockerfile"
  tags = ["anshuman/web:latest"]
}

group "default" {
  targets = ["api", "web"]
}
