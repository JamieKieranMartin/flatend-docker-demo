# flatend-docker-demo

## Starting the flatend instance

- docker build -t flatend .
- docker run -p 3000:3000 -p 9000:9000 -t flatend

## Starting your app instance

- docker build -t services services
- docker run --network host -t services