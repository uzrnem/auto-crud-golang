# docker build . -t uzrnem/auto-crud:0.1
FROM golang:1.21rc2-alpine3.18 as dev

LABEL image_name="AutoCRUD"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd cmd
COPY pkg pkg
COPY files files

RUN go build -o autocrud cmd/app/main.go

FROM scratch as prod

COPY --from=dev /app/autocrud autocrud

EXPOSE 9055

CMD [ "./autocrud" ]
