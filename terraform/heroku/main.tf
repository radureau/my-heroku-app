terraform {
  backend "gcs" {
    credentials = "/Users/gabrielradureau/my-heroku-app/gcs-backend-herge-lab-credentials.json"
    bucket  = "herge-lab-terraform-states"
    prefix  = "my-heroku-app"
  }
}

provider "heroku" {
  version = "~> 2.2"
}

