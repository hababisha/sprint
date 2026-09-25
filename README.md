# Spring CLI

A simple Go CLI for quickly creating Spring Boot projects using [Spring Initializr](https://start.spring.io).

Instead of manually opening Spring Initializr, configuring the same options, downloading the ZIP, extracting it, and deleting the ZIP file, you can simply run:

```bash
spring init --name myapp
```
The CLI handles the rest

## current config
* **Build tool:** Maven
* **Language:** Java
* **Spring Boot:** 4.1.1
* **Packaging:** JAR
* **Java:** 21
* **Dependencies:**

  * Spring Web
  * Spring Data JPA
  * PostgreSQL Driver

The generated project is extracted into a directory with the project name.

For example:

```text
my-project/
├── pom.xml
├── mvnw
├── mvnw.cmd
├── src/
└── ...
```

## Requirements

Currently, Spring CLI is designed for Linux and requires:

* Go
* `unzip`

Check that Go is installed:

```bash
go version
```

Check that `unzip` is installed:

```bash
unzip -v
```

If `unzip` is not installed on Ubuntu/Debian:

```bash
sudo apt install unzip
```

## Getting the Project

Clone the repository:

```bash
git clone https://github.com/hababisha/sprint.git
or
git clone git@github.com:hababisha/sprint.git
```
## Build it and add it to your path

```bash
go build -o spring ./cmd/spring

echo $PATH
```
copy the binary to one of the paths to make it global



