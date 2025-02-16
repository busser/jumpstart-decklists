package export

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/busser/jumpstart-decklists/pkg/mtg"
	"github.com/busser/jumpstart-decklists/pkg/render"
)

func DecklistsAsPNG(ctx context.Context, decklists []mtg.Decklist, destination string) error {
	if err := os.MkdirAll(destination, 0o755); err != nil {
		return fmt.Errorf("failed to create destination directory %q: %w", destination, err)
	}

	html, err := render.DecklistsAsWebPage(decklists)
	if err != nil {
		return fmt.Errorf("failed to render decklists as HTML: %w", err)
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "*.html")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(html); err != nil {
		return fmt.Errorf("failed to write HTML to temp file: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to save temp file: %w", err)
	}

	ctx, cancel := chromedp.NewContext(
		ctx,
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	log.Printf("starting up headless Chrome")
	if err := chromedp.Run(
		ctx,
		chromedp.Tasks{
			chromedp.Navigate("file://" + tmpFile.Name()),
			chromedp.EmulateViewport(10000, 10000),
			chromedp.WaitVisible("body"),
			chromedp.Sleep(time.Second), // hack to let background images load
		},
	); err != nil {
		return fmt.Errorf("failed to start headless Chrome: %w", err)
	}

	for i := range decklists {
		pngFilePath := filepath.Join(destination, fmt.Sprintf("%03d.png", i+1))
		log.Printf("saving decklist %d of %d to %q", i+1, len(decklists), pngFilePath)

		selector := fmt.Sprintf("#decklist-%d", i) // must match HTML template

		var png []byte
		if err := chromedp.Run(
			ctx,
			chromedp.Tasks{
				chromedp.Screenshot(selector, &png, chromedp.NodeVisible),
			},
		); err != nil {
			return fmt.Errorf("failed to take screenshot: %w", err)
		}

		if err := os.WriteFile(pngFilePath, png, 0o644); err != nil {
			return fmt.Errorf("failed to write PNG to %q: %w", pngFilePath, err)
		}
	}

	return nil
}

func HTMLToPNG(ctx context.Context, html []byte, selector string) ([]byte, error) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(html); err != nil {
		return nil, fmt.Errorf("failed to write HTML to temp file: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return nil, fmt.Errorf("failed to save temp file: %w", err)
	}

	ctx, cancel := chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	var png []byte
	if err := chromedp.Run(
		ctx,
		chromedp.Tasks{
			chromedp.Navigate("file://" + tmpFile.Name()),
			chromedp.Screenshot(selector, &png),
		},
	); err != nil {
		return nil, fmt.Errorf("failed to take screenshot: %w", err)
	}

	return png, nil
}
