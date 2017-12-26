package controller

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tanel/wardrobe-manager-app/db"
	"github.com/tanel/wardrobe-manager-app/ui"
)

// GetItem renders an item page
func GetItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params, userID string) {
	item, err := db.SelectItemWithImagesByID(ps.ByName("id"), userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	page := ui.NewItemPage(userID, *item)
	if err := Render(w, "item", page); err != nil {
		log.Println(err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
