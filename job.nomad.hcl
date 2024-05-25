job "aaallt2" {
  type = "service"

  group "aaallt2" {
    network {
      port "http" { }
    }

    service {
      name     = "aaallt2"
      port     = "http"
      provider = "nomad"
      tags = [
        "traefik-external.enable=true",
        "traefik-external.http.routers.aaallt2.rule=Host(`aaallt.datasektionen.se`)||HostRegexp(`{[^.]+}.datasektionen.se`)",
        "traefik-external.http.routers.aaallt2.priority=1",
        "traefik-external.http.routers.aaallt2.entrypoints=websecure",
        "traefik-external.http.routers.aaallt2.tls.certresolver=default",
      ]
    }

    task "aaallt2" {
      driver = "docker"

      config {
        image = var.image_tag
        ports = ["http"]
      }

      template {
        data        = "PORT={{ env \"NOMAD_PORT_http\" }}"
        destination = "local/.env"
        env         = true
      }

      resources {
        memory = 10
      }
    }
  }
}

variable "image_tag" {
  type = string
  default = "ghcr.io/datasektionen/aaallt2:latest"
}
