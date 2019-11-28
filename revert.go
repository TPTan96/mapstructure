package mapstructure

type RevertMapStruct interface{
	RevertMapStruct(interface{})  error
}