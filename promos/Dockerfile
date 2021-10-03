FROM golang:1.17-alpine3.14

#Install bash
RUN ["apk","add","--no-cache","bash"]

#Install tzdata for change timezone
RUN ["apk","add","tzdata"]

#Change timezone
RUN ["cp","/usr/share/zoneinfo/America/Mexico_City","/etc/localtime"]

#Create Dir
RUN ["mkdir","/opt/promos"]

WORKDIR /opt/promos

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ["go.mod","go.sum","./"]

RUN ["go","mod","download"]

COPY . .

# Build the Go app
#RUN ["go","build","-o","./out/promos"]
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o ./out/promos
# Run the binary program produced by `go install`
CMD ["./out/promos"]