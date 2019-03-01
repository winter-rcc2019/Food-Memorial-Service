package controller

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	"strconv"

	"server/mysql/sql_model"
	"server/mysql/view"

	//"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
    "github.com/go-xorm/xorm"
)

func GetFoods(w http.ResponseWriter, r *http.Request) {
	foods, err := sql_model.GetFoods()
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get foods error: %v", err))
		return
	}
	view.RenderFoods(w, foods)
}

// GetFood は path に含まれる uuid に一致する foods テーブルの レコードを返す
func GetFood(c *gin.Context) {
	w := c.Writer
	food_UUID := c.Param("uuid")
	//exist, err := sql_model.CheckFoodExist(food_UUID)
	//if err != nil {
		//view.RenderInternalServerError(w, fmt.Sprintf("check food exist error: %v", err))
		//return
	//}
	//if !exist {
		//view.RenderNotFound(w, "foods", food_UUID)
		//return
	//}

	food, err := sql_model.GetFood(food_UUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get foods error: %v", err))
		return
	}
	view.RenderFood(w, food, http.StatusOK)
}

func CreateFood(c *gin.Context) {
	w := c.Writer
	var food sql_model.Food
	food.Image_ID = c.PostForm("uuid")
	food.Comment = c.PostForm("comment")
	food.User_Name = c.PostForm("user_name")
	food.User_ID, _ = strconv.Atoi(c.PostForm("user_id"))
	food.Reply = c.PostFormMap("reply")
	engine, _ := xorm.NewEngine("mysql","root:password@/food_memory")
	engine.Insert(food)
	insertID, err := sql_model.CreateFood(&food)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create food error: %v", err))
		return
	}
	createdFood, err := sql_model.GetFoodByID(insertID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get food error: %v", err))
		return
	}
	view.RenderFood(w, createdFood, http.StatusCreated)
	c.JSON(200, food)
}	

func PutFood(c *gin.Context, w http.ResponseWriter) {
	food_UUID := c.Param("uuid")
	//exist, err := sql_model.CheckFoodExist(food_UUID)
	//if err != nil {
		//view.RenderInternalServerError(w, fmt.Sprintf("check food exist error: %v", err))
		//return
	//}
	//if !exist {
		//view.RenderNotFound(w, "foods", food_UUID)
		//return
	//}
	//body := c.PostForm("comment")
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
		//view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		//return
	//}

	//var food sql_model.Food
	//err := json.Unmarshal(body, &food)
	//if err != nil {
		//view.RenderInternalServerError(w, fmt.Sprintf("json parse error: %v", err))
		//return
	//}

	//err = sql_model.UpdateFood(&food, food_UUID)
	//if err != nil {
		//view.RenderInternalServerError(w, fmt.Sprintf("create food error: %v", err))
		//return
	//}
	updatedFood, err := sql_model.GetFood(food_UUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get food error: %v", err))
		return
	}
	view.RenderFood(w, updatedFood, http.StatusOK)
}

func DeleteFood(c *gin.Context) {
	w := c.Writer
	food_UUID := c.Param("uuid")
	//exist, err := sql_model.CheckFoodExist(food_UUID)
	//if err != nil {
		//view.RenderInternalServerError(w, fmt.Sprintf("check food exist error: %v", err))
		//return
	//}
	//if !exist {
		//view.RenderNotFound(w, "foods", food_UUID)
		//return
	//}
	
	err := sql_model.DeleteFood(food_UUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create food error: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id: %s is deleted!", food_UUID)})
}