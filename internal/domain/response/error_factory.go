package response

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func NewErrorResponse(message string, code int) *ErrorResponse {
	return &ErrorResponse{
		Status:  "error",
		Message: message,
		Code:    code,
	}
}

func NewApiErrorResponse(c echo.Context, statusText string, message string, code int) error {
	return c.JSON(code, ErrorResponse{
		Status:  statusText,
		Message: message,
		Code:    code,
	})
}

func ToGraphqlErrorFromErrorResponse(errResp *ErrorResponse) error {
	if errResp == nil {
		return nil
	}

	if errResp.Message == "" {
		errResp.Message = "unknown error occurred"
	} // func ToGrpcErrorFromErrorResponse(err *ErrorResponse) error {
	// 	if err == nil {
	// 		return nil
	// 	}
	// 	return status.Errorf(codes.Code(err.Code),
	// 		errors.GrpcErrorToJson(&pb.ErrorResponse{
	// 			Status:  err.Status,
	// 			Message: err.Message,
	// 			Code:    int32(err.Code),
	// 		}),
	// 	)
	// }

	// func NewGrpcError(statusText string, message string, code int) error {
	// 	return status.Errorf(codes.Code(code),
	// 		errors.GrpcErrorToJson(&pb.ErrorResponse{
	// 			Status:  statusText,
	// 			Message: message,
	// 			Code:    int32(code),
	// 		}),
	// 	)
	// }

	return fmt.Errorf("graphql error: [%d] %s - %s", errResp.Code, errResp.Status, errResp.Message)
}

func NewGraphqlError(statusText string, message string, code int) error {
	errResp := &ErrorResponse{
		Status:  statusText,
		Message: message,
		Code:    code,
	}

	return fmt.Errorf("graphql error: [%d] %s - %s", errResp.Code, errResp.Status, errResp.Message)
}
