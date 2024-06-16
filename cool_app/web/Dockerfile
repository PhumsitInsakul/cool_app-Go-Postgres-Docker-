# build server on top of linux with go installed
# base image
FROM golang:1.22

# take our app code
# copy code from src dir on our computer
# to /usr/src/app dir in the server (container)
COPY ./src /usr/src/app

#go to this directory
WORKDIR /usr/src/app

#compile the code
RUN go build -o /bin/app main.go

# run the app
CMD ["/bin/app"]

# docker build -t coolapp:latest .
# docker run coolapp:latest -p 3000:3000