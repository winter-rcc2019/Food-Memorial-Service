package sql_model

import (
	//"time"

	"server/mysql/db"
	"github.com/google/uuid"
)

type Food struct {
	ID        	int       			`json:"id"`
	UUID      	string    			`json:"uuid"`
	//Images		[]Images			`json:"images"`
	Image_ID	string				`json:"image_id"`
	Comment		string    			`json:"comment"`
	User_Name	string    			`json:"user_name"`
	User_ID		int					`json:"user_id"`
	Reply    	map[string]string   `json:"reply"`
	//CreatedAt 	time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
}

//type Images struct {
	//Image		[]byte				`json:"image"`
//}

func GetFoods() ([]*Food, error) {
	conn, err := db.Init()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var foods []*Food

	rows, err := conn.Query(`
	  SELECT
		id,
		uuid,
		image_id, 
		comment, 
		user_name,
		user_id,
		reply, 
		from foods`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var food Food
		//var createdDatetime string
		//var updateDatetime string
		err := rows.Scan(
			&(food.ID),
			&(food.UUID),
			&(food.Image_ID),
			&(food.User_Name),
			&(food.User_ID),
			&(food.Comment),
			&(food.Reply))

		if err != nil {
			return nil, err
		}

		foods = append(foods, &food)
	}

	return foods, nil
}


func GetFood(foodUUID string) (*Food, error) {
	conn, err := db.Init()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var food Food
	//var createdDatetime string
	//var updateDatetime string

	err = conn.QueryRow(`
      	SELECT 
			id,
			uuid,
			image_id, 
			comment, 
			user_name,
			user_id,
			reply, 
			from foods
		where uuid = ?,
		taskUUID`).Scan(
				&(food.ID),
				&(food.UUID),
				&(food.Image_ID),
				&(food.Comment),
				&(food.User_Name),
				&(food.User_ID),
				&(food.Reply),
	)
	if err != nil {
		return nil, err
	}

	return &food, nil
}

//func CheckFoodExist(foodUUID string) (bool, error) {
	//conn, err := db.Init()
	//if err != nil {
		//return false, err
	//}
	//defer conn.Close()

	//var fetchRecordCount int
	//err = conn.QueryRow(`
      //SELECT
       //count(*)
       //from foods
       //where uuid = ?`,
		//foodUUID).Scan(
		//&fetchRecordCount,
	//)
	//if err != nil {
		//return false, err
	//}

	//if fetchRecordCount > 0 {
		//return true, nil
	//}
	//return false, nil
//}

func CreateFood(food *Food) (int64, error) {
	conn, err := db.Init()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	result, err := conn.Exec(
		`INSERT INTO foods (
			uuid,
			image_id, 
			user_name,
			user_id,
			comment,
			reply
			) VALUES (?, ?, ?, ?, ?) `,
		uuid.New(),
		food.Image_ID,
		food.User_Name,
		food.User_ID,
		food.Comment,
		food.Reply,
	)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func GetFoodByID(foodID int64) (*Food, error) {
	conn, err := db.Init()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var food Food
	//var createdDatetime string
	//var updateDatetime string

	err = conn.QueryRow(`
		SELECT 
			id,
			uuid,
			image_id,
			user_name,
			user_id,
			comment,
			reply
		from foods
		where id = ?`,
				foodID).Scan(
				&(food.ID),
				&(food.UUID),
				&(food.Image_ID),
				&(food.User_Name),
				&(food.User_ID),
				&(food.Comment),
				&(food.Reply),
				)
	if err != nil {
		return nil, err
	}

	return &food, nil
}

func UpdateFood(food *Food, foodUUID string) error {
	conn, err := db.Init()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(
		`UPDATE foods 
			SET image = ?, 
				user_name = ?, 
				user_id = ?,
				comment = ?,
				reply = ?
		 WHERE uuid = ?`,
			food.Image_ID,
			food.User_Name,
			food.User_ID,
			food.Comment,
			food.Reply,
			foodUUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFood(foodUUID string) error {
	conn, err := db.Init()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Exec(
		`DELETE FROM foods 
         WHERE uuid = ?`,
		foodUUID,
	)
	if err != nil {
		return err
	}
	return nil
}