# GRPC-BLOCKCHAIN (Golang)

## Overview

    This repository is based on client and server communication using grpc in golang.
    The data that we exchange between client and server is blockchain.
    Each block in blockchain contains "Hash", "PrevBlockHash" and "Data", where "Data" store current time.

## Folder Structure

### blockchain

    blockchain.go : blockchain package

### client

    main.go : main client file

### proto

    blockchain.proto : message definition

### server

    main.go : main server file

## Command

### server : startserver

    go run server/main.go

### client : add a block to blockchain

    go run client/main.go --add

### client :  list a blockchain

    go run client/main.go --list

## Reference

Base project: <https://github.com/plutov/packagemain/tree/master/00-grpc>

Tutorial video: <https://www.youtube.com/watch?v=gju-bml4kdw&>
