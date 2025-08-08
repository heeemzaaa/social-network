#!/bin/zsh

set -e

echo "📦 Installing Docker (rootless)..."

# Download and install Docker rootless
curl -fsSL https://get.docker.com/rootless | sh

echo "✅ Docker rootless installed."

# Run rootless setup, skip iptables check, continue on error
echo "⚙️ Running dockerd-rootless-setuptool.sh with --skip-iptables..."
~/bin/dockerd-rootless-setuptool.sh install --skip-iptables || echo "⚠️ Skipping setup tool errors."

# Set up environment variables
echo "⚙️ Setting up environment variables..."

echo 'export PATH=$HOME/bin:$PATH' >> ~/.zshrc
echo 'export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock' >> ~/.zshrc
export PATH=$HOME/bin:$PATH
export DOCKER_HOST=unix://$XDG_RUNTIME_DIR/docker.sock

echo "✅ Environment configured."

# Start the rootless Docker daemon in background
echo "🚀 Starting Docker daemon (rootless) in background..."
nohup dockerd-rootless.sh > ~/docker-rootless.log 2>&1 &

echo "✅ Docker daemon started in background (log: ~/docker-rootless.log)"

# Optional: install Docker Compose v2
echo "📦 Installing Docker Compose v2..."
mkdir -p ~/.docker/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.26.1/docker-compose-linux-x86_64 \
  -o ~/.docker/cli-plugins/docker-compose
chmod +x ~/.docker/cli-plugins/docker-compose

echo 'export PATH=$HOME/.docker/cli-plugins:$PATH' >> ~/.zshrc
export PATH=$HOME/.docker/cli-plugins:$PATH

echo "✅ Docker Compose installed."

# Verify installation
echo "🔍 Verifying installation..."
docker --version || echo "⚠️ Docker not found in PATH yet."
docker compose version || echo "⚠️ Docker Compose not found in PATH yet."

echo ""
echo "✅ All done! Run 'source ~/.zshrc' or restart your terminal to apply environment changes."

