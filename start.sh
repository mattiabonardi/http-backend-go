#!/bin/bash
echo Starting http-backend-go container
docker run -d -p 8080:8080 http-backend-go