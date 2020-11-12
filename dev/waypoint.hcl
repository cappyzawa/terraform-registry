project = "terraform-registry"

app "api" {
  path = "../"

  labels = {
    "service" = "terraform-registry",
    "env" = "dev"
  }

  build {
    use "docker" {
      dockerfile = "Dockerfile"
    }
  }

  deploy {
    use "docker" {
      static_environment = {
        CONFIG_FILE="/testdata/config.yaml"
      }
    }
  }
}
