# Kubernets

in this secton we will try a simple example of kubernetes
- deployment
- service
- ingress
- hpa

## Deployment

deployment is provides declaraive updates for Pods and ReplicaSet
We describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate.

example of deployment

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
  namespace: yournamespace
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

In this example:

- A Deployment named `nginx-deployment` is created, indicated by the `.metadata.name` field.
- The Deployment creates three replicated Pods, indicated by the `replicas` field.
- The `selector` field defines how the Deployment finds which Pods to manage. In this case, you simply select a label that is defined in the Pod template `(app: nginx)`. However, more sophisticated selection rules are possible, as long as the Pod template itself satisfies the rule.
- The `template` field contains the following sub-fields:
    - The Pods are labeled `app: nginx` using the `labels` field.
    - The Pod template’s specification, or `.template.spec` field, indicates that the Pods run one container, nginx, which runs the `nginx` Docker Hub image at version 1.7.9.
    -Create one container and name it `nginx` using the `name` field.

To Apply any resource use:

```
kubectl apply -f filename.yml
```

To get any resource use
```
kubectl get resource -n namespace
```
in this case 
```
kubectl get deployments -n namespace
```


## Service

Kubernetes Pods are mortal. They are born and when they die, they are not resurrected. If you use a Deployment to run your app, it can create and destroy Pods dynamically.

Each Pod gets its own IP address, however in a Deployment, the set of Pods running in one moment in time could be different from the set of Pods running that application a moment later.

This leads to a problem: if some set of Pods (call them “backends”) provides functionality to other Pods (call them “frontends”) inside your cluster, how do the frontends find out and keep track of which IP address to connect to, so that the frontend can use the backend part of the workload?

### def

In Kubernetes, a Service is an abstraction which defines a logical set of Pods and a policy by which to access them (sometimes this pattern is called a micro-service). The set of Pods targeted by a Service is usually determined by a selector 

service example

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: nginx
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

What if we want to connect to external service

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
---
apiVersion: v1
kind: Endpoints
metadata:
  name: my-service
subsets:
  - addresses:
      - ip: 192.0.2.42
    ports:
      - port: 9376
```

## Ingress

Ingress exposes HTTP and HTTPS routes from outside the cluster to services within the cluster. Traffic routing is controlled by rules defined on the Ingress resource.

INgress example 
```
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: hpa-ingress
  namespace: test
spec:
  rules:
  - host: hpa-example.com
    http:
      paths:
      - backend:
          serviceName: hpa-example-service
          servicePort: 8080

```

its expose service `hpa-example-service` with host hpa-example.com.
set it on the `/etc/host` and you will see the result.
```
{IP} hpa-example.com
```

## HPA

The Horizontal Pod Autoscaler automatically scales the number of pods in a replication controller, deployment or replica set based on observed CPU utilization (or, with custom metrics support, on some other application-provided metrics).
