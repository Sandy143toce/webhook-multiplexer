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
2. Once docker-compose is up and running, you need to run the migrations to create the tables in the database. 
First Replace the Database Variables like DB User, DB Password, DB Name, DB Host, DB Port in the URL in apply.sh file in schema folder.
Then Give Acces to the apply.sh file by running the command chmod +x apply.sh. Then run the below command to apply the migrations.

```
cd schema
./apply.sh
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
    "name": "webhook1",
    "url": "http://localhost:9000/webhook-multiplexer/send-event"
  }
  ```

![Alt text](/webhook.png "Create Webhook")

2. Map Endpoint to Master Webhook
- Method: POST
- URL: `/webhook-multiplexer/add-customer-endpoint`
- Body:
  ```json
  {
    "webhook_id": "id returned by create webhook api",
    "url": "any_url"
  }
  ```

 ![Alt text](/endpoint.png "Create Webhook")
 

3. Send Event/Webhook Forwarding
- Method: POST
- URL: `/webhook-multiplexer/send-event`
- Body:
  ```json
  {
    "event_name": "event1",
    "metadata": json, // can send any data its optional
  }
  ```

 ![Alt text](/event.png "Create Webhook")



