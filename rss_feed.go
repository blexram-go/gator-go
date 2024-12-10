package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Add("User-Agent", "gator")

	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return &RSSFeed{}, err
	}
	rssFeed.Channel.Title = decodeEscapedHTML(rssFeed.Channel.Title)
	rssFeed.Channel.Description = decodeEscapedHTML(rssFeed.Channel.Description)
	for i, rssItem := range rssFeed.Channel.Item {
		rssItem.Title = decodeEscapedHTML(rssItem.Title)
		rssItem.Description = decodeEscapedHTML(rssItem.Description)
		rssFeed.Channel.Item[i] = rssItem
	}

	return &rssFeed, nil
}

func decodeEscapedHTML(s string) string {
	return html.UnescapeString(s)
}
