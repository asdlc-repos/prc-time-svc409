# Overview

A lightweight web service that tells you the current time in UTC when you ask for it.

# Personas

- Developer — calls the API to get the current UTC timestamp.

# Features

- Ask the service for the current time and get back a JSON response.
- The time comes back in RFC3339 format so it's easy to parse.
- The service runs on port 8080 and responds to GET requests at /time.