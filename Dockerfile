FROM golang:alpine as builder

RUN apk update && apk add --no-cache git && apk add --no-cache redis

WORKDIR /app

COPY . .

RUN go mod tidy

#RUN go build cmd/http/*
RUN CGO_ENABLED=0 go build main.go


FROM busybox:1.35.0

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/conf /conf

EXPOSE 9003 6379

CMD [ "./main" ]