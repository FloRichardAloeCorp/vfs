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

The client is developed with Next.js. There’s nothing particularly unique about it—it simply allows browsing and manipulating your files.

## Launch the project

### Prerequisites

* `Docker` installed on your machine

### Launch

To run the project, clone it and run `docker compose up` command in a terminal at the root of the project.

To access the VFS, open a web browser and go to <http://localhost:3000>.
