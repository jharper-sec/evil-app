# Introduction
Evil App is an intentionally vulnerable Golang application intended for learning about security vulnerabilities within Golang. Currently implemented vulnerabilities are:
* SQL Injection
* Reflected Cross-Site Scripting (XSS)

Upcoming vulnerabilities:
* Command Injection
* Path Traversal

# Pre-Requisites
## Normal
* Go >= 1.16

## Contrast
* contrast-go >= 0.14.0
* contrast-service >= 2.19.0

# Normal Build/Run Instructions
## Build
```bash
go build
```

## Run
```bash
./evil-app
```

# Contrast Build/Run Instructions
## Build with Contrast
Must have `contrast-go` installed.
```bash
contrast-go build -o evil-app
```

## Run with Contrast
1. Download `contrast_security.yaml` from Contrast to application directory

1. Start Contrast Service
```bash
contrast-service
```

1. Start application
```bash
./evil-app
```
