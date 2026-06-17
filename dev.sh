#!/usr/bin/env bash

SESSION="kaizen-development"

tmux new-session -d -s $SESSION
tmux rename-window -t $SESSION "kaizen"
tmux split-window -h -t $SESSION
tmux send-keys -t $SESSION:0.0 "cd backend && air" C-m
tmux send-keys -t $SESSION:0.1 "cd frontend && pnpm dev" C-m

tmux attach-session -t $SESSION
