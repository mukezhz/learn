# Java project deployment

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
4. run the docker
```
docker run --name javademo <yourname>/javademo:<tag>
```

---
:( it won't work now because there is no database container

[but you can the build process is sucessful]

I will be creating the k8s deployment soon.

Stay tunes 😋😋😋