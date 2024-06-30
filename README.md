# Webhook Multiplexer

Webhook Multiplexer is a Go-based application that allows you to create and manage webhooks, endpoints, and responses.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Running the Application](#running-the-application)
5. [API Documentation](#api-documentation)

## Prerequisites

- Go 1.21.4 or later
- Docker and Docker Compose
- PostgreSQL 15.0 or later

## Installation

1. Clone the repository:
    
```
git clone https://github.com/Sandy143toce/webhook-multiplexer.git
cd webhook-multiplexer
```

2. Install the dependencies:

```
go mod download
```

## Configuration

The application uses environment variables for configuration. These are set in the `docker-compose.yml` file:

- `DB_USER`: PostgreSQL username
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: Database name
- `DB_HOST`: Database host
- `DB_PORT`: Database port

## Running the Application

1. Start the application and database using Docker Compose:

```
docker-compose up --build
```
2. Once docker-compose is up and running, you need to run the migrations to create the tables in the database. You can do this by running the following command:

```
cd schema
Replace the Database Variables like DB User, DB Password, DB Name, DB Host, DB Port in the URL in apply.sh file
Give Acces to the apply.sh file by running the command chmod +x apply.sh
Run the apply.sh file by running the command ./apply.sh
```

3. The application will be available at `http://localhost:9000`

## API Documentation

The application provides the following API endpoints:

1. Create Webhook
- Method: POST
- URL: `/webhook-multiplexer/create-webhook`
- Body:
  ```json
  {
    "name": "string",
    "url": "string"
  }
  ```

2. Map Endpoint to Master Webhook
- Method: POST
- URL: `/webhook-multiplexer/add-customer-endpoint`
- Body:
  ```json
  {
    "webhook_id": "string",
    "url": "string"
  }
  ```

3. Send Event/Webhook Forwarding
- Method: POST
- URL: `/webhook-multiplexer/send-event`
- Body:
  ```json
  {
    "event_name": "string",
    "metadata": json, // can send any data its optional
  }
  ```


