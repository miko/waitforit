TAG=2.4.4
docker build -t miko/waitforit .
docker tag miko/waitforit miko/waitforit:${TAG}


