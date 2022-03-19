# Java project deployment

## Run app using docker

0. clone the repo
```
git clone https://gitlab.com/mukezhz/hamropatrofellowship.git
```
1. install docker in your machine
2. goto javaprojectdemo dir
```
cd hamropatrofelloship/javaprojectdemo
```
3. copy `database.properties_sample` to `database.properties`
```
cp app/src/main/resources/database.properties_sample app/src/main/resources/database.properties
```
4. build the docker file
```
docker build -t <yourname>/javademo:<tag> . 
```
5. create docker network
```
docker network create <networkname>
```

6. run mysql container
```
docker run --name <containerdbname> -d -e MYSQL_ROOT_PASSWORD=<password> mysql:latest
```

7. add mysql to the network
```
docker network connect <networkname> <containerdbname>
```

8. run the docker using network name
```
docker run --name <java_appcontainername> --network <networkname> <yourname>/javademo:<tag>
```

---
# Run app using kubernetes
---

0. go to k8s dir
```
cd k8s/
```

1. add kubernetes secret and configmap
```
kubectl apply -f mysql-config.yml 
kubectl apply -f mysql-secret.yml 
```

2. run mysql service and deployment:
```
kubectl apply -f mysql.yml
```

3. run java application
```
kubectl apply -f java-app.yml
```

**NOTE**: you must install desired requirement of kubernetes before executing the above commands

I am using: 

- kubectl [kubernetes client] &

- minikube [kubernetes component packaged in a single application best for the developement]

### For arch based linux to install minikube and kubectl

```
sudo pacman -S minikube kubectl
````

Start minikube using:
```
minikube start
```

---
### Happy leaning 😋😋😋