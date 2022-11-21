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

