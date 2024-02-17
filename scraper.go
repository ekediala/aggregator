package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/ekediala/aggregator/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequests time.Duration) {
	log.Printf("scraping on %v go routines every %s duration", concurrency, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for range ticker.C {

		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Printf("error fetching feeds %v", err)
			continue
		}

		wg := sync.WaitGroup{}

		for _, feed := range feeds {

			wg.Add(1)

			go scrapePost(&wg, feed, db)

		}

		wg.Wait()
	}

}

func scrapePost(wg *sync.WaitGroup, feed database.Feed, db *database.Queries) {
	defer wg.Done()

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Printf("error fetching rss feed %v", err)
		return
	}

	_, err = db.UpdateLastFetchedAt(context.Background(), feed.ID)

	if err != nil {
		log.Printf("error fetching marking feed as fetched %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Items {

		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			t = time.Now()
		}

		post, err := db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			FeedID:      feed.ID,
			PublishedAt: t,
		})

		if err != nil {
			log.Printf("error creating post %v", post)
			return
		}

		log.Printf("post created: %v", post)
	}

	log.Printf("feed %v collected. %d posts found", feed.Name, len(rssFeed.Channel.Items))
}
