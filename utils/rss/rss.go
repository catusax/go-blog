package rss

import (
	"blog/errors"
	"blog/models"
	"github.com/gorilla/feeds"
	"github.com/spf13/viper"
	"io/ioutil"
	"time"
)

func generateAtom() (string, error) {
	now := time.Now()
	host := viper.GetString("host")
	feed := &feeds.Feed{
		Title:       viper.GetString("siteName"),
		Link:        &feeds.Link{Href: host},
		Description: viper.GetString("description"),
		Author:      &feeds.Author{Name: viper.GetStringMapString("author")["name"], Email: viper.GetStringMapString("author")["email"]},
		Created:     now,
	}
	posts, err := models.GetAllPublishedPost()
	if err != nil {
		return "", errors.Errorf(err, "")
	}
	for _, post := range posts {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       post.Title,
			Link:        &feeds.Link{Href: host},
			Updated:     post.UpdatedAt,
			Description: post.Description,
			Content:     post.HTML,
		})
	}
	atom, err := feed.ToAtom()
	if err != nil {
		return "", errors.Errorf(err, "")
	}
	return atom, nil
}

func WriteAtom(filename string) error {
	var err error
	atom, err := generateAtom()
	if err != nil {
		return errors.Errorf(err, "generate rss failed")
	}
	err = ioutil.WriteFile(filename, []byte(atom), 0600)
	if err != nil {
		return errors.Errorf(err, "write rss failed")
	}
	return nil
}
