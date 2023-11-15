FROM golang

COPY . /app

WORKDIR /app
RUN [ "rm", "Dockerfile" ]
RUN [ "rm", "docker-compose.yml" ]
RUN [ "rm", ".env"]
RUN [ "go", "mod", "tidy" ]
# ENV GIN_MODE=release
RUN [ "go", "build", "." ]
RUN [ "chmod", "u+x", "cuciin-backend" ]
ENTRYPOINT [ "./cuciin-backend" ]