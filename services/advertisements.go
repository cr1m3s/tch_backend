package services

import (
	"fmt"
	"time"

	"github.com/cr1m3s/tch_backend/di"
	"github.com/cr1m3s/tch_backend/models"
	"github.com/cr1m3s/tch_backend/queries"
	"github.com/cr1m3s/tch_backend/repositories"
	"github.com/gin-gonic/gin"
)

type AdvertisementService struct {
	db repositories.AdvertisementsRepositoryInterface
}

func NewAdvertisementService() *AdvertisementService {
	return &AdvertisementService{
		db: di.NewAdvertisementsRepository(),
	}
}

func (t *AdvertisementService) AdvCreate(ctx *gin.Context, inputModel models.AdvertisementInput, userID int64) (queries.Advertisement, error) {

	user, err := t.db.GetUserById(ctx, userID)
	if err != nil {
		return queries.Advertisement{}, err
	}

	args := &queries.CreateAdvertisementParams{
		Title:       inputModel.Title,
		Provider:    user.Name,
		ProviderID:  user.ID,
		Attachment:  inputModel.Attachment,
		Experience:  inputModel.Experience,
		Category:    inputModel.Category,
		Time:        inputModel.Time,
		Price:       inputModel.Price,
		Format:      inputModel.Format,
		Language:    inputModel.Language,
		Description: inputModel.Description,
		MobilePhone: inputModel.MobilePhone,
		Email:       user.Email,
		Telegram:    inputModel.Telegram,
		CreatedAt:   time.Now(),
	}

	advertisement, err := t.db.CreateAdvertisement(ctx, *args)
	if err != nil {
		return queries.Advertisement{}, err
	}

	return advertisement, nil
}

func (t *AdvertisementService) AdvPatch(ctx *gin.Context, patch models.AdvertisementUpdate) (queries.Advertisement, error) {
	advertisement, err := t.db.GetAdvertisementByID(ctx, patch.ID)

	if err != nil {
		return queries.Advertisement{}, err
	}

	advertisementTmp := &queries.UpdateAdvertisementParams{
		ID:          advertisement.ID,
		Title:       advertisement.Title,
		CreatedAt:   time.Now(),
		Attachment:  advertisement.Attachment,
		Experience:  advertisement.Experience,
		Category:    advertisement.Category,
		Time:        advertisement.Time,
		Price:       advertisement.Price,
		Format:      advertisement.Format,
		Language:    advertisement.Language,
		Description: advertisement.Description,
		MobilePhone: advertisement.MobilePhone,
		Telegram:    advertisement.Telegram,
	}

	result, err := t.db.UpdateAdvertisement(ctx, *advertisementTmp)

	if err != nil {
		return queries.Advertisement{}, err
	}

	return result, nil
}

func (t *AdvertisementService) AdvDelete(ctx *gin.Context, advId int64, userId int64) error {
	advertisement, err := t.db.GetAdvertisementByID(ctx, advId)

	if err != nil {
		return fmt.Errorf("Advertisement not found.")
	}

	if advertisement.ProviderID != userId {
		return fmt.Errorf("Don't have permissions to delete advertisement.")
	}

	err = t.db.DeleteAdvertisementByID(ctx, advId)
	if err != nil {
		return err
	}

	return nil
}
