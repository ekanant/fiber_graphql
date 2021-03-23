package todo

type User struct {
	ID   string
	Name string
}

var userDatas = []*User{
	&User{
		ID:   "1",
		Name: "Moron",
	},
	&User{
		ID:   "2",
		Name: "Moron",
	},
}

func AddUser(user *User) int {
	idUserToCheck := user.ID

	userNotExist := true
	for _, u := range userDatas {
		if u.ID == idUserToCheck {
			userNotExist = false
			break
		}
	}

	if userNotExist {
		userDatas = append(userDatas, user)
	}
	return 1
}
