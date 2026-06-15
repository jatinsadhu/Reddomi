
# Reddit Block Investigation Guide

## Overview

This guide explains how to run a comprehensive investigation to determine exactly why Reddit is blocking the browser automation connection.

## Setup Summary

- Enhanced `browser_automation/reddit.go` with `dumpDiagnosticsOnBlock()` function
- Created test harness: `cmd/reddit_block_diagnostics/main.go`
- When Reddit blocks the connection, diagnostics are automatically captured
- All diagnostic data saved to `data/debugstore/` with session ID prefix

## Running the Investigation

### Step 1: Compile and Run

```bash
cd /workspaces/Reddomi/backend
source ../.envrc
go run ./cmd/reddit_block_diagnostics
```

### Step 2: Open LiveURL and Sign In

The test will print:
```
LiveURL: https://api.steel.dev/v1/sessions/<SESSION_ID>/player?interactive=true&showControls=false
```

- Open this URL in your browser (Chrome/Firefox/Safari)
- Sign in to Reddit normally
- The test will detect login automatically

### Step 3: Automatic Diagnostic Capture

If Reddit blocks the connection, the test will:
1. Detect the block
2. Automatically capture:
   - Full page HTML
   - All cookies
   - Screenshot
   - Detailed logs
3. Save all files to `data/debugstore/`

### Step 4: Review Diagnostic Files

Files created will be named:
- `block_diagnostics_<SESSION_ID>_page.html` - Full HTML of blocked page
- `block_diagnostics_<SESSION_ID>_cookies.json` - All cookies and session info
- `block_diagnostics_<SESSION_ID>_screenshot.png` - Screenshot of what user sees

### Step 5: Analyze Results

For each file, check:

**page.html**:
- Exact block message displayed
- Error code or error ID
- Which endpoint returned the block
- Any redirect URL
- JavaScript errors in console

**cookies.json**:
- Authentication tokens present?
- Session cookies valid?
- Cookie domain/path configuration
- Security flags (httpOnly, Secure, SameSite)

**screenshot.png**:
- Visual confirmation of block
- Type of block (Reddit native, CloudFlare, WAF, etc.)
- Message clarity

## Root Cause Diagnosis

Based on diagnostic output, determine:

| Evidence | Possible Cause | Solution |
|----------|---------------|----------|
| Block mentions "security" or "unusual" | IP reputation | Use paid provider with residential proxy |
| Block mentions "bot" or "automation" | Playwright detection | Enable stealth mode features |
| Block mentions CloudFlare | WAF/DDoS protection | Use provider with CloudFlare bypass |
| Geographic block | Geo-blocking | Use provider with different country |
| 429 status code | Rate limiting | Reduce request frequency |
| Missing auth cookies | Session issue | Fresh session creation |
| Screenshot shows login page | Page never loads | Navigation/timeout issue |

## Expected Investigation Time

- **Automated**: ~1-2 minutes to capture diagnostics
- **Manual**: ~5 minutes for you to open LiveURL and sign in
- **Analysis**: I'll analyze diagnostic data immediately

## Key Differences from Manual Testing

Unlike manual browser testing where you just see "blocked", this investigation captures:
- Full HTML (not just rendered view)
- All network information (cookies, headers, endpoints)
- Exact timing and sequence of events
- Browser properties being used
- Complete error stack

This allows precise root cause identification.
