# Objective 
Helping git repos markdown auto reload the repos info

# Live demo
https://github.com/AlphaWong/go-web-framework-stars

# Run
env.list
```sh
# Please insert your key.
GITHUB_CLIENT_ID=YOU_API_CLIENT
GITHUB_CLIENT_SECRET=YOU_API_SECRET
```
build & run
```sh
docker build -t app .
sudo docker run --env-file env.list -p 80:8080 -d app
```

# Remove
```sh
sudo docker rm -f name-of-container
```