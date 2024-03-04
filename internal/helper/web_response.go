package helper

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    any         `json:"meta"`
}

func WebResponseOK(ctx *fiber.Ctx, response interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(WebResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    response,
	})
}

func WebResponseOKWithMeta(ctx *fiber.Ctx, response interface{}, meta any) error {
	return ctx.Status(fiber.StatusOK).JSON(WebResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    response,
		Meta:    meta,
	})
}

func WebResponseError(ctx *fiber.Ctx, statusCode int, err string) error {
	return ctx.Status(statusCode).JSON(WebResponse{
		Code:    statusCode,
		Message: err,
	})
}

func GetWebResponse(ctx *fiber.Ctx, result interface{}, err error) error {
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(WebResponse{
			Code:    fiber.StatusExpectationFailed,
			Message: err.Error(),
		})
	} else {
		return WebResponseOK(ctx, result)
	}
}

func GetWebResponseWithMeta(ctx *fiber.Ctx, result any, meta any, err error) error {
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(WebResponse{
			Code:    fiber.StatusExpectationFailed,
			Message: err.Error(),
		})
	} else {
		return WebResponseOKWithMeta(ctx, result, meta)
	}
}

func GetSteamResponse(ctx *fiber.Ctx, result *bytes.Buffer, fileType, fileName string, err error) error {
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(WebResponse{
			Code:    fiber.StatusExpectationFailed,
			Message: err.Error(),
		})
	} else {
		if fileType == "excel" {
			ctx.Response().Header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", fileName))
			ctx.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

			return ctx.Send(result.Bytes())
		} else {
			return ctx.Status(fiber.StatusOK).JSON(WebResponse{
				Code:    fiber.StatusExpectationFailed,
				Message: "unknown file type",
			})
		}
	}
}
