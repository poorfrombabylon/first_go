package inmemory

import (
	"errors"
	"test_task/service"
)

type InMemory struct {
	UrlsMap map[string]string
}

func NewInMemory() InMemory {
	UrlsMap := make(map[string]string)
	return InMemory{
		UrlsMap: UrlsMap,
	}
}

func (db InMemory) GetUrl(shorturl string) (string, error) {
	if shorturl == "" {
		return "", errors.New("invalid input")
	}

	longUrl, ok := db.UrlsMap[shorturl]

	if !ok {
		return "", errors.New("no such url")
	} else {
		return longUrl, nil
	}

	return "", nil
}

func (db InMemory) GetShortUrl(longurl string) (string, error) {
	if longurl == "" {
		return "", errors.New("invalid input")
	}

	for k, v := range db.UrlsMap {
		if longurl == k {
			return v, nil
		}
	}
	return "", nil
}

func (db InMemory) PostUrl(url string) (string, error) {
	if url == "" {
		return "", errors.New("invalid input")
	}
	_, ok := db.UrlsMap[url]

	if ok {
		return "", errors.New("this url is already in db")
	}

	shorturl := service.Shorten(url)
	db.UrlsMap[url] = shorturl
	return shorturl, nil
}
