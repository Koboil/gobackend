package models

import (
	"errors"
	"fairground-backend/db"
	"strconv"
)

type Parent struct {
	User
	TokensBalance int `json:"tokens_balance"`
	PointsEarned  int `json:"points_earned"`
}

func GetParentByUserID(userID int) (*Parent, error) {

	// Step 1: Retrieve parent by user ID
	var parent Parent
	err := db.DB.QueryRow("SELECT parent_id, tokens_balance, points_earned FROM parents WHERE parent_id = $1", userID).Scan(&parent.UserID, &parent.TokensBalance, &parent.PointsEarned)

	if err != nil {
		return nil, err
	}

	// Step 2: Return the parent
	return &parent, nil

}

func GetAllParents() ([]User, error) {
	// Step 1: Retrieve all users
	rows, err := db.DB.Query("SELECT user_id, first_name, last_name, email, role FROM users where role = 'parent'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step 2: Iterate over the rows and create a slice of users
	var parents []User
	for rows.Next() {
		var parent User
		if err := rows.Scan(&parent.UserID, &parent.FirstName, &parent.LastName, &parent.Email, &parent.Role); err != nil {
			return nil, err
		}
		parents = append(parents, parent)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step 3: Return the slice of users
	return parents, nil
}

func DeleteStudentFromParent(studentId string, parentId string) error {

	// Step 1: Delete student from parent_student table
	_, err := db.DB.Exec("DELETE FROM parent_student WHERE parent_id = $1 AND student_id = $2", parentId, studentId)
	if err != nil {
		return err
	}

	// Step 2: Return success message
	return nil
}

func AddStudentToParent(studentId string, parentId string) error {

	// Step 1: Check if are you already the student's parent
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM parent_student WHERE parent_id = $1 AND student_id = $2", parentId, studentId).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("you are already the parent of this student")
	}

	// Step 1: Check if student has less than 2 parents
	var count2 int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM parent_student WHERE student_id = $1", studentId).Scan(&count2)
	if err != nil {
		return err
	}

	// Step: 2 - Return error if student has more than 2 parents
	if count2 >= 2 {
		return errors.New("can't add more parents to this student")
	}

	// Step: 3 - Add student to parent_student table
	_, err = db.DB.Exec("INSERT INTO parent_student (parent_id, student_id) VALUES ($1, $2)", parentId, studentId)
	if err != nil {
		return err
	}

	// Step: 4 - Return success message
	return nil

}

func GetParentsByStudentID(studentID int) ([]Parent, error) {

	// Step: 1 - Get all student_id from parent_student table with the parent ID
	rows, err := db.DB.Query("SELECT parent_id FROM parent_student WHERE student_id = $1", studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Step: 2 - Use GetStudentByUserID to get all students using the student_id
	var parents []Parent
	for rows.Next() {
		var parentID int
		if err := rows.Scan(&parentID); err != nil {
			return nil, err
		}
		// Step: 2a - Get student from student_id
		parent, err := GetParentByUserID(parentID)
		if err != nil {
			return nil, err
		}

		// Step: 2b - Get user from student_id
		user, err := GetUserByID(strconv.Itoa(parentID))
		if err != nil {
			return nil, err
		}

		// Step: 2c - Add user and student to students
		parent.User = *user
		parents = append(parents, *parent)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Step: 3 - If there are no students, return an empty slice
	if len(parents) == 0 {
		return []Parent{}, nil
	}

	// Step: 4 - Return the students
	return parents, nil

}
