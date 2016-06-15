package api


import (
	"reflect"
	"errors"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"net/http"
)

// Use omitFields to specify fields that should not be validated as required in the incoming DTO.
func decodeAndValidateRequest(r http.Request, dtoStruct interface{}, omitFields map[string]bool) (map[string]interface{}, error){
	if omitFields == nil{
		omitFields = map[string]bool{}
	}
	// Save the buffer so we can decode it a couple of times.
	buf, _ := ioutil.ReadAll(r.Body)
	var err error

	var requestVals = make(map[string]interface{})

	// Try and deserialize the body into the dto type.  If this works it means
	// we have valid JSON and something like a valid dto.
	buffer1 := bytes.NewBuffer(buf)
	paramDecoder := json.NewDecoder(buffer1)
	err = paramDecoder.Decode(&requestVals)
	if err != nil {return requestVals, err}

	//  Now we deserialize the body into a map and check for required and extra properties.
	io := bytes.NewBuffer(buf)
	decoder := json.NewDecoder(io)
	var int map[string]interface{}
	err = decoder.Decode(&int)

	err = validateFields(int, dtoStruct, omitFields)

	return requestVals, err
}

func validateFields(requestMap map[string]interface{}, dto interface{}, omitFields map[string]bool) (error){
	// Create the fields map from the dto.
	var dtoFields = make(map[string]bool)
	dtoVal := reflect.ValueOf(dto)
	for i := 0; i < dtoVal.NumField(); i++ {
		dtoFields[dtoVal.Type().Field(i).Name] = true
	}

	// Loop the request map and check to make sure all fields in the request are present in the dto
	for fieldName, _ := range requestMap{
		_, ok := dtoFields[fieldName]
		if !ok{
			return errors.New(fieldName + " is not a valid field.")
		}
	}

	// Loop the dto fields and make sure all fields in the dto are present in the request.
	for fieldName, _ := range dtoFields{
		_, ok := requestMap[fieldName]
		if !ok && !omitFields[fieldName] {
			return errors.New(fieldName + " is required.")
		}
	}

	return nil
}

func checkReadonlyFields(existing interface{}, new interface{}) (error) {
	dtoType := reflect.TypeOf(new)
	updatedValue := reflect.ValueOf(new)
	existingValue := reflect.ValueOf(existing)

	for i := 0; i < updatedValue.NumField(); i++ {
		field := dtoType.Field(i)
		if field.Tag.Get("ark-readonly") == "true" {
			newVal := updatedValue.Field(i).Interface()
			oldVal := existingValue.Field(i).Interface()

			if newVal != oldVal {
				return errors.New(field.Name + " is read only.")
			}
		}
	}
	return nil
}