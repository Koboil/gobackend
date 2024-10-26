package models

import (
	"errors"
	"fairground-backend/db"
	"strconv"
)

type Student struct {
	User
	TokensBalance int `json:"tokens_balance"`
	PointsEarned  int `json:"points_earned"`
}

func GetStudentByUserID(userID int) (*Student, error) {

	// Step 1: Retrieve student by user ID
	var student Student
	err := db.DB.QueryRow("SELECT student_id, tokens_balance, points_earned FROM students WHERE student_id = $1", userID).Scan(&student.UserID, &student.TokensBalance, &student.PointsEarned)

	if err != nil {
		return nil, err
	}

	// Step 2: Return the student
	return &student, nil
}

func GetAllStudents() ([]User, error) {
	// Step 1: Retrieve all users
	rows, err := db.DB.Query("SELECT user_id, first_name, last_name, email, role FROM users where role = 'student'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of users
	var students []User
	for rows.Next() {
		var student User
		if err := rows.Scan(&student.UserID, &student.FirstName, &student.LastName, &student.Email, &student.Role); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: Return the slice of users
	return students, nil
}

func GetStudentsByParentID(parentID int) ([]Student, error) {

	// Step: 1 - Get all student_id from parent_student table with the parent ID
	rows, err := db.DB.Query("SELECT student_id FROM parent_student WHERE parent_id = $1", parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step: 2 - Use GetStudentByUserID to get all students using the student_id
	var students []Student
	for rows.Next() {
		var studentID int
		if err := rows.Scan(&studentID); err != nil {
			return nil, err
		}
		// Step: 2a - Get student from student_id
		student, err := GetStudentByUserID(studentID)
		if err != nil {
			return nil, err
		}

		// Step: 2b - Get user from student_id
		user, err := GetUserByID(strconv.Itoa(studentID))
		if err != nil {
			return nil, err
		}

		// Step: 2c - Add user and student to students
		student.User = *user
		students = append(students, *student)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step: 3 - If there are no students, return an empty slice
	if len(students) == 0 {
		return []Student{}, nil
	}

	// Step: 4 - Return the students
	return students, nil

}

func AddParentToStudent(parentId string, studentId string) error {

	// Step 1: Check if are you already the student of the parent
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM parent_student WHERE parent_id = $1 AND student_id = $2", parentId, studentId).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("you are already the student of this parent")
	}

	// Step 1: Check if student has less than 2 parents
	var count2 int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM parent_student WHERE student_id = $1", studentId).Scan(&count2)
	if err != nil {
		return err
	}

	// Step: 2 - Return error if student has more than 2 parents
	if count2 >= 2 {
		return errors.New("can't add more than two parents")
	}

	// Step: 3 - Add student to parent_student table
	_, err = db.DB.Exec("INSERT INTO parent_student (parent_id, student_id) VALUES ($1, $2)", parentId, studentId)
	if err != nil {
		return err
	}

	// Step: 4 - Return success message
	return nil

}
