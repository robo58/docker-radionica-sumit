FROM python:3.10.8
# /home/user
WORKDIR /scripts
# cd /scripts
# /scripts

RUN set -x \
     && apt-get -y update \
     && apt-get -y upgrade \
     && apt-get install -y ffmpeg \
     && pip install youtube_dl

COPY . .
# cp . .

#CMD [ "python3", "script.py" ]
ENTRYPOINT ["python3", "script.py"]