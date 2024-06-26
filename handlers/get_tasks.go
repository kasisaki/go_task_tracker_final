package handlers

import (
	"errors"
	"go_final_project/db"
	"go_final_project/models"
	"net/http"
	"strconv"
)

const TasksNumberLimit = 50

func HandleGetTasks(w http.ResponseWriter, req *http.Request) {
	// Проверяем GET-запрос или нет
	if req.Method == http.MethodGet {
		search := req.URL.Query().Get("search")
		tasks, err := db.GetTasksFromDB(search, TasksNumberLimit) // Получаем список задач из базы данных лим
		if err != nil {
			HandleError(w, http.StatusInternalServerError, err)
			return
		}
		respMap := make(map[string][]models.TaskDTO)
		respMap["tasks"] = tasks

		// Преобразуем список задач в JSON
		HandleNormalResponse(w, respMap)
		return
	}
}

func HandleGetTaskById(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		id, err := strconv.Atoi(req.URL.Query().Get("id"))
		if err != nil {
			HandleError(w, http.StatusBadRequest, errors.New("Не указан идентификатор"))
			return
		}

		task, err := db.GetTaskById(id)
		if HandleGetError(w, err) {
			return
		}

		HandleNormalResponse(w, task)
		return
	}
}
