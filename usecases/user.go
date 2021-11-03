package usecases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"spos/users/constant"
	"spos/users/models/views"

	"github.com/s-pos/go-utils/utils/response"
)

func (u *usecase) Profile(ctx context.Context, userId int) response.Output {
	user, err := u.repo.FindUserById(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.Errors(ctx, http.StatusNotFound, string(constant.UserNotFound), constant.Message[constant.FailedGetProfile], constant.Reason[constant.UserNotFound], err)
		}

		return response.Errors(ctx, http.StatusInternalServerError, string(constant.ErrorQuery), constant.Message[constant.FailedGetProfile], constant.ErrorGlobal, err)
	}

	userStores, err := u.repo.FindUserStoreByUserId(user.GetId())
	if err != nil {
		return response.Errors(ctx, http.StatusInternalServerError, string(constant.ErrorQuery), constant.Message[constant.FailedGetProfile], constant.ErrorGlobal, err)
	}
	user.SetUserStores(userStores)

	var viewUserStores = make([]*views.UserStore, 0)
	for _, userStore := range user.GetUserStores() {
		store, err := u.repo.FindStoreById(userStore.GetStoreID())
		if err != nil {
			log.Println(err, userStore.GetStoreID())
			continue
		}

		logo := fmt.Sprintf("%s/%s/store/%s", os.Getenv("URL_STORAGE"), os.Getenv("GCS_BUCKET_NAME"), store.GetLogo())
		store.SetLogo(logo)
		store.CreatedAt = store.GetCreatedAt()
		store.SetFormattedCreatedAt(store.GetCreatedAt())

		vStore := &views.UserStore{
			ID:      userStore.GetID(),
			Enabled: userStore.GetEnabled(),
			Store:   store,
		}

		viewUserStores = append(viewUserStores, vStore)
	}

	userView := views.ResponseProfile{
		ID:              user.GetId(),
		Name:            user.GetName(),
		Email:           user.GetEmail(),
		Phone:           user.GetPhone(),
		IsEmailVerified: user.IsEmailVerified(),
		Stores:          viewUserStores,
	}

	return response.Success(ctx, http.StatusOK, string(constant.SuccessGetProfile), constant.Message[constant.SuccessGetProfile], userView)
}
