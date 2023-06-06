# Helloworld

## Install Micro
```bash
go install github.com/devexps/go-micro/cmd/micro/v2@latest
```

## Build & Run
```bash
# navigate into service folder
cd helloworld

# Generate proto code, wire, etc.
make all

# Run
make run

# Or 
micro run
```
