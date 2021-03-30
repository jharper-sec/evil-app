# Introduction
Evil App is an intentionally vulnerable Golang application intended for learning about security vulnerabilities within Golang. Currently implemented vulnerabilities are:
* SQL Injection
* Reflected Cross-Site Scripting (XSS)

Upcoming vulnerabilities:
* Command Injection
* Path Traversal

# Pre-Requisites
* Go 1.16

# Build
## 1. Run locally
```bash
go build
```

# Run
```bash
./evil-app
```

# Build with Contrast
Must have `contrast-go` installed.
```bash
contrast-go build -o evil-app
```

# Run with Contrast
1. Download `contrast_security.yaml` from Contrast to application directory

1. Start Contrast Service
```bash
contrast-service
```

1. Start application
```bash
./evil-app
```
