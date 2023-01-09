FROM golang:1.13-alpine
FROM node:16

# Set the Current Working Directory inside the container
WORKDIR /email_search_engine

COPY ./interfaz_seeker/package*.json ./interfaz_seeker/
RUN npm install ./interfaz_seeker/

#back:
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod ./
COPY go.sum ./

COPY . ./
RUN cd interfaz_seeker/ && npm run build

# This container exposes port 8080 to the outside world
EXPOSE 3000

CMD ["./interfaz"]
