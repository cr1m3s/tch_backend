package services

import(
	"time"
	"reflect"
	
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/models"
	"github.com/gin-gonic/gin"
)

func (t *UserService) AdvCreate(ctx *gin.Context, inputModel models.AdvertisementInput, userID int64) (queries.Advertisement, error) {
	
	user, err := t.db.GetUserById(ctx, userID)
	if err != nil {
		return queries.Advertisement{}, err
	}
	
	args := &queries.CreateAdvertisementParams{
		Title: inputModel.Title,
		Provider: user.Name,
		ProviderID: user.ID,
		Attachment: inputModel.Attachment,
		Experience: inputModel.Experience,
		Category: inputModel.Category,
		Time: inputModel.Time,
		Price: inputModel.Price,
		Format: inputModel.Format,
		Language: inputModel.Language,
		Description: inputModel.Description,
		MobilePhone: inputModel.MobilePhone,
		Email: user.Email,
		Telegram: inputModel.Telegram,
		CreatedAt: time.Now(),
	}

	advertisement, err := t.db.CreateAdvertisement(ctx, *args)
	if err != nil {
		return queries.Advertisement{}, err
	}

	return advertisement, nil
}

func (t *UserService) AdvPatch(ctx *gin.Context, patch models.AdvertisementUpdate) (queries.Advertisement, error) {
	advertisement, err := t.db.GetAdvertisementByID(ctx, patch.ID)

	if err != nil {
		return queries.Advertisement{}, err
	}

	advertisementTmp := &queries.UpdateAdvertisementParams{
		ID:	advertisement.ID, 
  		Title: advertisement.Title,
  		CreatedAt: time.Now(),
  		Attachment: advertisement.Attachment,
  		Experience: advertisement.Experience,
  		Category: advertisement.Category,
  		Time: advertisement.Time,
  		Price: advertisement.Price,
  		Format: advertisement.Format,
  		Language: advertisement.Language,
  		Description: advertisement.Description,
  		MobilePhone: advertisement.MobilePhone,
  		Telegram: advertisement.Telegram,
	}

	// this BS needed to transfer data only from fields with values
	// both stucts should have same number of fields
	advValue := reflect.ValueOf(advertisementTmp).Elem()
	patchValue := reflect.ValueOf(patch)	
	for i := 0; i < advValue.NumField(); i++ {
		field := advValue.Field(i)
		updateField := patchValue.Field(i)

		if updateField.IsValid() && !updateField.IsZero() {
			field.Set(updateField)
		}
	}

	result, err := t.db.UpdateAdvertisement(ctx, *advertisementTmp)

	if err != nil {
		return queries.Advertisement{}, err
	}

	return result, nil
}