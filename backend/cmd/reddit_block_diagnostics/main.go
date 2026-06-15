package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/shank318/doota/browser_automation"
	"github.com/streamingfast/dstore"
	"go.uber.org/zap"
)

func main() {
	// Setup logger
	logger := zap.NewExample()
	defer logger.Sync()

	// Load tokens
	steelToken := os.Getenv("DOOTA_START_COMMON_STEEL_API_KEY")
	browserlessToken := os.Getenv("DOOTA_START_COMMON_BROWSERLESS_API_KEY")
	browserlessWarmup := os.Getenv("DOOTA_START_COMMON_BROWSERLESS_WARMUP_API_KEY")
	if browserlessWarmup == "" {
		browserlessWarmup = "2SIxpPBYG6XJqLj5ec45cd436c170abdbec8713fd1bbaffe4"
	}

	if steelToken == "" && browserlessToken == "" {
		fmt.Println("No browser automation provider token found.")
		fmt.Println("Set DOOTA_START_COMMON_STEEL_API_KEY or DOOTA_START_COMMON_BROWSERLESS_API_KEY")
		os.Exit(1)
	}

	// Setup debug store (all diagnostics will be saved here)
	debugStore, err := dstore.NewStore("data/debugstore", "", "", true)
	if err != nil {
		fmt.Printf("Failed to create debug store: %v\n", err)
		os.Exit(2)
	}

	// Create browser providers
	steel := browser_automation.NewSteelBrowser(steelToken, logger)
	browserless := browser_automation.NewBrowserLessBrowser(browserlessToken, browserlessWarmup, logger)
	provider := browser_automation.NewFallbackBrowserAutomation(steel, browserless, logger)

	// Create Reddit automation client
	redditBA := browser_automation.NewRedditBrowserAutomation(provider, logger, debugStore)

	fmt.Println("================================================================================")
	fmt.Println("REDDIT BLOCK DIAGNOSTICS TEST")
	fmt.Println("================================================================================")
	fmt.Println("")
	fmt.Println("This test will:")
	fmt.Println("  1. Create a browser session")
	fmt.Println("  2. Navigate to Reddit login")
	fmt.Println("  3. If blocked, capture comprehensive diagnostics")
	fmt.Println("  4. Save diagnostics to data/debugstore/")
	fmt.Println("")
	fmt.Println("Diagnostics captured on block:")
	fmt.Println("  - Full page HTML (block_diagnostics_*.html)")
	fmt.Println("  - All cookies and session info (block_diagnostics_*.json)")
	fmt.Println("  - Full page screenshot (block_diagnostics_*.png)")
	fmt.Println("  - Complete diagnostic logs")
	fmt.Println("")
	fmt.Println("80 character separator line")
	fmt.Println("Starting test...")
	fmt.Println("80 character separator line")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Step 1: Start login session (creates browser with LiveURL)
	fmt.Println("[1/2] Creating browser session and generating LiveURL...")
	cdp, err := redditBA.StartLogin(ctx, "")
	if err != nil {
		fmt.Printf("ERROR: Failed to start login: %v\n", err)
		os.Exit(3)
	}

	fmt.Printf("SUCCESS: Browser session created\n")
	fmt.Printf("  Session ID: %s\n", cdp.SessionID)
	fmt.Printf("  LiveURL: %s\n", cdp.LiveURL)
	fmt.Println("")

	// Step 2: Wait for cookies (this is where diagnostics will be captured if blocked)
	fmt.Println("[2/2] Waiting for Reddit login (polling for navigation)...")
	fmt.Println("")
	fmt.Println("IMPORTANT: Open the LiveURL in your browser and sign in to Reddit.")
	fmt.Println("The test will automatically detect the login and extract cookies.")
	fmt.Println("")
	fmt.Println("If Reddit shows 'Your request has been blocked by network security',")
	fmt.Println("the test will automatically capture detailed diagnostics.")
	fmt.Println("")
	fmt.Printf("Polling timeout: 5 minutes\n")
	fmt.Println("")

	cfg, err := redditBA.WaitAndGetCookies(ctx, cdp)
	if err != nil {
		fmt.Println("")
		fmt.Println("80 character separator line")
		fmt.Println("ERROR: Login failed or blocked")
		fmt.Println("80 character separator line")
		fmt.Printf("Error message: %v\n", err)
		fmt.Println("")
		fmt.Println("DIAGNOSTIC DATA SAVED")
		fmt.Println("80 character separator line")
		fmt.Println("The test captured comprehensive diagnostics in data/debugstore/")
		fmt.Println("Look for files named: block_diagnostics_<session-id>_*")
		fmt.Println("")
		fmt.Println("These files contain:")
		fmt.Println("  - Full page HTML at time of block")
		fmt.Println("  - All cookies and session information")
		fmt.Println("  - Screenshots showing the block message")
		fmt.Println("  - Complete diagnostic logs")
		fmt.Println("")
		fmt.Printf("Session ID for reference: %s\n", cdp.SessionID)
		fmt.Println("")
		os.Exit(4)
	}

	fmt.Println("")
	fmt.Println("80-char-separator-line-success-login")
	fmt.Println("SUCCESS: Login detected and cookies extracted!")
	fmt.Println("80-char-separator-line-success-login")
	fmt.Printf("Username: %s\n", cfg.Username)
	fmt.Printf("Cookies extracted: %d bytes\n", len(cfg.Cookies))
	fmt.Println("")
	fmt.Println("If you were testing for blocks, you would need to:")
	fmt.Println("  1. Use a Reddit account that triggers the block")
	fmt.Println("  2. Or use a provider/IP that Reddit blocks")
	fmt.Println("")
	os.Exit(0)
}
