package factories

import (
	"github.com/bxcodec/faker/v3"
	"go-mall-api/models/user"
	"go-mall-api/pkg/helpers"
)

func MakeUsers(times int) []user.User {
	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, model)
	}

	return objs
}
