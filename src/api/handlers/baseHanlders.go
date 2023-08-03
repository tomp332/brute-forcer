package handlers

//func CrudGetWithPagination(c echo.Context) error {
//	paginationParams := new(src.Paginate)
//	err := c.Bind(paginationParams)
//	if err != nil || paginationParams == nil {
//		log.Printf("Error binding pagination struct")
//		return c.JSONBlob(http.StatusBadRequest,
//			utils.BadRequestError("One or more of the parameters specified for pagination was incorrect", err))
//	}
//	creds, err := crud.GetCreds(paginationParams.Limit, paginationParams.Page)
//	if err != nil {
//		log.Printf("Error getting creds: %s", err.Error())
//		return c.JSONBlob(http.StatusInternalServerError,
//			utils.BadRequestError("Error fetching credentials from the database", err))
//	}
//	return c.JSON(http.StatusOK, creds)
//}
