package newsfeed

import "testing"

func TestAdd(t *testing.T) {
  feed := New()
  feed.Add(Item{})

  if len(feed.Items) != 1 {
    f.Errorf("Item was not added")
  }
}

func TestGetAll(t *testing.T) {
  feed := New()
  feed.Add(Item{})

  results := feed.GetAll()

  if len(feed.results) != 1 {
    f.Errorf("Item was not added")
  }
}
