package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Ravinder2102/gator/internal/database"
	"github.com/google/uuid"
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
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// build the GET req
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	// do the req
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal fetched data
	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}

	return &rssFeed, nil
}

func scrapeFeeds(s *state) {
	// get next feed to fetch
	feed_to_fetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get next feed to fetch", err)
		return
	}

	// mark feed as fetched
	_, err = s.db.MarkFeedFetched(context.Background(), feed_to_fetch.ID)
	if err != nil {
		log.Printf("couldn't mark feed %s fetched: %v", feed_to_fetch.Name, err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed_to_fetch.Url)
	if err != nil {
		log.Printf("couldn't fetch feed %s: %v", feed_to_fetch.Name, err)
	}

	// save feed to posts
	for _, item := range rssFeed.Channel.Item {
		pubAt := sql.NullTime{}
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err == nil {
			pubAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: pubAt,
			FeedID:      feed_to_fetch.ID,
		}
		_, err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("couldn't save post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed_to_fetch.Name, len(rssFeed.Channel.Item))
}
