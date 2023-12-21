package apiman

import (
	"github.com/gin-gonic/gin"
	"github.com/philiphil/apiman/orm/entity"

	"github.com/philiphil/apiman/router"
)

func (r *ApiRouter[T]) Patch(c *gin.Context) {
	id := c.Param("id")
	obj, err := r.Orm.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(404, "Resource not found")
		return
	}
	if !router.UnserializeBodyAndMerge(c, obj) {
		return
	}
	var cast entity.IEntity
	cast = *obj
	cast = cast.SetId(id)
	convertedEntity, _ := cast.(T)
	err = r.Orm.Update(&convertedEntity)
	if err != nil {
		c.AbortWithStatusJSON(500, "Database issue")
		return
	}

	c.JSON(204, nil)
}