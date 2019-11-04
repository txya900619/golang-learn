package vm

type LoginViewModel struct {
	BaseViewModel
}
type LoginViewModelOp struct{}

func (LoginViewModelOp) GetMV() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}
