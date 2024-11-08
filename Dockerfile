FROM golang:alpine 

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .
COPY --from=build /app/app.env app.env

RUN go mod tidy 

RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]