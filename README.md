# cask 
## Description
This repository contains a simple PoC demonstrating how to communicate with the ChatOn API in order to leverage their 
AI compute.

The ChatOn API is a simple API that wraps the OpenAI and Claude AI APIs. 

Communicating with the API is done through a simple HTTP POST request, however the integrity of the request is verified
using a HMAC-SHA256 signature passed in the Authorization header.

The code provided in this repository demonstrates how to generate the HMAC-SHA256 signature and the sending of a valid
request to the ChatOn API.

## Usage
To compile and install:
```shell
git clone https://github.com/lumaaaaaa/cask
cd cask
go install
```

Then simply:
```shell
cask <prompt>
```

Beyond this, take a look at the `ask()` function in "api.go" and utilize the functions however you please.
