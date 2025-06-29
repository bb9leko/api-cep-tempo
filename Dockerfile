FROM golang:1.24 as build
gcloud auth loginWORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-cep-tempo 

FROM scratch
WORKDIR /app
COPY --from=build /app/api-cep-tempo .
ENTRYPOINT ["./api-cep-tempo"]