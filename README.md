# Objective 
Helping git repos markdown auto reload the repos info

# Live demo
https://github.com/AlphaWong/go-web-framework-stars

# Run
```sh
docker build -t app .
sudo docker run --env-file env.list -p 80:8080 -d app
```

# Remove
```sh
sudo docker rm -f name-of-container
```