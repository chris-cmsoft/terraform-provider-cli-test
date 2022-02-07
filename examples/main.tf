terraform {
  required_providers {
    cli = {
      version = "0.2"
      source  = "hashicorp.com/edu/cli"
    }
  }
}

provider "cli" {}

data "cli_onepassword_version" "local" {}

output "onepassword-version" {
  value = data.cli_onepassword_version.local.version
}

data "cli_onepassword_users" "cs" {}

output "onepassword-users" {
  value = data.cli_onepassword_users.cs.users
}
