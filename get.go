package ApiMan

import (
	"github.com/gin-gonic/gin"
	"github.com/philiphil/apiman/router"
	"github.com/philiphil/apiman/serializer/format"
)

func (r *ApiRouter[T]) Get(c *gin.Context) {
	object, err := r.Orm.GetByID(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(404, "Resource not found")
		return
	}

	c.Render(200, router.SerializerRenderer{
		Data:   object,
		Format: format.JSON,
		Groups: []string{},
	})
}