FROM golang:latest

WORKDIR /app

COPY . .

# Run unit tests
RUN make ci-unit -i

# Run integration tests
# TODO