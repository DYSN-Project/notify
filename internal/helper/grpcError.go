package helper

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetGrpcUnauthenticatedError() error {
	st := status.New(codes.Unauthenticated, "not authenticated")

	return st.Err()
}

func IsGrpcInvalidData(err error) bool {
	return status.Code(err) == codes.InvalidArgument
}

func IsUnauthorizedData(err error) bool {
	return status.Code(err) == codes.Unauthenticated
}

func GetBadRequestError(err error) error {
	st := status.New(codes.InvalidArgument, err.Error())

	return st.Err()
}

func MakeGrpcBadRequestError(errDetails map[string]error) error {
	st := status.New(codes.InvalidArgument, "bad request error")
	rsp := getErrorResponse(errDetails)
	st, err := st.WithDetails(rsp)
	if err != nil {
		return err
	}

	return st.Err()
}

func MakeGrpcNotFoundError(errDetails error) error {
	st := status.New(codes.NotFound, "not found")
	rsp := getErrorResponse(map[string]error{"errors": errDetails})
	st, err := st.WithDetails(rsp)
	if err != nil {
		return err
	}

	return st.Err()
}

func getErrorResponse(errDetails map[string]error) *errdetails.BadRequest {
	rsp := &errdetails.BadRequest{
		FieldViolations: make([]*errdetails.BadRequest_FieldViolation, len(errDetails)),
	}

	i := 0
	for errKey, errItem := range errDetails {
		fv := &errdetails.BadRequest_FieldViolation{
			Field:       errKey,
			Description: errItem.Error(),
		}
		rsp.FieldViolations[i] = fv
		i++
	}

	return rsp
}
