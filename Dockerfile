FROM golang:1.22.2-alpine AS start

WORKDIR /forum-projct/

RUN apk add gcc musl-dev
COPY . .

RUN go build -o forum .

FROM alpine
WORKDIR /myProject
COPY --from=start /forum-projct/* /myProject/
LABEL version="1.0"
LABEL projectname="EduTalks"
CMD ["./forum-projct"]