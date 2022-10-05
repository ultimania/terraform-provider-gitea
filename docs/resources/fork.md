---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitea_fork Resource - terraform-provider-gitea"
subcategory: ""
description: |-
  gitea_fork manages repository fork to the current user or an organisation
  Forking a repository to a dedicated user is currently unsupported
  Creating a fork using this resource without an organisation will create the fork in the executors name
---

# gitea_fork (Resource)

`gitea_fork` manages repository fork to the current user or an organisation
Forking a repository to a dedicated user is currently unsupported
Creating a fork using this resource without an organisation will create the fork in the executors name

## Example Usage

```terraform
resource "gitea_org" "org1" {
  name = "org1"
}

resource "gitea_org" "org2" {
  name = "org2"
}

resource "gitea_repository" "repo1_in_org1" {
  username = gitea_org.org1.name
  name     = "repo1-in-org1"
}

resource "gitea_fork" "user_fork_of_repo1_in_org1" {
  owner = gitea_org.org1.name
  repo  = gitea_repository.repo1_in_org1.name
}

resource "gitea_fork" "org2_fork_of_repo1_in_org1" {
  owner        = gitea_org.org1.name
  repo         = gitea_repository.repo1_in_org1.name
  organization = gitea_org.org2.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `owner` (String) The owner or owning organization of the repository to fork
- `repo` (String) The name of the repository to fork

### Optional

- `organization` (String) The organization that owns the forked repo

### Read-Only

- `id` (String) The ID of this resource.

