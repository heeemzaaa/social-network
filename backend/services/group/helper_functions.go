package group

import (
	"errors"

	"social-network/backend/utils"
)

func (service *GroupService) ValidateGroupTitle(title string) error {
	_, has_title, _ := service.gRepo.GetItem("groups", "title", title)

	if has_title {
		return errors.New("title already in use")
	}

	if err := utils.ValidateTitle(title); err != nil {
		return err
	}
	return nil
}
