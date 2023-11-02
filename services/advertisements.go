package services

import(
	
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/gin-gonic/gin"
)

func (t *UserService) AdvCreate(ctx *gin.Context, inputModel queries.Advertisement) (queries.Advertisement, error) {
	args := &queries.CreateAdvertisementParams{
		Title: inputModel.Title,
		Provider: inputModel.Provider,
		Attachment: inputModel.Attachment,
		Experience: inputModel.Experience,
		Category: inputModel.Category,
		Time: inputModel.Time,
		Price: inputModel.Price,
		Format: inputModel.Format,
		Language: inputModel.Language,
		Description: inputModel.Description,
		MobilePhone: inputModel.MobilePhone,
		Email: inputModel.Email,
		Telegram: inputModel.Telegram,
	}

	advertisement, err := t.db.CreateAdvertisement(ctx, *args)
	if err != nil {
		return queries.Advertisement{}, err
	}

	return advertisement, nil
}