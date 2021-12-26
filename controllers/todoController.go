package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	guuid "github.com/google/uuid"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"body"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func CreateTodoTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError := db.CreateTable(&Todo{}, options)

	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}

	log.Printf("Todo table created")

	return nil
}

func GetAllTodos(c *gin.Context) {
	var todoList []Todo

	err := dbConnect.Model(&todoList).Select()

	if err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "error to getting list",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "todo List",
		"data":    todoList,
	})

	return
}

func CreateTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	title := todo.Title
	text := todo.Text
	comleted := todo.Completed
	id := guuid.New().String()
	insertError := dbConnect.Insert(&Todo{
		ID:        id,
		Title:     title,
		Text:      text,
		Completed: comleted,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if insertError != nil {
		log.Printf("Inserting error: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error Inserting",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created",
	})

	return
}

func GetTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}

	err := dbConnect.Select(todo)

	if err != nil {
		log.Printf("Error get todo: %v\n", err)

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo with this id not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo found",
		"data":    todo,
	})

	return
}

func EditTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	var todo Todo
	c.BindJSON(&todo)
	completed := todo.Completed

	_, err := dbConnect.Model(&Todo{}).Set("completed = ?", completed).Where("id = ?", todoId).Update()

	if err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Todo Edited Successfully",
	})

	return

}

func DeleteTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}

	err := dbConnect.Delete(todo)

	if err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted",
	})
	return
}
