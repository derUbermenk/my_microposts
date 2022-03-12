package data

import (
	"time"

	_ "github.com/lib/pq"
)

type Micropost struct {
	Id        int
	Uuid      string
	Title     string
	Content   string
	CreatedAt time.Time
}

// gets all microposts
func Microposts() (microposts []Micropost, err error) {
	rows, err := Db.Query("SELECT * FROM microposts ORDER BY created_at DESC")

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		micropost := Micropost{}
		if err = rows.Scan(&micropost.Id, &micropost.Uuid,
			&micropost.Title, &micropost.Content, &micropost.CreatedAt); err != nil {
			panic(err)
		}
		microposts = append(microposts, micropost)
	}
	rows.Close()
	return
}

func NewMicropost(title, content string) (micropost Micropost) {
	micropost = Micropost{
		Uuid:      createUUID(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	return
}

func CreateMicropost(title, content string) (saved_micropost Micropost, err error) {
	micropost := NewMicropost(title, content)

	statement := "INSERT INTO microposts (uuid, title, content, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, title, content, created_at"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	// gets executed once function returns
	// 	if there was an error then this executes
	// 	if not then this executes after code below
	defer stmt.Close()

	err = stmt.QueryRow(micropost.Uuid, micropost.Title, micropost.Content, micropost.CreatedAt).Scan(
		&saved_micropost.Id, &saved_micropost.Uuid,
		&saved_micropost.Title, &saved_micropost.Content,
		&saved_micropost.CreatedAt,
	)
	return
}
