# Compiling

## Requirements
- [glide](http://glide.sh/)

## Compilation
Ensure that all requirements are present, and checkout the repo from github.

```bash
# Install dependencies
glide install

cd build
bash ./build-linux.sh

# Optional: Ensure that you target the right docker registry (assuming minikube)
eval $(minikube docker-env)

# Build the docker image
bash ./docker.sh
```

### Workflow Parser
Although this will be integrated in the fission CLI, currently in order to write workflows, you will need the `wfparser` to manually convert your yaml workflow definitions to json.
To build the `wfparser` tool:
```bash
go install github.com/fission/fission-workflows/cmd/wfparser/
``` 

### CLI
There is an experimental CLI available, called `wfcli`.
The intent is to integrate it into the Fission CLI, removing the need for the separate CLI.
```bash
go install github.com/fission/fission-workflows/cmd/wfcli/
```
