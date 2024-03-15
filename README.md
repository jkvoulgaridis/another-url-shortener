# What? 

This is implements a functional url shortener.

# How? 

* Golang is used on Bacnkend and React on Frontend.
* For the persistance layer, there 2 implementations, 1 for Redis and one for Postgres DB. I tried to use the 3-tier architecture where the business logic is implemented isolated and uses
  Data Abtraction Layers (DAOs) to fetch or store items to persistance layers and is agnostic. So, it is straigtforward to change to your desired persistance layer from the initial config.
* Also, prometheus and grafana are integrated for monitoring.
* Docker and Docker compose are used to pack the stack and easily use it. 


# Getting Started

1. ```make build```
2. ```make up```
3. Navigate to localhost:3000 and start shorten urls.

# Disclaimer 
* The UI is pretty poor
* This is a toy project to get some friction with golang  programming language. It can be used a demo, tutorial or anything educational but nothing more. 
