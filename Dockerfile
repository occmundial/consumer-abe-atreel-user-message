##################################
# STEP 1 build executable binary #
##################################
# Get image of golang
FROM golang:1.20 as builder
# go modules = on
ENV GO111MODULE=on
# Set workdir
WORKDIR /go/src/consumer-abe-atreel-user-message
# Copy all from local to image
ADD . .
# Get libraries (including test libraries)
RUN go get -t ./... && \
# Run tests
go test ./services/... && \
# Create binary (binary name = "api-build")
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o api-build .

##############################
# STEP 2 build a small image #
##############################
# image small of linux
FROM alpine

ARG version

RUN apk add --no-cache tzdata
ENV TZ="America/Mexico_City"

# Copy our static executable and required files.
WORKDIR /go/src/consumer-abe-atreel-user-message
ADD ./config ./config
RUN rm ./config/config.go

# 
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Set environment variables
ENV VERSION=$version

# Copy binary and json files
COPY --from=builder /go/src/consumer-abe-atreel-user-message/api-build /go/src/consumer-abe-atreel-user-message/api-build
# Run the app binary.
CMD ["./api-build"]
EXPOSE 8023
