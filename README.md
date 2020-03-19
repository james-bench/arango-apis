# ArangoDB Cloud API's

[Public API definitions](https://arangodb-managed.github.io/apis/) of ArangoDB Cloud

## Build

The build requires a github personal access token with the following scopes.

* repo
    * repo:status
    * repo_deployment
    * public_repo
    * repo:invite

The token must be placed in a file called `${HOME}/.arangodb/ms/github-readonly-code-acces.token`.

To compile the protobuffer specifications and build
Go wrappers for them, run:

```bash
make build-image
make
```
