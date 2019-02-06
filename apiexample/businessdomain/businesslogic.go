package businessdomain

// ResponseData contains business logic response data
type ResponseData struct {
	Data string
}

// BusinessLogic runs domain business logic
func BusinessLogic() (resp *ResponseData, err error) {
	// ... do something ...
	resp = &ResponseData{Data: "business logic result"}
	err = nil
	return
}
