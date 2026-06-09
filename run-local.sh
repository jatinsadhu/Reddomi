#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

PATH_add() {
  export PATH="$PATH:$*"
}

if [[ -f .envrc ]]; then
  # shellcheck source=/dev/null
  source .envrc
fi

export NEXT_PUBLIC_API_URL="${NEXT_PUBLIC_API_URL:-http://127.0.0.1:8787}"
export NEXT_PUBLIC_APP_URL="${NEXT_PUBLIC_APP_URL:-http://127.0.0.1:3000}"

if ! command -v docker >/dev/null 2>&1; then
  echo "docker is required but not installed" >&2
  exit 1
fi

if ! command -v pnpm >/dev/null 2>&1; then
  echo "pnpm is required but not installed" >&2
  exit 1
fi

if ! command -v go >/dev/null 2>&1; then
  echo "go is required but not installed" >&2
  exit 1
fi

echo "Starting local dependencies..."
docker compose up -d postgres redis

echo "Building backend..."
cd "$ROOT/backend"
go build -o "$ROOT/build/doota" ./cmd/doota

echo "Applying database migrations..."
"$ROOT/backend/script/migrate.sh" up --yes

BACKEND_LOG="$ROOT/backend.log"
FRONTEND_LOG="$ROOT/frontend.log"

echo "Starting backend..."
"$ROOT/build/doota" start \
  --portal-auth0-domain="${DOOTA_START_PORTAL_AUTH0_DOMAIN:-}" \
  --portal-auth0-portal-client-id="${DOOTA_START_PORTAL_AUTH0_PORTAL_CLIENT_ID:-}" \
  --portal-auth0-portal-client-secret="${DOOTA_START_PORTAL_AUTH0_PORTAL_CLIENT_SECRET:-}" \
  --portal-auth0-api-redirect-uri="${DOOTA_START_PORTAL_AUTH0_API_REDIRECT_URI:-http://127.0.0.1:8787/auth/callback}" \
  --portal-reddit-redirect-url="${DOOTA_START_PORTAL_REDDIT_REDIRECT_URL:-http://127.0.0.1:3000/auth/callback}" \
  --portal-reddit-client-id="${DOOTA_START_PORTAL_REDDIT_CLIENT_ID:-}" \
  --portal-reddit-client-secret="${DOOTA_START_PORTAL_REDDIT_CLIENT_SECRET:-}" \
  --google-client-id="${GOOGLE_CLIENT_ID_DEV:-}" \
  --google-client-secret="${GOOGLE_CLIENT_SECRET_DEV:-}" \
  --common-browserless-api-key="${DOOTA_START_COMMON_BROWSERLESS_API_KEY:-}" \
  --common-browserless-warmup-api-key="${DOOTA_START_COMMON_BROWSERLESS_WARMUP_API_KEY:-2SIxpPBYG6XJqLj5ec45cd436c170abdbec8713fd1bbaffe4}" \
  --common-steel-api-key="${DOOTA_START_COMMON_STEEL_API_KEY:-}" \
  --common-resend-api-key="${DOOTA_START_COMMON_RESEND_API_KEY:-}" \
  > "$BACKEND_LOG" 2>&1 &
BACKEND_PID=$!

echo "Backend started with PID $BACKEND_PID (log: $BACKEND_LOG)"

cd "$ROOT/frontend/portal"
if [[ ! -d node_modules ]]; then
  echo "Installing frontend dependencies..."
  pnpm install
fi

echo "Starting frontend..."
export PORT=3000
if lsof -i tcp:3000 -sTCP:LISTEN >/dev/null 2>&1; then
  echo "ERROR: port 3000 is already in use. Stop the process using it or use a different port." >&2
  exit 1
fi
pnpm dev > "$FRONTEND_LOG" 2>&1 &
FRONTEND_PID=$!

echo "Frontend started with PID $FRONTEND_PID (log: $FRONTEND_LOG)"

echo "Waiting for services to come up..."
sleep 8

echo "--- Status ---"
echo "Backend log tail:"
tail -n 10 "$BACKEND_LOG" || true

echo "Frontend log tail:"
tail -n 10 "$FRONTEND_LOG" || true

echo "Listening ports:"
ss -ltnp | grep -E ':8787|:3000' || true

echo "You can stop with: kill $BACKEND_PID $FRONTEND_PID"

wait
