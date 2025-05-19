package error

type Error struct {
	Code        int    `json:"errCode,omitempty"`
	Type        string `json:"errType,omitempty"`
	ShortError  string `json:"shortError,omitempty"`
	Description string `json:"description,omitempty"`
}

type apiError struct {
	DefaultError        *Error
	InternalServerError *Error
	JsonMarshalError    *Error
	JsonUnmarshalError  *Error
	Unauthorized        *Error
	BadRequest          *Error
	NoDataFound         *Error
	OracleDBConnError   *Error
	GetDBError          *Error
	AddDBError          *Error
	DelDBError          *Error
	GetCacheError       *Error
	SetCacheError       *Error
}

var ApiErrors = ErrorInit()

func ErrorInit() *apiError {
	return &apiError{
		DefaultError:        &Error{Code: 1000, Type: "DefaultError", ShortError: "Something went wrong"},
		InternalServerError: &Error{Code: 1001, Type: "InternalServerError", ShortError: "Internal server error"},
		JsonMarshalError:    &Error{Code: 1002, Type: "JsonMarshalError", ShortError: "Failed to marshal the data"},
		JsonUnmarshalError:  &Error{Code: 1003, Type: "JsonUnmarshalError", ShortError: "Failed to unmarshal the data"},
		Unauthorized:        &Error{Code: 1004, Type: "Unauthorized", ShortError: "Unauthorised request"},
		BadRequest:          &Error{Code: 1005, Type: "BadRequest", ShortError: "Missing or invalid request"},
		NoDataFound:         &Error{Code: 1006, Type: "NoDataFound", ShortError: "No data found in system"},
		OracleDBConnError:   &Error{Code: 1007, Type: "OracleDBConnError", ShortError: "Unable to connect Oracle DB"},
		GetDBError:          &Error{Code: 1008, Type: "GetDBError", ShortError: "Failed to get the data"},
		AddDBError:          &Error{Code: 1009, Type: "AddDBError", ShortError: "Failed to add data into table"},
		DelDBError:          &Error{Code: 1010, Type: "DelDBError", ShortError: "Failed to delete data from table"},
		GetCacheError:       &Error{Code: 1011, Type: "GetCacheError", ShortError: "Failed to get data from cache"},
		SetCacheError:       &Error{Code: 1012, Type: "SetCacheError", ShortError: "Failed to set data into cache"},
	}
}

func (e *Error) Error() string {
	if e == nil {
		return "UndefinedError"
	}
	return e.Type + ":-" + e.Description
}

func (e *Error) WithErrorDescription(errorDesc string) Error {
	var err Error
	if e == nil {
		err = *ApiErrors.DefaultError
	} else {
		err = *e
	}
	err.Description = errorDesc
	return err
}

// func (e *Error) GetErrorSlice(str string) (out []Error) {
// 	out = append(out, e.WithErrorDescription(str))
// 	return
// }
