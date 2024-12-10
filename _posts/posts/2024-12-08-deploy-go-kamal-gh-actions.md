---
layout: post
title: "Deploying Go with Kamal via GitHub Actions"
date: '2024-12-08'
tags: ["go", "kamal", "github-actions", "devops"]
image:
  path: https://github.com/aziflaj/aziflaj.github.io/blob/main/assets/images/20241208/3way.png?raw=true
  alt: "Three-way handshake"
---

Deployments, the last mile of delivering software projects. It's not surprising
most developers don't know how to do deployments, let alone do them properly. 
It used to be a simple "FTP the files to the server" process, but now there are
a lot of moving parts, especially if you're working with a modern stack. Docker 
is supposed to be a simple way to **package & run** apps, and Kubernetes is supposed to
be a **simple** way to and manage them once the packaged apps reach the server.
But it's not that simple, is it? Even if you go [the LeMuR way](https://aziflaj.github.io/posts/2024-06-20-lemur-ship-your-machine-to-customers/),
it's still not that simple.

In the Rails world, we used to have [Capistrano](https://capistranorb.com/).
It would SSH into the server, clone the repository, run the migrations,
and restart the server. It was simple, and it worked. But it didn't adapt to the
ever-changing landscape of software development. Last year, the Rails benevolent
dictator David Heinemeier Hansson (_the infamous DHH_)
announced [Kamal](https://world.hey.com/dhh/kamal-1-0-5304ff9e), the 
spiritual successor to Capistrano for the modern world. It does almost
exactly the same thing as Capistrano, but it's container-aware and it comes with a proxy
to manage a "cluster" of containers in multiple hosts.

Even though Kamal came from the Rails world, since it only cares about containers,
it can be used with any language. I've been using it with Go for a while now, and
in this post, I'll show you how to deploy a Go app with Kamal via GitHub Actions.
There are some caveats here and there, but we have the power of Trial and Error on
our side.

For a real-world, modern production app, we need:
- A Frontend app
- A Backend app
- A Database and migrations for it
- A background worker (optional, but most likely it'll be there)

So let's get busy.

## Prepping the Frontend

Most of the frontends you'll see in the wild are written in some SPA-ish framework
like React, Angular, or Vue. It doesn't matter which one you choose; most likely
it will be built into a bunch of static files that need to be served from a server.

With Go, we can embed these files into the same binary as the backend, so we don't need to serve
them from a separate server. That means we don't need to worry
about CORS, cookies, or any other frontend-backend communication issues. But that's too easy...

Most of the time you want to be able to deploy the frontend separately from the
backend. You might even want to deploy the frontend to a CDN, so you don't need to
worry about serving the files at all. Kamal can't help you with that CDN part, but
we can still use it to deploy the frontend to a server... Nginx, for example.

Speaking of `Dockerfile`s, here's one for the frontend:

```Dockerfile
FROM node:22-alpine AS dev
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .


FROM node:22-alpine AS build
ARG BACKEND_URL=http://localhost:8080
ENV PUBLIC_API_URL=$BACKEND_URL
WORKDIR /app
COPY --from=dev /app /app
RUN npm run build


FROM nginx:alpine AS prod
COPY --from=build /app/build /etc/nginx/html/
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

Here we're using a multi-stage build to keep the image size small. There's a dev
stage where we install the dependencies, a build stage where we build the frontend,
and a prod stage where we copy the built files to the Nginx server. Unsurprisingly,
the prod stage needs a `nginx.conf` file to work properly:

```nginx
events {}

http {
  include /etc/nginx/mime.types;
  default_type application/octet-stream;

  server {
    listen 80;

    gzip_static on;
    gzip_disable "MSIE [1-6]\\.(?!.*SV1)";
    gzip_proxied any;
    gzip_comp_level 5;
    gzip_types text/plain text/css application/javascript application/x-javascript text/xml application/xml application/rss+xml text/javascript image/x-icon image/bmp image/svg+xml;
    gzip_vary on;

    location / {
      try_files $uri $uri/ /index.html;
    }
  }
}
```

You will probably have a SPA router that will handle the routes, so you need to
redirect all requests to the `index.html` file. Everything else is pretty standard
Nginx configuration, listening on port 80 for HTTP requests and gzipping the files
before sending them to the client.

## Handling migrations

In standard library Go, there's no built-in way to handle migrations like in some backend frameworks.
You can use [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
which can help you generate and run migrations, but I prefer something else.

I have been using [uptrace/bun](https://github.com/uptrace/bun) as a lightweight
ORM, and they have a way to (optionally) support migrations. Two birds, one stone.
It is minimalistic enough to not need much more than that.

When I build Go backends, I ship 3 different binaries in the same image:
- The backend itself, which will handle your HTTP requests
- The migrations binary, which will run the migrations
- The worker binary, which will handle background tasks (though that's missing from this post)

To handle migrations, I (re)use a simple `cli` that will run the migrations
when the binary is executed:

```go
// cmd/cli/main.go
package main

var Migrations = migrate.NewMigrations()

//go:embed migrations/*.sql
var sqlMigrations embed.FS

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}

//... rest of https://github.com/uptrace/bun/blob/master/example/migrate/main.go
```

Along with this cli, and the obvious backend code, you will need a `Dockerfile` to
build the image:

```Dockerfile
FROM golang:1.23-bookworm AS base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .


FROM base AS cli-builder
RUN CGO_ENABLED=0 go build -o my-cli cmd/cli/main.go
RUN mv my-cli /usr/local/bin/


FROM base AS server-builder
RUN CGO_ENABLED=0 go build -o my-server cmd/server/main.go
RUN mv my-server /usr/local/bin/


FROM gcr.io/distroless/static-debian12 AS prod
COPY --from=server-builder /usr/local/bin/my-server /usr/local/bin/
COPY --from=cli-builder /usr/local/bin/my-cli /usr/local/bin/

# or whatever port you're using
EXPOSE 8000
CMD ["my-server"]
```

You might have seen people using `scratch` as the base image, but I prefer to use
`distroless` because... you can [read the reasons here](https://iximiuz.com/en/posts/containers-distroless-images/).

> What you have so far is a containerized frontend and backend, each ready to be
deployed into a server. If you push them to your Docker Registry of choice, and you
pull them in a server, you can run them via `docker run` and they will work as
expected.
> 
> But we've been committed to DevOps since 2016 and we want to spend 5 hours to
automate a 5-minute task, right?!

## Setting up Github Container Registry

While you can use any Registry you want, the title of this post is _"Deploying Go
with Kamal via **GitHub** Actions"_, so we're going to use GitHub Container Registry (GHCR).
It's free (up to a point) and unlike Docker Hub, it allows you to push more than one
private image for the low price of $0.00.

You can read more about how to set up and authenticate with GHCR [here](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-in-a-github-actions-workflow),
but the gist is that you need a Personal Access Token (PAT) with access to:
- `read:packages`
- `write:packages`
- `delete:packages`
- `repo`

You'll want to store this PAT in your GitHub repository's `Secrets > Actions`,
so your GitHub Actions can use it to push images to GHCR.

## Kamal-ifying our deployment

Kamal can be used as a Ruby gem, or as a Docker image. I already have Ruby installed
so I don't mind using the gem. You can read more about Kamal in their [official docs](https://kamal-deploy.org/docs/installation/).

With Kamal installed, run `kamal init` for both your frontend and backend projects;
they must be 2 different projects, regardless of being a monorepo or not. Kamal
can only deploy one project at a time. The frontend Kamal config is not going to be that
interesting, since it only needs to build the image, push it to GHCR, and deploy it
to a server. The backend Kamal config will have to also set up an _accessory service_:

```yaml
# backend/config/deploy.yml
service: my-backend

# Name of the container image.
image: aziflaj/my-backend

# Deploy to these servers.
servers:
  web:
    - 123.4.20.69

proxy: 
  ssl: true
  host: api.my-backend.com
  app_port: 8000

registry:
  server: ghcr.io
  username: aziflaj
  password:
    - KAMAL_REGISTRY_PASSWORD

builder:
  arch: amd64
  target: prod

env:
  secret:
    - POSTGRES_DSN

accessories:
  db:
    host: 123.4.20.69
    image: postgres:13
    port: 5432
    env:
      POSTGRES_USER: pguser
      POSTGRES_PASSWORD: pgpwd
      POSTGRES_DB: my_db
    volumes:
      - pgdata:/var/lib/postgresql/data
```

A lot of magic in a single file. When you run `kamal setup`, Kamal will set up
all your hosts (in this case, only one web server on `123.4.20.69`) by installing
Docker, running the accessories (the database in this case), setting up the
`kamal-proxy` with automagically managed SSL certs, and deploying the app. 

All that by a single command and a single file.

When you _commit changes_, you can run `kamal deploy` and watch the magic happen again and again.

> You need to commit changes, otherwise Kamal will ignore your changes when it deployes.
I ~~spent~~ lost an hour debugging faulty deployments before I realized this, you 
don't have to do the same.

## Automating via GitHub Actions

Kamal SSH-es into your server to deploy the app, so you need to set up SSH keys in
your server and in your GitHub Actions. If you, like me, get access into an empty server
and do things manually, start by generating an SSH key for the GH Action:

```bash
ssh-keygen -t ed25519 -C "youremail+ghactions@mail.com"
```

Copy the **public key** to your server:

```bash
ssh-copy-id -i /path/to/.ssh/key.pub root@123.4.20.69
```

Now, the same way you added your PAT to your GitHub Secrets, add the **private key**
to your GitHub Secrets. You can name it `PRIVATE_SSH_KEY` or something similar.

Before deploying, we'll need to set up the tools that Kamal needs to deploy. 
We'll need to install Docker, Kamal, and to give our GitHub Actions the ability
to SSH into the server. Since this'll be done for both the frontend and the backend,
I prefer putting it in a reusable action:

```yaml
# .github/actions/setup-cd/action.yml
name: Setup CD Pipeline
description: "Setup the CD pipeline by: \
              - Setting up Docker Buildx \
              - Setting up Ruby \
              - Installing Kamal"
inputs:
  ssh_key:
    description: "The SSH key to use for deployment"
    required: true

runs:
  using: "composite"
  steps:
    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v2

    - name: Set up Ruby
      uses: ruby/setup-ruby@v1
      with:
        ruby-version: 3

    - name: Install Kamal
      shell: bash
      run: |
        gem install kamal

    - name: Add SSH key
      shell: bash
      env:
        SSH_AUTH_SOCK: /tmp/ssh_agent.sock
      run: |
        mkdir -p /home/runner/.ssh
        ssh-keyscan 199.247.6.56 >> /home/runner/.ssh/known_hosts
        echo "${{ inputs.ssh_key }}" > /home/runner/.ssh/github_actions
        chmod 600 /home/runner/.ssh/github_actions
        ssh-agent -a $SSH_AUTH_SOCK > /dev/null	
        ssh-add /home/runner/.ssh/github_actions
```

Kudos to Max Schmitt for that last step, I ~~stole~~ borrowed it from [his blog post](https://maxschmitt.me/posts/github-actions-ssh-key).

And finally, for the deployment workflow:

```yaml
# .github/workflows/cd.yml
name: Build and Push Docker Containers

on:
  push:
    branches:
      - main

env:
  GHCR_REGISTRY_PASSWORD: ${{ secrets.GHCR_PAT }}

jobs:
  build-backend:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Setup CD
        uses: ./.github/actions/setup-cd
        with:
          ssh_key: ${{ secrets.PRIVATE_SSH_KEY }}

      - name: Build and push backend service
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: |
          cd backend && kamal deploy

      - name: Migrate database
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run:
          cd backend && kamal app exec my-cli db migrate

  build-frontend:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Setup CD
        uses: ./.github/actions/setup-cd
        with:
          ssh_key: ${{ secrets.PRIVATE_SSH_KEY }}

      - name: Build and push frontend service
        env:
          SSH_AUTH_SOCK: /tmp/ssh_agent.sock
        run: |
          cd frontend && kamal deploy
```

In the `build-backend` job, there's a `kamal deploy` command which will build the
image, push it to GHCR, and deploy it to the server. After that, we run the migrations
via `kamal app exec my-cli db migrate`. Since running migrations is idempotent,
there's no harm in running them every time you deploy.

And with all that, you have a fully automated deployment pipeline to deploy your
Go app, run migrations, and deploy your frontend. You can extend this pipeline to
deploy your worker, or to run tests before deploying, or [hack the planet](https://www.youtube.com/watch?v=5y_SbnPx_cE)
