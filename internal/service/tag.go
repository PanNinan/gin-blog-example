package service

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State int    `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `json:"name" binding:"required,min=2,max=100"`
	CreatedBy string `json:"created_by" binding:"required,min=3,max=100"`
	State     int    `json:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	Name       string `json:"name" binding:"min=2,max=100"`
	State      int    `json:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `json:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID int `form:"id" binding:"required,gte=1"`
}
