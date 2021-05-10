# ButterCMS Hugo Starter Project

This is a starter project for Hugo with integration of ButterCMS. This includes a posts generator written in Go which uses the ButterCMS SDK for Go to retrieve posts from ButterCMS.

## Installation

1. Install Hugo.

   You can check on how to install Hugo on your operating system [here](https://gohugo.io/getting-started/installing/)

2. Install Go

   You can check on how to install Go compiler [here](https://golang.org/doc/install)

## Setup

You should replace YOUR_AUTH_TOKEN in go/generate.go with your auth token provided by ButterCMS.


## Usage

1. Run the generate script to retrieve all of the posts from ButterCMS
```shell
go run go/generate.go
```

2. Check the posts if they are loaded into HUGO by running
```shell
hugo serve
```

Web Server will be available at http://localhost:1313/