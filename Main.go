package main

import (
	"fmt"
	"log"
	"./ProCon2018"
)

func main() {
	/* JSON組立 ----------------------------------------------------------------------------------------------------- */
	// tiles で指定する y, x は 0-index <-？
	status := ProCon2018.Status {
		Field: ProCon2018.Field {
				Scores: []int{
					-2, 1, 0, 1, 2, 0, 2, 1, 0, 1, -2,
					1, 3, 2, -2, 0, 1, 0, -2, 2, 3, 1,
					1, 3, 2, 1, 0, -2, 0, 1, 2, 3, 1,
					2, 1, 1, 2, 2, 3, 2, 2, 1, 1, 2,
					2, 1, 1, 2, 2, 3, 2, 2, 1, 1, 2,
					1, 3, 2, 1, 0, -2, 0, 1, 2, 3, 1,
					1, 3, 2, -2, 0, 1, 0, -2, 2, 3, 1,
					-2, 1, 0, 1, 2, 0, 2, 1, 0, 1, -2,
				},
				Height: 8,
				Width: 11,
			},
		Teams: []ProCon2018.Team{
			ProCon2018.Team {
				Tiles: []ProCon2018.Tile {
					ProCon2018.Tile {
						Y: 1,
						X: 1,
					},
					ProCon2018.Tile {
						Y: 6,
						X: 9,
					},
				},
			},
			ProCon2018.Team {
				Tiles: []ProCon2018.Tile {
					ProCon2018.Tile {
						Y: 1,
						X: 9,
					},
					ProCon2018.Tile {
						Y: 6,
						X: 1,
					},
				},
			},
		},
	}

	// JSON文字列化
	str, err := ProCon2018.GenerateJSONString(status)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(str) // おみせいたす
	/* ---------------------------------------------------------------------------------------------------------------- */
	
	/* POSTリクエスト ---------------------------------------------------------------------------------------------- */
	ProCon2018.GetResult(status);
	/* ---------------------------------------------------------------------------------------------------------------- */

	/* レスポンス解析 ----------------------------------------------------------------------------------------------- */
	/* ---------------------------------------------------------------------------------------------------------------- */
}

