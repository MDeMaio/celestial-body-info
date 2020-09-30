FROM dependencies

# Copy the application source code.
COPY . .
# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o main
EXPOSE 50051
ENTRYPOINT [ "./main" ]