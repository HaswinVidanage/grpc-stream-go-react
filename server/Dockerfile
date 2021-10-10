FROM golang:1.17

# All these steps will be cached
RUN mkdir /server
WORKDIR /server
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .



ENTRYPOINT [ "go", "run", "server.go" ]