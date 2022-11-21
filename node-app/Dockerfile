FROM node:16.17.0

# Odabrati radno okruzenje unutar kontejnera
# `WORKDIR lokacija` isto kao cd /var/www
WORKDIR /var/www

# Trebamo kopirati package.json da znamo postavke i pakete koje nasa aplikacija koristi
# Wildcard (*) se koristi da budemo sigurni da se i package-lock kopira ukoliko postoji
# Sintaksa je naredba `COPY datoteka/e lokacija` ova lokacije se odnosi na WORKDIR koji smo postavili iznad
# znaci ovo ce biti kopirano u /var/www
COPY package*.json ./


# Instalacija paketa
RUN npm install

# Ukoliko zelimo da se i nas kod nalazi unutar image-a, mozda zbog lakseg deploymenta, 
# koristimo i ovu naredbu da kopiramo sav nas kod
COPY . .

# Naredba `EXPOSE port` otvara port kontejnera da bi se spojili na njega izvana
EXPOSE 8080

# na kraju Dockerfile-a pisemo naredbu da pokrenemo samu aplikaciju
# Imamo vise opcija
# Jedna je preko naredbe CMD
CMD [ "node", "app.js" ]

# Druga preko naredbe ENTRYPOINT
#ENTRYPOINT [ "node", "app.js" ]
