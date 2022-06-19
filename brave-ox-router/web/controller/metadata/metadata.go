package metadata

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"brave-ox-web/web/response"
	"brave-ox-web/web/services/metadata"
)

type metadataController struct{ service metadata.MetadataService }

func MetadataController() *metadataController {
	return &metadataController{
		service: metadata.Service(),
	}
}

func (this *metadataController) GetCattleNFT(c *gin.Context) {
	pid := c.Param("id")

	id, err := strconv.ParseUint(pid, 0, 0)
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}

	cattle := this.service.Cattle(uint32(id))

	if cattle == nil {
		response.RespondWithRawJsonData(c, "NoSuchKey")
		return
	}

	response.RespondWithRawJsonData(c, cattle)
}

func (this *metadataController) CattleSkin(c *gin.Context) {
	pid := c.Param("id")

	id, err := strconv.ParseUint(pid, 0, 0)
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}

	skin := this.service.Skin(uint32(id))
	if skin == nil {
		response.RespondWithRawJsonData(c, "NoSuchKey")
		return
	}

	response.RespondWithRawJsonData(c, skin)
}

func (this *metadataController) Planet(c *gin.Context) {
	pid := c.Param("id")
	t, err := strconv.ParseUint(pid, 0, 0)
	if err != nil {
		response.RespondErr(c, err.Error())
		return
	}
	planet := this.service.Planet(uint32(t))
	if planet == nil {
		response.RespondWithRawJsonData(c, "NoSuchKey")
		return
	}
	response.RespondWithRawJsonData(c, planet)
}

type BoxURI struct {
	Types   uint32 `uri:"types"`
	Address string `uri:"address"`
}

func (this *metadataController) Box(c *gin.Context) {
	var req BoxURI
	if err := c.ShouldBindUri(&req); err != nil {
		response.RespondErr(c, err.Error())
		return
	}

	planet := this.service.Box(req.Types, req.Address)
	if planet == nil {
		response.RespondWithRawJsonData(c, "NoSuchKey")
		return
	}

	response.RespondWithRawJsonData(c, planet)
}
