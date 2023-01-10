FROM ubuntu:latest
USER root

# Set the Current Working Directory inside the container
WORKDIR /email_search_engine

COPY ./interfaz_seeker/package*.json ./interfaz_seeker/
# install go and node
RUN apt-get update
RUN apt-get -y install curl gnupg
RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash - 
RUN apt-get -y install nodejs
RUN npm install ./interfaz_seeker/
RUN apt-get -y install wget
RUN apt-get -y install zip
RUN curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile

COPY go.mod ./
COPY go.sum ./

COPY . ./
RUN cd interfaz_seeker/ && npm run build
RUN chmod +x install_zincsearch.sh
RUN cd ..

EXPOSE 3000
EXPOSE 4080

CMD ["./install_zincsearch.sh"]
