package handler

import (
	"StudentManagement/storage"
	"log"
	"net/http"
)


func (h connection) MarksHandler(w http.ResponseWriter, r *http.Request, class int, studentID int) error {

	subject, err := h.storage.GetSubjectByClassID(class)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	for _, s := range subject {
		b := storage.StudentSubject{
			StudentID: studentID,
			SubjectID: s.ID,
			Marks:     0,
		}

		_, err := h.storage.InsertMark(b)
		if err != nil {
			log.Fatalf("%v", err)
			return err
		}
	}
	return nil
}