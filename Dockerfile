FROM alpine:3.5
LABEL source_repository="https://github.com/sapcc/gotee"

COPY bin/gotee-linux /usr/local/bin/gotee

ENTRYPOINT [ "gotee" ]
CMD [ "--listen", "8000", "-1", "8001", "-2", "8002" ]
