# Grofers Assignment

Problem Statement Solution:

I have built a 

## Installation

Install docker on your workstation
Link for installation guide: https://docs.docker.com/engine/install/

### Running the Web Server
Go to the downloaded assignment directory and then run the following command:

``docker compose up``

### Running the CLI Client
Run the following docker command:
``docker run sarthak321321/GroferClient``

## Usage
Run the server on one terminal and on another terminal run the CLI client. For `watch` operation, run it on another terminal.
You will get a CLI upon running the client side, which supports 4 operations:
```
1.) put <key> <value>
2.) get <key>
3.) watch
4.) STOP
```

> Note: The commands are case-sensitive

## Tech Stack Used

```
Golang for Server side and client side implementation
Redis for storing the key value pair.
```
> I used redis since it is an in-memory data structure store, cache and message broker, with optional durability.
