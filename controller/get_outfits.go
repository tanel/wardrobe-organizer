package controller

import (
	"github.com/juju/errors"
	"github.com/tanel/wardrobe-organizer/db"
	"github.com/tanel/wardrobe-organizer/ui"
	"github.com/tanel/webapp/http"
	commonui "github.com/tanel/webapp/ui"
)

// GetOutfits renders outfits page
func GetOutfits(request *http.Request, userID string) {
	outfits, err := db.SelectOutfitsByUserID(userID)
	if err != nil {
		request.InternalServerError(errors.Annotate(err, "selecting outfits by user ID failed"))
		return
	}

	request.Render("outfits", ui.OutfitsPage{
		Page: commonui.Page{
			UserID: userID,
		},
		Outfits: outfits,
	})
}
