# Auto CRUD API Generator in Golang
This Auto CRUD help create simple CRUD API's without writing backend, you just need to specify [config.yaml](https://github.com/uzrnem/auto-crud-golang/blob/main/files/config.yaml)

## Installation with Docker

Create Directory `files` and create / update [config.yaml](https://github.com/uzrnem/auto-crud-golang/blob/main/files/config.yaml) file, bind this directory with container directory
Sample `docker-compose.yml`
```sh
version: '3.7'

services:
  auto-crud:
    image: uzrnem/auto-crud:0.1
    container_name: auto-crud
    volumes:
      - $PWD/files:/app/files
    ports:
      - 8080:8080
```