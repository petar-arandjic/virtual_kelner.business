package grpc

//func transformErr(extendedError *error2.ExtendedError) *identitypb.ExtendedError {
//	if extendedError == nil {
//		return nil
//	}
//
//	var errors []*identitypb.Error
//	for _, value := range extendedError.Errors {
//		errors = append(errors, &identitypb.Error{
//			Key:     value.Key,
//			Message: value.Message,
//		})
//	}
//
//	return &identitypb.ExtendedError{
//		Key:     extendedError.Key,
//		Message: extendedError.Message,
//		Errors:  errors,
//	}
//}
