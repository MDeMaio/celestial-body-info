FROM dependencies

COPY planet_server/ planet/planet_server/
COPY planetpb/ planet/planetpb/

WORKDIR /app/planet/planet_server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o server
EXPOSE 50051
ENTRYPOINT ["./server"]