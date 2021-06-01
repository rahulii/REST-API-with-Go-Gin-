package newsfeed

import "testing"


func TestAdd(t *testing.T){
	for _,test := range []struct {
		name string
		input Item
	}{
		{
			name: "add news item",
			input: Item{
				Title : "Covid",
				Post: "pandemic",
			},
		},
	}{
		t.Run(test.name,func(t *testing.T) {
			feed := New()
			feed.Add(test.input)
			if len(feed.Items) != 1 {
				t.Errorf("Item was not added")
			}
		})
	}
}


 func TestGetAll(t *testing.T){
	for _,test := range []struct {
		name string
		input Item
	}{
		{
			name : "Get All Items",
			input: Item{
				Title: "Covid",
				Post: "Pandemic",
			},
		},
	}{
		t.Run(test.name,func(t *testing.T) {
			feed := New()
			feed.Add(test.input)
			result := feed.GetAll()
			if len(result) != 1{
				t.Errorf("Expected 1 Item!")
			}
		})
	}
 }