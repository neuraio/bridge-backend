FROM golang

WORKDIR /go/src/app
COPY ./ ./
RUN GOOS=linux go build -a -o ../server ./app/cmd/main.go
RUN rm -rf ./app & rm -rf ./etc & rm -rf k8.yml
EXPOSE 80
CMD ["../server"]

