FROM golang AS build

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app

FROM debian

WORKDIR /app

COPY --from=build /src/app /bin/app

CMD ["app"]