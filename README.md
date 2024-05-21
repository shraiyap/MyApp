# MyApp

MyApp is a web application consisting of three independent components developed in Go, Next.js (TypeScript), and WordPress. Each component has its own CI/CD pipeline to ensure quality and efficiency. This repository contains the setup instructions, Docker configuration, and CI/CD workflows for these components.

## Table of Contents

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Setting Up](#setting-up)
- [CI/CD Pipelines](#cicd-pipelines)
  - [Go Application](#go-application)
  - [Next.js Application](#nextjs-application)
  - [WordPress Site](#wordpress-site)
- [Docker Setup](#docker-setup)
  - [Docker Compose](#docker-compose)
  - [Running Locally](#running-locally)
- [Deployment](#deployment)
- [Coding Standards](#coding-standards)

## Project Structure

```
.
├── .github
│   └── workflows
│       ├── go-ci.yml
│       ├── nextjs-ci.yml
│       └── wordpress-ci.yml
├── go-app
│   ├── Dockerfile
│   ├── main.go
│   ├── golangci.go
│   └── go.mod
├── nextjs-app
│   ├── Dockerfile
│   ├── .eslintrc.json
│   ├── .prettierrc
│   ├── package.json
│   └── src
├── wordpress-site
│   ├── Dockerfile
│   ├── composer.yaml
│   └── wp-content
├── docker-compose.yml
└── README.md
```

## Getting Started

### Prerequisites

- Docker and Docker Compose installed on your local machine.
- Git installed on your local machine.

### Setting Up

Clone the repository:

```sh
git clone https://github.com/your-username/MyApp.git
cd MyApp
```

## CI/CD Pipelines

### Go Application

The CI pipeline for the Go application is defined in `MyApp/.github/workflows/go-ci.yml`. It includes steps for:

- Checking out the code.
- Setting up Go.
- Installing dependencies.
- Running the linter (GolangCI-Lint).
- Running tests.
- Building the Docker image.
- Pushing the Docker image to Docker Hub.
- Deploying to the staging environment.

### Next.js Application

The CI pipeline for the Next.js application is defined in `MyApp/.github/workflows/nextjs-ci.yml`. It includes steps for:

- Checking out the code.
- Setting up Node.js.
- Installing dependencies.
- Running ESLint.
- Running Prettier.
- Running tests.
- Building the Docker image.
- Pushing the Docker image to Docker Hub.
- Deploying to the staging environment.

### WordPress Site

The CI pipeline for the WordPress site is defined in `MyApp/.github/workflows/wordpress-ci.yml`. It includes steps for:

- Checking out the code.
- Setting up PHP.
- Installing dependencies via Composer.
- Running PHPCS.
- Running tests.
- Building the Docker image.
- Pushing the Docker image to Docker Hub.
- Deploying to the staging environment.

## Docker Setup

### Docker Compose

The `docker-compose.yml` file defines the services for the Go application, Next.js application, and WordPress site.

### Running Locally

To spin up the entire application stack locally, run the following command:

```sh
docker-compose up
```

This will start all the services defined in the `docker-compose.yml` file.

## Deployment

Deployment is integrated into the CI/CD pipelines. After a successful build, lint, and test, the Docker image is pushed to Docker Hub, and the application is deployed to the staging environment. The deployment steps are included in each component's CI workflow.

## Coding Standards

### Go Application

- **Linting**: GolangCI-Lint
  - Configuration: `go-app/.golangci.yml`
  - CI Integration: Defined in `MyApp/.github/workflows/go-ci.yml`

### Next.js Application

- **Linting and Formatting**: ESLint and Prettier
  - Configuration: `nextjs-app/.eslintrc.json` and `nextjs-app/.prettierrc`
  - CI Integration: Defined in `MyApp/.github/workflows/nextjs-ci.yml`

### WordPress Site

- **Coding Standards**: PHPCS with WordPress coding standards
  - Configuration: `wordpress-site/composer.json`
  - CI Integration: Defined in `MyApp/.github/workflows/wordpress-ci.yml`
