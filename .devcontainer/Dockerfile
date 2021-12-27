# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.208.0/containers/debian/.devcontainer/base.Dockerfile

# [Choice] Debian version (use bullseye or stretch on local arm64/Apple Silicon): bullseye, buster, stretch
ARG VARIANT="bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

# ** [Optional] Uncomment this section to install additional packages. **
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

RUN curl -LO https://go.dev/dl/go1.18beta1.linux-amd64.tar.gz && sudo tar -C /usr/local -xzf go1.18beta1.linux-amd64.tar.gz
