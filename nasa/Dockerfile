FROM dependencies

COPY nasa_server/ nasa/nasa_server/
COPY nasapb/ nasa/nasapb/

WORKDIR /app/nasa/nasa_server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o server
EXPOSE 50053
ENTRYPOINT ["./server"]