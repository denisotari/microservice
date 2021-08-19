package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
	"gitlab.com/faemproject/backend/faem/pkg/logs"
	"gitlab.com/faemproject/backend/faem/services/bootstrap/proto"
)

const (
	errBind = "failed to parse the request"
)

func (r *Rest) Hello(c echo.Context) error {
	var in proto.User
	if err := c.Bind(&in); err != nil {
		logs.LoggerForContext(c.Request().Context()).
			WithFields(logrus.Fields{
				"event":  "hello",
				"in":     in,
				"reason": errBind,
			}).
			Error(err)
		res := logs.OutputRestError(errBind, err)
		return c.JSON(http.StatusBadRequest, res)
	}

	out, err := r.Handler.Hello(c.Request().Context(), &in)
	if err != nil {
		logs.LoggerForContext(c.Request().Context()).
			Error(err) // you may add additional fields here
		res := logs.OutputRestError("can't greet", err)
		return c.JSON(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, out)
}
