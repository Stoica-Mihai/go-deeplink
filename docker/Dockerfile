FROM alpine:latest
LABEL author "github.com/Stoica-Mihai"

# Add curl
RUN apk add curl

# Set up the working directory
WORKDIR /tmp

# Set up GO file name
ENV GONAME=go1.22.4.linux-amd64.tar.gz

# Download and install GO
RUN wget -q https://go.dev/dl/$GONAME
RUN tar -C /usr/local -xzf $GONAME

# Set up PATH
ENV PATH=$PATH:/usr/local/go/bin

# Set working dir
WORKDIR /app

# Install GO dependecies
# RUN go mod download