package handler


func (h *Handler) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			// h.HTTPErrorHandler(errNotAuthorized, c)
			return errNotAuthorized
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			// h.HTTPErrorHandler(errNotAuthorized, c)
			return errNotAuthorized
		}
		tokenString := headerParts[1]
		userId, err := h.service.Auth.ParseToken(tokenString)
		if err != nil || userId == 0 {
			// h.HTTPErrorHandler(errIvalidJWTToken, c)
			return errIvalidJWTToken
		}
		c.Set("userId", userId)
		return next(c)
	}
}