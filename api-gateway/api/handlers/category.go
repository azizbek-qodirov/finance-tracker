package handlers

import (
	"net/http"
	"strconv"

	pb "gateway-service/genprotos" // Update with your actual package path

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateCategory godoc
// @Summary     Create Category
// @Description Creates a new category
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       category body     pb.CategoryCReqForSwagger true "Category creation request"
// @Success     200     {object} string "Category created successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     500     {object} string "Server error"
// @Router      /v1/category [post]
// @Security    BearerAuth
func (h *HTTPHandler) CreateCategory(c *gin.Context) {
	var body pb.CategoryCReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	claims, _ := c.Get("claims")
	user_id := claims.(jwt.MapClaims)["user_id"].(string)

	req := pb.CategoryCReq{
		UserId: user_id,
		Name:   body.Name,
		Type:   body.Type,
	}

	_, err := h.Category.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}

// GetCategory godoc
// @Summary     Get Category
// @Description Get category by ID
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Category ID"
// @Success     200     {object} pb.CategoryGRes "Category details"
// @Failure     400     {object} string "Invalid category ID"
// @Failure     404     {object} string "Category not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/category/{id} [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetCategory(c *gin.Context) {
	categoryId := c.Param("id")

	res, err := h.Category.GetByID(c, &pb.ByID{Id: categoryId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateCategory godoc
// @Summary     Update Category
// @Description Update category details
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       id       path     string             true "Category ID"
// @Param       category body     pb.CategoryUReqForSwagger true "Category update request"
// @Success     200     {object} string "Category updated successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     404     {object} string "Category not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/category/{id} [put]
// @Security    BearerAuth
func (h *HTTPHandler) UpdateCategory(c *gin.Context) {
	categoryId := c.Param("id")
	var body pb.CategoryUReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req := pb.CategoryUReq{
		Id:   categoryId,
		Name: body.Name,
	}

	_, err := h.Category.Update(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// DeleteCategory godoc
// @Summary     Delete Category
// @Description Delete category by ID
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Category ID"
// @Success     204     {object} string "Category deleted successfully"
// @Failure     400     {object} string "Invalid category ID"
// @Failure     404     {object} string "Category not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/category/{id} [delete]
// @Security    BearerAuth
func (h *HTTPHandler) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")

	_, err := h.Category.Delete(c, &pb.ByID{Id: categoryId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// GetAllCategories godoc
// @Summary     List Categories
// @Description List all categories with pagination
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       user_id      query     string         false "User ID"
// @Param       limit        query     int            false "Pagination limit"
// @Param       offset       query     int            false "Pagination offset"
// @Success     200     {object} pb.CategoryGARes "List of categories"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Router      /v1/categories [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetAllCategories(c *gin.Context) {
	var req pb.CategoryGAReq

	req.UserId = c.Query("user_id")

	if limitStr := c.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
		req.Pagination = &pb.Pagination{Limit: int64(limit)}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
		if req.Pagination == nil {
			req.Pagination = &pb.Pagination{}
		}
		req.Pagination.Offset = int64(offset)
	}

	res, err := h.Category.GetAll(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
