package apiman

import (
	"github.com/gin-gonic/gin"
	"github.com/philiphil/apiman/orm/entity"
	"github.com/philiphil/apiman/router"
)

func (r *ApiRouter[T]) Put(c *gin.Context) {
	id := c.Param("id")
	obj, err := r.Orm.GetByID(id)
	if err != nil {
		bfr := r.Orm.NewEntity()
		obj = &bfr
	}

	if err = r.WritingCheck(c, obj); err != nil {
		c.AbortWithStatusJSON(err.(ApiError).Code, err.(ApiError).Message)
		return
	}

	if err = router.UnserializeBodyAndMerge(c, obj); err != nil {
		c.AbortWithStatusJSON(err.(ApiError).Code, err.(ApiError).Message)
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
