package browser_automation

import (
    "context"
    "errors"
    "go.uber.org/zap"
)

// FallbackBrowserAutomation tries primary provider first, then secondary on error.
type FallbackBrowserAutomation struct {
    primary   BrowserAutomationProvider
    secondary BrowserAutomationProvider
    logger    *zap.Logger
}

func NewFallbackBrowserAutomation(primary, secondary BrowserAutomationProvider, logger *zap.Logger) BrowserAutomationProvider {
    return &FallbackBrowserAutomation{primary: primary, secondary: secondary, logger: logger}
}

func (f *FallbackBrowserAutomation) GetCDPInfo(ctx context.Context, input CDPInput) (*CDPInfo, error) {
    info, err := f.primary.GetCDPInfo(ctx, input)
    if err == nil {
        return info, nil
    }

    // If primary fails due to a provider plan limitation or similar, log and try secondary.
    f.logger.Warn("primary browser automation provider failed, attempting secondary", zap.Error(err))

    // If no secondary configured, return original error
    if f.secondary == nil {
        return nil, err
    }

    // Try secondary provider
    info2, err2 := f.secondary.GetCDPInfo(ctx, input)
    if err2 != nil {
        // aggregate errors
        return nil, errors.New(err.Error() + "; fallback error: " + err2.Error())
    }

    f.logger.Info("fallback provider succeeded")
    return info2, nil
}

// ReleaseSession is optional; implement if providers expose it via CDPInfo.ReleaseSession
// This wrapper doesn't add extra release logic.
