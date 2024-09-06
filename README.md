# VFS (Virtual file server)

This VFS implentation is an in-memory VFS shipped with its REST API and a web client.

DISCLAIMER: This project is for demo purposes, it is not intented to be used in production.

## Motivations

I've developed many things throughout my career, but I've realized that I can't showcase much of it. Most of the projects I've worked on are private. One day, I asked myself: How can people evaluate my work if they can't see what I'm capable of? That's why I built this VFS.

For those who are curious and reading this, I will guide you through the application, demonstrating my skills—or at least a subset of them.

I hope that you will enjoy to read my code.

## The project

### Architecture

The project is composed of three parts: the `vfs` module, the `server`, and the `client`. The backend is developed in Go, while the frontend is developed in Next.js.

### VFS

The `vfs` module is the database engine. It’s a tree of nodes representing the directories and files of the filesystem. It provides convenient functions to manipulate files.

### Server

The server is a Go REST API developed with `gin`. I built it following the principles of `screaming architecture`. The server allows manipulation of the VFS via HTTP.

### Client

The client is developed with `Next.js`. There’s nothing particularly unique about it—it simply allows browsing and manipulating your files.

### Features

The product allows you do to the following things:

* Upload files
* Create directories
* Delete files or directories
* Nest files and directories
* Renaming files and directories
* Download files

## Launch the project

### Prerequisites

* `Docker` installed on your machine

### Launch

To run the project, clone it and run `docker compose up` command in a terminal at the root of the project.

To access the VFS, open a web browser and go to <http://localhost:3000>.

### FAQ

#### Why did you choose mono repo for this project ?

I chose to organize the project in a monorepo for simplicity and ease of demonstration. It makes navigation more straightforward. However, in a production environment, I would split the VFS, server, and client into separate repositories. This approach provides clearer versioning and management of each component individually.

#### Why did you choose this stack?

For the backend, I have over 4 years of experience developing backend applications in Go, primarily focusing on microservices. As a dedicated Gopher, I believe Go is one of the best languages for building backend applications.

For the frontend, although I love React and also appreciate Vue.js (Nuxt), I chose to use Next.js. In my opinion, Next.js allows for faster development compared to Nuxt, even though Nuxt offers more structure to the code.

#### Choose one: backend or frontend?

The eternal war between backend and frontend engineers. Based on my previous experiences, I would say backend. It is my specialty and where I am most proficient. However, I enjoy taking breaks to work on frontend as well. I like frontend too. I would position myself as a backend-oriented fullstack engineer (if that’s a valid term).

#### What about deployment?

As you’ve seen, I know how to use Docker. To deploy real applications to the world, I am familiar with using a Kubernetes cluster and deploying apps on it using `GitOps` principles. Additionally, when a Kubernetes cluster is not needed, I prefer using PaaS platforms such as Vercel or Heroku.

For Kubernetes, here is what I do:

* Use [Flux](https://fluxcd.io/) for continuous delivery and GitOps
* Secret encryption with [SOPS](https://github.com/getsops/sops)
* Monitoring, alerting, and logging stack: Prometheus, Alloy, Loki, and a Grafana instance
* Security scans: [Kubescape](https://github.com/kubescape/kubescape)
* Ingress: Traefik, LetsEncrypt, CertManager

This stack has been used in production and works well.
