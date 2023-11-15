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

func (t *AdvertisementService) AdvCreate(ctx *gin.Context, inputModel models.AdvertisementInput, userID int64, u UserService) (queries.Advertisement, error) {
	user, err := u.UserInfo(ctx, userID)
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

func (t *AdvertisementService) AdvPatch(ctx *gin.Context, patch models.AdvertisementUpdate, userID int64) (queries.Advertisement, error) {
	adv, err := t.db.GetAdvertisementByID(ctx, patch.ID)

	if err != nil {
		return queries.Advertisement{}, err
	}

	if adv.ProviderID != userID {
		return queries.Advertisement{}, fmt.Errorf("Unauthorized")
	}

	advertisementTmp := &queries.UpdateAdvertisementParams{
		ID:          patch.ID,
		Title:       patch.Title,
		CreatedAt:   time.Now(),
		Attachment:  patch.Attachment,
		Experience:  patch.Experience,
		Category:    patch.Category,
		Time:        patch.Time,
		Price:       patch.Price,
		Format:      patch.Format,
		Language:    patch.Language,
		Description: patch.Description,
		MobilePhone: patch.MobilePhone,
		Telegram:    patch.Telegram,
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

func (t *AdvertisementService) AdvGetAll(ctx *gin.Context) ([]queries.Advertisement, error) {

	advertisements, err := t.db.GetAdvertisementAll(ctx)

	if err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (t *AdvertisementService) AdvGetByID(ctx *gin.Context, id int64) (queries.Advertisement, error) {
	advertisement, err := t.db.GetAdvertisementByID(ctx, id)

	if err != nil {
		return queries.Advertisement{}, err
	}
	return advertisement, nil
}

func (t *AdvertisementService) AdvGetFiltered(ctx *gin.Context, filter models.AdvertisementFilter) ([]queries.FilterAdvertisementsRow, error) {

	argFilter := queries.FilterAdvertisementsParams{
		Orderby:      filter.Orderby,
		Sortorder:    filter.Sortorder,
		Offsetadv:    filter.Offsetadv,
		Limitadv:     filter.Limitadv,
		Advcategory:  filter.Category,
		Timelength:   filter.Timelength,
		Advformat:    filter.Format,
		Minexp:       filter.Minexp,
		Maxexp:       filter.Maxexp,
		Minprice:     filter.Minprice,
		Maxprice:     filter.Maxprice,
		Advlanguage:  filter.Language,
		Titlekeyword: filter.Titlekeyword,
	}
	advertisements, err := t.db.FilterAdvertisements(ctx, argFilter)
	if err != nil {
		return nil, err
	}

	return advertisements, nil
}

func (t *AdvertisementService) AdvGetMy(ctx *gin.Context, userID int64) ([]queries.Advertisement, error) {
	advertisements, err := t.db.GetAdvertisementMy(ctx, userID)

	if err != nil {
		return nil, err
	}

	return advertisements, nil
}
