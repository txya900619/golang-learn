package vm

import (
	"../model"
)

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}
type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	u1 := model.User{Username: "wayne"}
	u2 := model.User{Username: "lolicon"}

	posts := []model.Post{
		model.Post{User: u1, Body: "naniiiiii"},
		model.Post{User: u2, Body: "bagana"},
	}
	v := IndexViewModel{BaseViewModel{Title: "home"}, u1, posts}
	return v

}
