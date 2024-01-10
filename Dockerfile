FROM golang:1.21 AS build
RUN mkdir -p /project/
WORKDIR /project/
COPY go.mod go.sum *.go /project/
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o opsgenie-reminder

FROM scratch AS prod
WORKDIR /project/
COPY --from=build /project/opsgenie-reminder /project/opsgenie-reminder
ENTRYPOINT ["./opsgenie-reminder"]
CMD ["--help"]
