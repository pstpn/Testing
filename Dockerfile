FROM golang:latest

WORKDIR /app

COPY . .

# Run unit tests
RUN make ci-unit

# Run integration tests
# TODO