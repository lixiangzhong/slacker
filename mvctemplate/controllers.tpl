//Generated by slacker 
//gin router
/*
api.GET("/{{.LowerName}}/:id",controllers.Take{{.CamelCaseName}})
api.GET("/{{.LowerName}}",controllers.List{{.CamelCaseName}})
api.POST("/{{.LowerName}}",controllers.Create{{.CamelCaseName}})
api.PUT("/{{.LowerName}}/:id",controllers.Update{{.CamelCaseName}})
api.PATCH("/{{.LowerName}}/:id",controllers.Patch){{.CamelCaseName}}
api.DELETE("/{{.LowerName}}/:id",controllers.Delete{{.CamelCaseName}}) 
api.PATCH("/{{.LowerName}}",controllers.BatchPatch{{.CamelCaseName}})
api.DELETE("/{{.LowerName}}",controllers.BatchDelete{{.CamelCaseName}})
api.PUT("/{{.LowerName}}",controllers.BatchUpdate{{.CamelCaseName}}) 
*/

package controllers

import(
	 "{{"gosrc/models"| .ImportLibrary}}/{{.Name}}"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
 

func   List{{.CamelCaseName}}(c *gin.Context)  {
    var {{.LowerName}} {{.LowerName}}.{{.CamelCaseName}}
	offset,_:=strconv.ParseUint(c.DefaultQuery("offset","0"), 10, 64)
 
    limit,_:=strconv.ParseUint(c.DefaultQuery("limit","0"),10,64)
 	c.ShouldBindQuery(&{{.LowerName}})
   data,total,err:= Service.List{{.CamelCaseName}}(offset,limit,{{.LowerName}})
			JSON(c, gin.H{
		"data":  data,
		"total": total,
	},err)
}

func   Take{{.CamelCaseName}}(c *gin.Context)  {
    var {{.LowerName}} {{.LowerName}}.{{.CamelCaseName}}
		c.ShouldBindQuery(&{{.LowerName}})
    id, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		JSON(c, nil,errcode.New(errcode.BadRequest,err))
		return
	} 
    {{.LowerName}}, err = Service.Take{{.CamelCaseName}}(id) 
			JSON(c, {{.LowerName}},err)
}

func  Create{{.CamelCaseName}}(c *gin.Context)  {
    var {{.LowerName}} {{.LowerName}}.{{.CamelCaseName}}
    if err := c.ShouldBindJSON(&{{.LowerName}}); err != nil {
      		JSON(c, nil,errcode.New(errcode.BadRequest,err))
        return
	}
    {{.LowerName}},err := Service.Create{{.CamelCaseName}}({{.LowerName}})
		JSON(c, {{.LowerName}},err)
}

func Update{{.CamelCaseName}}(c *gin.Context)  {
	  id, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
	JSON(c, nil,errcode.New(errcode.BadRequest,err))
		return
	}
    var {{.LowerName}} {{.LowerName}}.{{.CamelCaseName}}
	if err := c.ShouldBindJSON(&{{.LowerName}}); err != nil {
		JSON(c, nil,errcode.New(errcode.BadRequest,err))
		return
	}
	{{.LowerName}}.{{.PrimaryKeyColumn.CamelCaseName}} = id
	err = Service.Update{{.CamelCaseName}}({{.LowerName}})
		JSON(c, nil,err)
}

func  Patch{{.CamelCaseName}}(c *gin.Context)  {
	id, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
	JSON(c, nil,errcode.New(errcode.BadRequest,err))
		return
	}
	var updatefields = make(map[string]interface{})
	if err := c.ShouldBindJSON(&updatefields); err != nil {
	JSON(c, nil,errcode.New(errcode.BadRequest,err))
		return
	} 
	err = Service.Patch{{.CamelCaseName}}(id,updatefields)
	JSON(c, nil,err)
}

func  Delete{{.CamelCaseName}}(c *gin.Context)  {
    id, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil { 
		JSON(c, nil,errcode.New(errcode.BadRequest,err))
		return
	} 
	err = Service.Delete{{.CamelCaseName}}(id)
		JSON(c, nil,err)
}

/*
func  BatchDelete{{.CamelCaseName}}(c *gin.Context) {
	ids := strings2int64s(c.QueryArray("id"))
	ids = append(ids, strings2int64s(c.QueryArray("id[]"))...) 
	err := Service.BatchDelete{{.CamelCaseName}}(ids)
	JSON(c, nil,err)
}



func  BatchPatch{{.CamelCaseName}}(c *gin.Context) {
	var updatefields = make(map[string]interface{})
	if err := c.ShouldBindJSON(&updatefields); err != nil {
			JSON(c, nil,errcode.New(errcode.BadRequest,err))
		return
	}
	ids := strings2int64s(c.QueryArray("id"))
	ids = append(ids, strings2int64s(c.QueryArray("id[]"))...) 
	err := Service.BatchPatch{{.CamelCaseName}}(ids, updatefields)
		JSON(c, nil,err)
}

func  BatchUpdate{{.CamelCaseName}}(c *gin.Context) {
	var {{.LowerName}} {{.LowerName}}.{{.CamelCaseName}} 
	if err := c.ShouldBindJSON(&{{.LowerName}}); err != nil {
		JSON(c, nil,err)
		return
	}
	ids := strings2int64s(c.QueryArray("id"))
	ids = append(ids, strings2int64s(c.QueryArray("id[]"))...)
	err := Service.BatchUpdate{{.CamelCaseName}}(ids) 
		JSON(c, nil,err)
} 
*/