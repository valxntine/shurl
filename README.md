# Cronofy URL Shortener Project

## Requirements

This project assumes you have `docker` and `docker compose` installed locally

## Running the Project

To run the project locally, clone the repo

`git clone git@github.com:valxntine/shurl.git`

Then either using the provided `Makefile` bring up the project

`make run`

alternatively use `docker compose` directly

`docker compose up --build --no-recreate -d api client`

Open your browser at `http://localhost:4242` and use the web app to generate short url's
