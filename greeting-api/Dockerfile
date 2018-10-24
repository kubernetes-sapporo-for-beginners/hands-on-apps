# Build Image ####################################
FROM golang:1.10.3-alpine3.8 as builder

RUN set -eux; \
	apk add --no-cache git

RUN go get github.com/ant0ine/go-json-rest/rest; 

ENV GO_APP_HOME=${GOPATH}/src/github.com/kubernetes-sapporo-for-beginners/hands-on-apps/greeting-api

WORKDIR ${GO_APP_HOME}
ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./main.go
RUN chmod 755 ${GO_APP_HOME}/main

# Execution Image #########################docke#######
FROM scratch
EXPOSE 8080
COPY --from=builder /go/src/github.com/kubernetes-sapporo-for-beginners/hands-on-apps/greeting-api/main /app/main
CMD [ "/app/main" ]
