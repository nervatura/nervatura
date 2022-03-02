FROM golang:1.17-alpine AS builder

ARG APP_MODULES=all

# Move to working directory (/build).
WORKDIR /build

RUN apk add --no-cache upx make

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environmet variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN make release TAGS=${APP_MODULES}

FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/nervatura", "/"]

# Export necessary port.
EXPOSE 5000

# Command to run when starting the container.
ENTRYPOINT ["/nervatura"]
