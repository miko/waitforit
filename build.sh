TAG=v2.4.6
docker build -t miko/waitforit .
docker tag miko/waitforit miko/waitforit:${TAG}


