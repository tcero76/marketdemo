terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

# Workers
resource "digitalocean_droplet" "swarm_worker" {
  count     = var.worker_count
  name      = "swarm-worker-${count.index}"
  vpc_uuid  = var.vpc
  region    = var.region
  size      = var.size
  image     = var.image
  ssh_keys  = var.ssh_keys

  depends_on = [var.depends]
  user_data = templatefile("${path.module}/scripts/install_docker_worker.sh.tmpl", {
    manager_ip = var.manager_ip
    is_manager = false
    internal_privkey = var.internal_privkey
    internal_pubkey = ""
    network = ""
  })
}

resource "digitalocean_volume" "postgres_data" {
  name   = "swarm-postgres-data"
  region = var.region
  size   = 10
}

resource "digitalocean_volume_attachment" "postgres_data_attachment" {
  droplet_id = digitalocean_droplet.swarm_worker[0].id
  volume_id  = digitalocean_volume.postgres_data.id
}