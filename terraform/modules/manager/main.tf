terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

# Manager
resource "digitalocean_droplet" "swarm_manager" {
  name      = "swarm-manager"
  vpc_uuid  = var.vpc
  region    = var.region
  size      = var.size
  image     = var.image
  ssh_keys  = var.ssh_keys

  user_data = templatefile("${path.module}/scripts/install_docker_manager.sh.tmpl", {
    manager_ip = ""
    is_manager = true
    internal_pubkey = var.internal_pubkey
    internal_privkey = ""
    network = var.overlay_network
  })
}

resource "digitalocean_volume" "postgres_data" {
  name   = "swarm-postgres-data"
  region = var.region
  size   = 10
}

resource "digitalocean_volume_attachment" "postgres_data_attachment" {
  droplet_id = digitalocean_droplet.swarm_manager.id
  volume_id  = digitalocean_volume.postgres_data.id
}