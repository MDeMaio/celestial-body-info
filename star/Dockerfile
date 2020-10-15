FROM dependencies

COPY star_server/ star/star_server/
COPY starpb/ star/starpb/

WORKDIR /app/star/star_server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o server
EXPOSE 50052
ENTRYPOINT ["./server"]