FROM python:3.10.8

WORKDIR /scripts

RUN set -x \
     && apt-get -y update \
     && apt-get -y upgrade \
     && apt-get install -y ffmpeg \
     && pip install youtube_dl

COPY . .

ENTRYPOINT ["python3", "script.py"]