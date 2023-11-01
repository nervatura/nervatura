FROM golang:1.21.3-alpine AS builder

ARG APP_MODULES=all
ARG APP_VNUM=dev
ARG APP_NAME=nervatura

# Move to working directory (/build).
WORKDIR /build

RUN apk add --no-cache upx 
# RUN apk add --no-cache make gcc g++

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environmet variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags "$APP_MODULES" -ldflags="-w -s -X main.version=$APP_VNUM" -o $APP_NAME main.go
RUN upx --best --lzma $APP_NAME
#RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=c-shared -tags "$APP_MODULES" -ldflags="-w -s -X main.version=$APP_VNUM" -o $APP_NAME.so shared/main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/nervatura", "/"]
#COPY --from=builder ["/build/nervatura.so", "/shared/nervatura.so"]

# Export necessary port.
EXPOSE 5000

# Command to run when starting the container.
ENTRYPOINT ["/nervatura"]
