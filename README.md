# Uvod u docker

Korisne naredbe
 - 
 - `docker pull <image_name>` (povlaci image sa dockerhub-a)
 - `docker run <image_name>` (pokrece kontajner iz image-a)
 - `docker build . -t <tag_name>` (kreira novi image iz Dockerfile-a, -t je oznaka za tag, praksa je username/ime-kontejnera, npr robo58/python-app)
 - `docker ps` (izlistava popis aktivnih kontejnera)
 - `docker stop <container_id>` (zaustavlja kontejner sa naglasenim id-om)
 - `docker rm <container_id>` (brise kontejner)

 ## docker run korisni flag-ovi
- `-it` (interaktivni shell, otvara terminal unutar kontejnera)
- `-v local_path container_path` (mountanje mape, volume-a, prvo ide lokalna mapa koju zelimo mountat, drugo ide mapa unutar kontejnera)
- `--rm` (obrise kontejner nakon izvrsavanja, pozeljno koristiti svaki put)

## kreiranje node+react+mysql crud aplikacije

- `cd docker-compose-react-nodejs-mysql/bezkoder-api`
- `docker build -t <username>/bezkoder-node-test`
- `docker run -p 3307:3306 -v /home/robo58/Projects/docker-sumit/mysql:/var/lib/mysql --network=backend -e MYSQL_ROOT_PASSWORD='123456' -e MYSQL_DATABASE='bezkoder_db' --env-file .env --name mysqldb mysql:5.7`
- `docker run -p 6868:8080 -e DB_HOST='mysqldb' -e DB_USER='root' -e DB_PASSWORD='123456' -e DB_NAME='testdb' -e DB_PORT='3306' -e CLIENT_ORIGIN='http://127.0.0.1:8081' --env-file ./bezkoder-api/.env --network=backend --name node-backend <username>/bezkoder-node-test`
- `cd docker-compose-react-nodejs-mysql/bezkoder-ui`
- `docker build --build-arg REACT_APP_API_BASE_URL='http://127.0.0.1:6868/api' -t <username>/bezkoder-react-test`
- `docker run -p 8081:80 -e REACT_DOCKER_PORT='80' -e REACT_APP_API_BASE_URL='http://127.0.0.1:6868/api' --network=frontend --name react-frontend <username>/bezkoder-react-test`
