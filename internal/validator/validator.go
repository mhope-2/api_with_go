package validator

type Validator struct{
	Errors map[string]string
}

func New() *Validator{
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string){
	if _, exists := v.Errors; !exists{
		v.Errors[key] = message
	}
}

func Check(ok bool, key, message string){
	if !ok{
		v.AddError(key, message)
	}
}