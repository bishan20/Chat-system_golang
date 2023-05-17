# Chat-System for Fintech

A chat application for private messaging between users.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [API Endpoints](#api-endpoints)
- [Web Socket](#web-socket)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## Overview

This project is a chat application that allows users to send private messages to each other securely. It enables users to discuss details of their transactions, such as invoices and banking details, within the app itself.

## Features

- 1-on-1 private messaging between users.
- Sender can edit or delete the message after it has been sent.
- Sender receives an acknowledgement when the message is successfully delivered to the recipient.

## Technologies Used

- Go programming language
- Gin framework for building APIs
- SQLC for database query generation
- Docker for containerization
- Viper for configuration management
- WebSocket for one-to-one communication

## Configuration

The project uses the following environment variables for configuration. Make sure to set these variables before running the application.

- `DB_DRIVER`: Specifies the database driver to be used. In this case, it is set to postgres, indicating that the PostgreSQL database driver will be used.
- `DB_SOURCE`: Represents the connection string or the data source URL for connecting to the PostgreSQL database.
- `SERVER_ADDRESS`: Specifies the network address and port on which the server will listen for incoming HTTP requests.
- ...

You can also provide configuration through a configuration file. Refer to the `config.yaml` file for an example configuration.

## API Endpoints

The following API endpoints are available:

### User Endpoints

- `POST /v1/users`: Creates a new user. This endpoint is used to register a new user in the system.

### WebSocket Endpoint

- `GET /v1/ws`: Establishes a WebSocket connection for one-to-one communication between users. This endpoint is used for real-time messaging.

### Miscellaneous Endpoints

- `GET /v1/hello`: Returns a simple "Hello" message. This endpoint is used for testing purposes or as a basic health check.

### Message Endpoints

- `POST /v1/messages`: Retrieves a list of messages. This endpoint is used to fetch a list of messages.

- `PATCH /v1/messages`: Updates a message. This endpoint is used to edit or modify an existing message.

- `DELETE /v1/messages`: Deletes a message. This endpoint is used to remove a message from the system.



## Web Socket

The project uses WebSocket for one-to-one communication between users. The WebSocket endpoint is:

- `ws://localhost:8000/v1/ws`

1. Sender id should be included in the URL query params. As an example, `ws://localhost:8080/v1/ws?sender_id=1`
2. Message body should be in JSON format. As an example,
 `{
    "message":"good?",
    "receiver_id":3
  }`
3. After message is succcessfully delivered to the recepient, the recepient receives the message in the format which looks like
    `{
    "id": 9,
    "message": " i am good",
    "sender_id": 1,
    "receiver_id": 3,
    "sent_at": "2023-05-17T13:31:44.673873Z"
}`
4. After message is successfully delivered to the recepient, sender receives an acknowledgement message which looks like 
  `{
    "Msg": "How are you?",
    "IsDelivered": true
   }`

## Getting Started

Follow these instructions to get the project up and running on your local machine.

# Steps to run

1. Install Docker in your system.
2. Clone this repo using the command "git clone https://github.com/bishan20/Chat-system_golang.git".
3. Open the project in Visual Studio Code.
4. For Windows, OPEN start.sh and SELECT " LF " in end of line sequence.
5. To run your project in docker, use command "docker-compose up" in your terminal.
6. To test all REST API, check out the "postman-collection" folder and for web-socket, check out the instructions written above [Web-Socket](#web-socket).
7. After you are done with the testing, stop the server using the command "docker-compose down" in your terminal.
