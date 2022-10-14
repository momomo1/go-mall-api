package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-mall-api/api/admin/v1/entity"
	"go-mall-api/pkg/helpers"
	"go-mall-api/pkg/response"
	"strconv"
)

//营销-品牌推荐

func HomeBrandListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeBrandListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.BrandName = query["brandName"]
		request.RecommendStatus = query["recommendStatus"]

		out, _ := c.HomeBrandList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func HomeBrandUpdateRecommendStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeBrandUpdateRecommendStatusRequest
		request.Ids = ctx.PostForm("ids")
		request.RecommendStatus = ctx.PostForm("recommendStatus")

		err := c.HomeBrandUpdateRecommendStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeBrandCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request []entity.HomeBrandCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.HomeBrandCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeBrandUpdateSortHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeBrandUpdateSortRequest
		query := helpers.GetQueryParams(ctx)
		request.Id = query["id"]
		request.Sort = query["sort"]

		err := c.HomeBrandUpdateSort(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeBrandDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeBrandDeleteRequest
		request.Ids = ctx.PostForm("ids")

		err := c.HomeBrandDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

//营销-营销专题推荐

func HomeRecommendSubjectListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendSubjectListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.SubjectName = query["subjectName"]
		request.RecommendStatus = query["recommendStatus"]

		out, _ := c.HomeRecommendSubjectList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func HomeRecommendSubjectUpdateRecommendStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendSubjectUpdateRecommendStatusRequest
		request.Ids = ctx.PostForm("ids")
		request.RecommendStatus = ctx.PostForm("recommendStatus")

		err := c.HomeRecommendSubjectUpdateRecommendStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeRecommendSubjectCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request []entity.HomeRecommendSubjectCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)

		err := c.HomeRecommendSubjectCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeRecommendSubjectUpdateSortHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendSubjectUpdateSortRequest
		query := helpers.GetQueryParams(ctx)
		request.Id = query["id"]
		request.Sort = query["sort"]

		err := c.HomeRecommendSubjectUpdateSort(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeRecommendSubjectDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendSubjectDeleteRequest
		request.Ids = ctx.PostForm("ids")

		err := c.HomeRecommendSubjectDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

//营销-营销广告列表

func HomeAdvertiseHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeAdvertiseRequest
		request.Id = ctx.Param("id")

		out, _ := c.HomeAdvertise(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func HomeAdvertiseListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeAdvertiseListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.Type = query["type"]
		request.Name = query["name"]
		request.EndTime = query["endTime"]

		out, _ := c.HomeAdvertiseList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func HomeAdvertiseUpdateStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeAdvertiseUpdateStatusRequest
		query := helpers.GetQueryParams(ctx)
		request.Id = ctx.Param("id")
		request.Status = query["status"]

		err := c.HomeAdvertiseUpdateStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeAdvertiseCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeAdvertiseCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.HomeAdvertiseCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeAdvertiseUpdateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeAdvertiseUpdateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		request.Id = ctx.Param("id")
		err := c.HomeAdvertiseUpdate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeAdvertiseDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeAdvertiseDeleteRequest
		request.Ids = ctx.PostForm("ids")

		err := c.HomeAdvertiseDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

//营销-营销新品推荐

func HomeNewProductListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeNewProductListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.RecommendStatus = query["recommendStatus"]
		request.ProductName = query["productName"]

		out, _ := c.HomeNewProductList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func HomeNewProductUpdateRecommendStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeNewProductUpdateRecommendStatusRequest
		request.Ids = ctx.PostForm("ids")
		request.RecommendStatus = ctx.PostForm("recommendStatus")

		err := c.HomeNewProductUpdateRecommendStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeNewProductCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request []entity.HomeNewProductCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.HomeNewProductCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeNewProductUpdateSortHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeNewProductUpdateSortRequest
		query := helpers.GetQueryParams(ctx)
		request.Id = query["id"]
		request.Sort = query["sort"]

		err := c.HomeNewProductUpdateSort(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeNewProductDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeNewProductDeleteRequest
		request.Ids = ctx.PostForm("ids")
		err := c.HomeNewProductDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

//营销-营销人气推荐

func HomeRecommendProductListHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendProductListRequest
		query := helpers.GetQueryParams(ctx)
		pageSize, _ := strconv.Atoi(query["pageSize"])
		request.PageSize = pageSize
		request.RecommendStatus = query["recommendStatus"]
		request.ProductName = query["productName"]

		out, _ := c.HomeRecommendProductList(ctx, &request)
		response.OkWithData(ctx, out)
	}
}

func HomeRecommendProductUpdateRecommendStatusHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendProductUpdateRecommendStatusRequest
		request.Ids = ctx.PostForm("ids")
		request.RecommendStatus = ctx.PostForm("recommendStatus")

		err := c.HomeRecommendProductUpdateRecommendStatus(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeRecommendProductCreateHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request []entity.HomeRecommendProductCreateRequest
		j := helpers.GetRequestPayload(ctx)
		json.Unmarshal(j, &request)
		err := c.HomeRecommendProductCreate(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeRecommendProductUpdateSortHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendProductUpdateSortRequest
		query := helpers.GetQueryParams(ctx)
		request.Id = query["id"]
		request.Sort = query["sort"]

		err := c.HomeRecommendProductUpdateSort(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}

func HomeRecommendProductDeleteHttpHandler(c AdminController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request entity.HomeRecommendProductDeleteRequest
		request.Ids = ctx.PostForm("ids")
		err := c.HomeRecommendProductDelete(ctx, &request)
		if err != nil {
			response.FailWithMessage(ctx, err.Error())
			return
		}
		response.Ok(ctx)
	}
}
