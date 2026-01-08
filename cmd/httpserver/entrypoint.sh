#!/bin/bash
echo "Starting container..."
./svr || echo "Server crashed with exit code $?"
echo "Server exited, keeping container alive for debugging"
# Keep container running for debugging
tail -f /dev/null