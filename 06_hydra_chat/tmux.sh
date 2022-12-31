#!/usr/bin/env bash

tmux new-session -s HydraChat -d ./bin/server
tmux splitw -h ./bin/client
tmux splitw -v ./bin/client
tmux attach-session -d