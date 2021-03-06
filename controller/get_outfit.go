package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
)

// GetOutfit renders an outfit page
func GetOutfit(request *http.Request, userID string) {
	outfitID := request.ParamByName("id")

	outfit, err := db.SelectOutfitByID(outfitID, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting outfit by ID failed"))
		return
	}

	outfitItems, err := db.SelectOutfitItemsByOutfitID(outfitID, userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting outfit items by outfit ID failed"))
		return
	}

	outfit.OutfitItems = outfitItems

	request.Render("outfit", ui.NewOutfitPage(userID, *outfit))
}
