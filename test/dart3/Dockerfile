FROM dart:3.8.1

WORKDIR /workspace


RUN dart --disable-analytics

WORKDIR /workspace
COPY . /workspace
COPY pubspec.yaml .
RUN dart pub get

