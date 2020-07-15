package eventtimer

import (
	"groupevent/internal/pkg/db"
	"log"
)

const updateeventsstatussql string = `UPDATE events 
SET 
    status = (CASE
        WHEN DATEDIFF(NOW(), start_date) < 0 THEN 'deactived'
        WHEN
            DATEDIFF(NOW(), start_date) > 0
                AND DATEDIFF(NOW(), expired_on) < 0
        THEN
            'actived'
        ELSE 'expired'
    END)
WHERE
    id > 0`

func UpdateEventStatus() error {
	log.Println("start update mysql...")
	_, err := db.NewDB().Exec(updateeventsstatussql)
	if err != nil {
		return err
	}

	return nil
}
