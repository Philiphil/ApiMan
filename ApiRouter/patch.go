package ApiRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/philiphil/apiman/Gin"
	"github.com/philiphil/apiman/Gorm"
)

func (r *ApiRouter[T]) Patch(c *gin.Context) {
	id := c.Param("id")
	entity, err := r.Orm.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(404, "Resource not found")
		return
	}
	fmt.Println(entity)
	if !Gin.UnserializeBodyAndMerge(c, entity) {
		return
	}
	fmt.Println(entity)
	var cast Gorm.IEntity
	cast = *entity
	cast = cast.SetId(id)
	convertedEntity, _ := cast.(T)
	err = r.Orm.Update(&convertedEntity)
	if err != nil {
		c.AbortWithStatusJSON(500, "Database issue")
		return
	}

	c.JSON(204, nil)
}
