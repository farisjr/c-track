package controllers

// import (
// 	"app/config"
// 	"app/models"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/labstack/echo"
// 	"github.com/stretchr/testify/assert"
// )

// // func tesGetTestCategoriesController(c echo.Context) error {
// // 	var testCategories []models.TestCategories
// // 	err := config.DB.Debug().Model(&models.TestCategories{}).Find(&testCategories).Error
// // 	//testCategories, err := database.GetTestCategories()
// // 	if err != nil {
// // 		return c.JSON(http.StatusInternalServerError, models.TestCategoriesResponse{
// // 			false, "Failed get database test category", nil,
// // 		})
// // 	}
// // 	return c.JSON(http.StatusOK, models.TestCategoriesResponse{
// // 		true, "Success", testCategories,
// // 	})
// // }

// // func AddTestCategoriesData() bool {
// // 	testctg := models.TestCategories{ID: 1, TestCategoriesName: "pcr"}
// // 	err := config.DB.Create(&testctg)
// // 	if err == nil {
// // 		return false
// // 	}
// // 	return true
// // }

// // func TestGetTestCategoriesController(t *testing.T) {
// // 	config.InitDBTest()
// // 	e := echo.New()
// // 	config.DB.Migrator().DropTable(&models.TestCategories{})
// // 	config.DB.Migrator().AutoMigrate(&models.TestCategories{})
// // 	AddTestCategoriesData()
// // 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// // 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// // 	rec := httptest.NewRecorder()
// // 	c := e.NewContext(req, rec)

// // 	c.SetPath("/testcategories")
// // 	if assert.NoError(t, tesGetTestCategoriesController(c)) {
// // 		assert.Equal(t, http.StatusOK, rec.Code)
// // 		body := rec.Body.String()
// // 		var responseTestCategories models.TestCategoriesResponse
// // 		fmt.Println(body)
// // 		json.Unmarshal([]byte(body), &responseTestCategories)

// // 		assert.Equal(t, responseTestCategories.Status, true)
// // 		assert.Equal(t, len(responseTestCategories.Data), 1)
// // 		assert.Equal(t, responseTestCategories.Data[0].TestCategoriesName, "pcr")
// // 	}
// // }

// // func TestFailGetTestCategoriesController(t *testing.T) {
// // 	config.InitDBTest()
// // 	e := echo.New()
// // 	config.DB.Migrator().DropTable(&models.TestCategories{})
// // 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// // 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// // 	rec := httptest.NewRecorder()
// // 	c := e.NewContext(req, rec)

// // 	c.SetPath("/testcategories")
// // 	if assert.NoError(t, tesGetTestCategoriesController(c)) {
// // 		assert.Equal(t, http.StatusInternalServerError, rec.Code)
// // 		body := rec.Body.String()
// // 		var responseTestCategories models.TestCategoriesResponse
// // 		fmt.Println(body)
// // 		json.Unmarshal([]byte(body), &responseTestCategories)

// // 		assert.Equal(t, responseTestCategories.Status, false)
// // 		assert.Equal(t, len(responseTestCategories.Data), 0)
// // 	}
// // }
