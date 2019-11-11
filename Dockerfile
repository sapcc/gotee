FROM alpine:3.5

COPY gotee /usr/local/bin/gotee

ENTRYPOINT [ "gotee" ]
CMD [ "--listen", "8000", "-1", "8001", "-2", "8002" ]
