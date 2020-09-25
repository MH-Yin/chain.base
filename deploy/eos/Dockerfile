FROM ubuntu:18.04

RUN apt-get update && apt-get install -y curl libicu60 libusb-1.0-0 libcurl3-gnutls

RUN curl -LO https://github.com/EOSIO/eos/releases/download/v1.8.14/eosio_1.8.14-1-ubuntu-18.04_amd64.deb \
    && dpkg -i eosio_1.8.14-1-ubuntu-18.04_amd64.deb

RUN curl -LO https://github.com/EOSIO/eosio.cdt/releases/download/v1.7.0/eosio.cdt_1.7.0-1-ubuntu-18.04_amd64.deb \
    && dpkg -i eosio.cdt_1.7.0-1-ubuntu-18.04_amd64.deb