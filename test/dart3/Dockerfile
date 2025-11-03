FROM dart:3.8.1

WORKDIR /workspace


RUN dart --disable-analytics

WORKDIR /workspace
COPY . /workspace
RUN dart pub get

