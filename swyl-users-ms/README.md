<p align="center">
<br />
<a href="https://github.com/SWYLy/servers"><img src="https://github.com/SWYLy/materials/blob/master/logo.svg?raw=true" width="150" alt=""/></a>
<h1 align="center">SWYL - Support Who You Love - v2.0</h1>
<h4 align="center">User Microservice</h4>

## Folder structure

    .
    ├── controllers
    │   └── swyl.user-controller.go
    ├── dao
    │   ├── swyl.dao-impl.go
    │   └── swyl.dao-interface.go
    ├── db
    │   └── swyl.user-db.go
    ├── middleware
    │   └── swyl.auth-middleware.go
    ├── models
    │   └── swyl.user-model.go
    ├── routers
    │   └── swyl.router.go
    ├── utils
    │   └── swyl.dao-impl.go
    ├── .example.env
    ├── .gitignore
    ├── Makefile
    ├── README.md
    ├── go.mod
    ├── go.sum
    └── swyl.user-main.go

## Getting Started

### Requirement

- [git](https://git-scm.com/)
- [golang](https://go.dev/)
<!-- - [docker](https://www.docker.com/) -->

### Clone the repo

```bash
git clone https://github.com/SWYLy/servers.git
cd swyl-users-ms
```

### Set up environment variables

At the root of the directory, create a .env file using .env.example as the template and fill out the variables.

### Running the project

Build and run `swyl-users-ms` locally using `Make` scripts

```bash
make dev-mode
```

<!-- 2. Build and run `agent` on Docker using `Make` scripts

```bash
make build-app
``` -->
