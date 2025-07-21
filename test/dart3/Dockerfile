FROM dart:3.8.1

WORKDIR /workspace


RUN dart --disable-analytics
COPY pubspec.yaml .
RUN dart pub get

