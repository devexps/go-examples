# Go-Micro on Kubernetes

## Overview

This project mainly demonstrates how the microservice (Go-Micro) framework runs on a Kubernetes cluster and performs
service discovery registration and configuration management by calling APIServer.

Ensure that you have installed the Docker on your local machine before, but you can do it
now: https://docs.docker.com/engine/install

You also need an account and logging in to Docker Hub, for example:

```shell
docker login -u devexps -p MY_PASSWORD
```

## Build & Docker services

### Admin Service

```shell
cd admin-srv
make docker
make docker-push
```

### User Service

```shell
cd user-srv
make docker 
make docker-push
```

### Token Service

```shell
cd token-srv
make docker
make docker-push
```

## Kubernetes

### Start k8s

There are so many ways to start a mini kubernetes cluster on your dev environment (your local machine.)
I prefer one that I usually do:

```shell
minikube start --network=socket_vmnet
```

According to check whether your minikube is working well, you can do this command: `kubectl get svc` and you'll see the
result like this:

```text
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   40h
```

### Apply configuration resources

We have to create a `Namespace`, a `Service Account`, and roles to get access permissions like `list`, `update`, `watch`
,... which you can work closely with `configmap` and `pod`.

```shell
kubectl apply -f .k8s/namespace.yaml
kubectl apply -f .k8s/serviceaccount.yaml
kubectl apply -f .k8s/configmap-rbac.yaml
kubectl apply -f .k8s/pod-rbac.yaml
```

Yeah! we're so lucky, the result will be showed:

```text
namespace/go-micro created
serviceaccount/go-micro-services created
role.rbac.authorization.k8s.io/go-micro-config created
rolebinding.rbac.authorization.k8s.io/go-micro-config created
clusterrole.rbac.authorization.k8s.io/go-micro-registry created
clusterrolebinding.rbac.authorization.k8s.io/go-micro-registry created
```

In the same, let's apply the configurations for three of our services

```shell
kubectl apply -f .k8s/admin-srv-configmap.yaml
kubectl apply -f .k8s/admin-srv-deployment.yaml
kubectl apply -f .k8s/admin-srv-service.yaml
kubectl apply -f .k8s/user-srv-configmap.yaml
kubectl apply -f .k8s/user-srv-deployment.yaml
kubectl apply -f .k8s/user-srv-service.yaml
kubectl apply -f .k8s/token-srv-configmap.yaml
kubectl apply -f .k8s/token-srv-deployment.yaml
kubectl apply -f .k8s/token-srv-service.yaml
```

Using this command to check the result: `kubectl get pods -n go-micro -o wide`

```text
NAME                         READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE   READINESS GATES
admin-srv-65df4c5bcd-gs26q   1/1     Running   0          42s   10.244.0.8   minikube   <none>           <none>
token-srv-54b8c998f5-qpjnv   1/1     Running   0          27m   10.244.0.5   minikube   <none>           <none>
user-srv-5fc4d67dd5-wkg7c    1/1     Running   0          27m   10.244.0.4   minikube   <none>           <none>
```

Using this command to get the target host:port for service `admin-srv`: `minikube service admin-srv -n go-micro`

```text
|-----------|-----------|---------------------|---------------------------|
| NAMESPACE |   NAME    |     TARGET PORT     |            URL            |
|-----------|-----------|---------------------|---------------------------|
| go-micro  | admin-srv | admin-srv-http/8080 | http://192.168.49.2:30000 |
|           |           | admin-srv-grpc/9090 | http://192.168.49.2:30001 |
|-----------|-----------|---------------------|---------------------------|
🏃  Starting tunnel for service admin-srv.
|-----------|-----------|-------------|------------------------|
| NAMESPACE |   NAME    | TARGET PORT |          URL           |
|-----------|-----------|-------------|------------------------|
| go-micro  | admin-srv |             | http://127.0.0.1:64630 |
|           |           |             | http://127.0.0.1:64631 |
|-----------|-----------|-------------|------------------------|
[go-micro admin-srv  http://127.0.0.1:64630   http://127.0.0.1:64631]
```

### Testing

#### HTTP

We can run an HTTP API test with this endpoint: `http://127.0.0.1:64630`

```shell
curl -X 'POST' \
  'http://127.0.0.1:64630/admin/v1/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "userName": "thangn",
  "password": "123456"
}'
```

And the result will be shown:

```json
{
  "id": "1",
  "userName": "thangn",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2OTUzNTkyNDYsInBsYXRmb3JtIjoiVE9LRU5fUExBVEZPUk1fV0VCIiwic3ViIjoidGhhbmduIiwidWlkIjoiMSJ9.kieQAm75DWlb6Pa7hmX76dpzNDJXwGNpbLpfA-Jcy_o"
}
```

#### gRPC

We can also run a gRPC API test with this endpoint: `http://127.0.0.1:64631`

```shell
grpcurl -plaintext -d '{"userName":"thangn","password":"123456"}' 127.0.0.1:64631 admin_srv.v1.AuthenticationService/Login
```

And then we'll see the result like this:

```json
{
  "id": "1",
  "userName": "thangn",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2OTUzNTk0NjcsInBsYXRmb3JtIjoiVE9LRU5fUExBVEZPUk1fV0VCIiwic3ViIjoidGhhbmduIiwidWlkIjoiMSJ9.tLKmZUe2TW6X4DOWVDfkLBz8fKwPN6LgYjcQuLTZlu0"
}
```

## Docker Compose

We can modify the `type` value inside the `remote.yaml` file in each of the services, then re-dockerize them.

```yaml
config:
  #  type: "k8s"
  type: "consul"

  ...
```

Run the docker-compose and we'll have all the same above

```shell
docker-compose up -d
```

- HTTP endpoint: 127.0.0.1:8080
- gRPC endpoint: 127.0.0.1:9090