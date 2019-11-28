package mapstructure

type ConvertMapStruct interface{
	Convert() (interface{}, error)
}
