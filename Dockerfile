##################################################
#Stage1 :builds go binary
##################################################
FROM golang:1.23 AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

##################################################
#Stage2 :builds minimal image
##################################################

FROM gcr.io/distroless/base
COPY --from=base /app/main .
COPY --from=base /app/profile.html ./profile.html
EXPOSE 8080
CMD [ "./main" ]