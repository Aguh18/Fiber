package request

type PhotoRequest struct {
	CategotyId uint64 `json:"categoryid" form:"categoryId" validate:"required"`
}
