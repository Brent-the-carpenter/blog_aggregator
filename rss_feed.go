package main

import (
	"context"
	"encoding/xml"
	"fmt"
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
		Item        []RRSItem `xml:"item"`
	} `xml:"channel"`
}
type RRSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("There was a error creating the request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	res, err := httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error occurred while making the request to %s, error:%w", feedURL, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error occurred while reading response: %w", err)
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error occurred while unmarshalling the XML : %w", err)
	}
	unescapeRSSFeed(&feed)
	return &feed, nil
}

func unescapeRSSFeed(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
	}
}
