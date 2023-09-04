# Basic GitOps Operator

This simple Go program acts as a basic GitOps operator. It continually synchronizes a specified Git repository to a local path and then applies Kubernetes manifests from that repository.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)


## Features

- Automatic syncing of a given Git repository.
- Application of Kubernetes manifests using `kubectl`.
- Logging of sync and apply operations.
- Efficient error handling and reporting.

## Prerequisites

- Go (at least version 1.15)
- `kubectl` installed and configured with appropriate cluster access.
- `github.com/go-git/go-git/v5` Go package.

## Installation

1. Clone this repository:
git clone https://github.com/yourusername/basic-gitops-operator.git
cd basic-gitops-operator


2. Install the required Go packages:
go get -v

3. Build the program:
go build

This will generate an executable named `basic-gitops-operator` (or `basic-gitops-operator.exe` on Windows).

## Usage

1. Run the program:
./basic-gitops-operator

2. By default, the program will sync the repository at `https://github.com/richinex/basic-gitops-operator.git` every 5 seconds and apply the manifests found in the `basic-gitops-operator-config` directory. Modify the source code if you need to sync a different repository or adjust the interval.

3. Monitor the logs for information on syncing and applying operations. Errors, if any, will also be reported here.