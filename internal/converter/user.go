package converter

import "vetback/internal/model"

func UserToStorage(user model.UserInfo, hash string) model.UserToStorage {
	return model.UserToStorage{
		Hash:           hash,
		FullName:       user.FullName,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		Role:           user.Role,
		Specialization: user.Specialization,
	}
}

func UserInfoToSession(user model.UserInfo, id int) model.Claims {
	return model.Claims{
		UserId:         id,
		FullName:       user.FullName,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		Role:           user.Role,
		Specialization: user.Specialization,
	}
}

func UserToSession(user model.User) model.Claims {
	return model.Claims{
		UserId:         user.UserId,
		FullName:       user.FullName,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		Role:           user.Role,
		Specialization: user.Specialization,
	}
}

func UserToSafe(user model.User) model.SafeUser {
	return model.SafeUser{
		UserId:         user.UserId,
		FullName:       user.FullName,
		Role:           user.Role,
		Specialization: user.Specialization,
	}
}

func UserToUpdate(user model.UserToUpdate, hash string) model.UpdatedUserToStorage {
	return model.UpdatedUserToStorage{
		Hash:           hash,
		FullName:       user.FullName,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		Specialization: user.Specialization,
	}
}

func UpdatedUserToClaims(user model.UserToUpdate, id int) model.Claims {
	var role string
	if user.Specialization != "" {
		role = "doctor"
	} else {
		role = "owner"
	}
	return model.Claims{
		UserId:         id,
		FullName:       user.FullName,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		Role:           role,
		Specialization: user.Specialization,
	}
}
