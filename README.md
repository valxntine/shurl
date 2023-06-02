# URL Shortener Project

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

To generate a short URL, enter any url in the first input field and click `Shorten!`

Your short URL will appear in the second field and the `Copy` button will copy the short URL to your clipboard.

When you visit the short URL you'll be redirected to the original URL you shortened.

![image](https://github.com/valxntine/shurl/assets/24855366/5188611c-23f6-4741-9190-45f86657e1e6)
