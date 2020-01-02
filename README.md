# Realize

**Discontinued in favor of [loop](https://github.com/grzegorz-zur/loop/)**

Task runner.

## Goal

The goal of this fork is to simplify the tool and maintain it to support the latest Go relaese.

## Installation

To install or update run the following command.

```
go get -u github.com/grzegorz-zur/realize/cmd/realize
```

## Configuration

Create `.realize.yml`.

### Sample

```
schema:
- name: realize
  path: .
  commands:
    fmt:
      status: true
    generate:
      status: true
    vet:
      status: true
    test: 
      status: true
    build:
      status: true
    install:
      status: true
  watcher:
    paths:
      - /
    extensions:
      - go
      - mod
    ignored_paths:
      - .git
```

## Running

```
realize
```
