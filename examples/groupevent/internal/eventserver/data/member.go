package data

import (
	"groupevent/internal/pkg/db"
)

const selectmemberssql string = "select student_name, student_id, g_m, college, level, profession from event_members where event_id=?"

func MembersBy(eventid string) ([]map[string]string, error) {
	rows, err := db.NewDB().Query(selectmemberssql, eventid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []map[string]string
	for rows.Next() {
		var (
			studentName, studentId, gm, college, level, profession string
		)
		err := rows.Scan(&studentName, &studentId, &gm, &college, &level, &profession)
		if err != nil {
			return nil, err
		}
		temp := make(map[string]string)
		temp["studentName"] = studentName
		temp["studentId"] = studentId
		temp["gm"] = gm
		temp["college"] = college
		temp["level"] = level
		temp["profession"] = profession
		result = append(result, temp)
	}
	return result, nil
}
