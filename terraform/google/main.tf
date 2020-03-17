terraform {
  backend "local" {
    path = "/Users/gabrielradureau/terraform-states/my-heroku-app/terraform.tfstate"
  }
}

provider "google" {
  credentials = "/Users/gabrielradureau/my-heroku-app/gcs-backend-herge-lab-credentials.json"
  project     = "herge-lab"
  # 5 GB-months of Regional Storage per month (US regions only — excluding Northern Virginia [us-east4])
  region      = "us-east1"
  zone        = "us-east1-b"
}

resource "google_storage_bucket" "tf_backend" {
  name = "herge-lab-terraform-states"
  location = "US-EAST1"
  storage_class = "REGIONAL"
}

output "bucket_name" {
  value = google_storage_bucket.tf_backend.name
  description = "bucket name used as a gcs terraform backend"
}