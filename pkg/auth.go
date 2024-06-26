package pkg

import (
	"errors"
)

// GetUserID checks if the token is valid by calling auth.getMe method and returns the user id
func GetUserID(token string) (string, error) {
	request, err := ApiRequest(token, "auth.getMe", nil)
	if err != nil {
		return "", err
	}
	if request.Error.Code != 0 {
		currentLogger("Failed to get user: %s", request.Error.Message)
		return "", errors.New(request.Error.Message)
	}

	user, err := MapToStruct[UserInfo](request.Result)
	if err != nil {
		return "", err
	}

	if user.ID != "" {
		return user.ID, nil
	}

	return "", errors.New("failed to get user id from response")
}
