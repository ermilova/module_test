package uobject

type UObject struct {
	properties map[string]interface{}
}

func NewUObject() *UObject {
	return &UObject{
		properties: make(map[string]interface{}),
	}
}

func (u *UObject) GetProperty(property string) interface{} {
	return u.properties[property]
}

func (u *UObject) SetProperty(property string, value interface{}) {
	u.properties[property] = value
}
