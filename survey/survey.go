package survey

import (
	"fmt"
	"go-search-history/dbutil"
	"go-search-history/inspect"
)

func Search(){
	fmt.Println("Search:")
	var search string
	fmt.Scanln(&search)
	var inquiry = inspect.Inspect(search)
	saveSearch(inquiry)
}

func saveSearch(inquiry inspect.Inquiry) {
	dbutil.ConnectMongoDB()

	defer dbutil.CloseMongoDB()
}