#!/usr/bin/env bash

curl -H "Content-type: application/json" -d '{"username": "user1", "points": 25}' localhost:8080/points
curl -H "Content-type: application/json" -d '{"username": "user2", "points": 25}' localhost:8080/points
curl -H "Content-type: application/json" -d '{"username": "user3", "points": 18}' localhost:8080/points
curl -H "Content-type: application/json" -d '{"username": "user7", "points": 100}' localhost:8080/points
curl -H "Content-type: application/json" -d '{"username": "user1", "points": 81}' localhost:8080/points
echo "######"

curl -s localhost:8080/points/user3 | jq
curl -s localhost:8080/leaderboard | jq