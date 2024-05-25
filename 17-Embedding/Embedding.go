package main

import "fmt"

// <結構變數>.<內嵌結構型別>.<欄位名稱>
// <結構變數>.<提升的欄位名稱>

type id int          // 自訂一個名為 name 的字串型別
type hunter struct { // 定義名為 hunter 的結構型別
	role string
	age  int
}
type skill struct { // 定義名為 skill 的結構型別
	ability string
	clan    string
}

// 內嵌(embedding)
type myFavorite struct { // 定義名為 myFavorite 的結構型別
	id     // 內嵌自訂型別 => type id int
	hunter // 內嵌結構 	=> type hunter struct 宣言
	skill  // 內嵌結構  => type skill struct 宣言
}

func getMyFavorites() []myFavorite {
	var favorite1 myFavorite // 定義 favorite1 為 myFavorite 的結構變數，沒給初始值就會是零值
	favorite2 := myFavorite{}

	// type id int
	favorite2.id = 2 // 賦予結構欄位初始值

	// type hunter struct
	favorite2.role = "奇犽"
	favorite2.age = 12

	// type myFavorite struct
	favorite2.ability = "變化系"
	favorite2.clan = "揍敵客"

	favorite3 := myFavorite{
		id:     3,
		hunter: hunter{role: "小傑", age: 12},
		skill:  skill{ability: "強化系", clan: "富力士"},
	}

	favorite4 := myFavorite{}
	favorite4.id = 4
	favorite4.hunter.role = "西索"
	favorite4.hunter.age = 26
	favorite4.skill.ability = "變化系"
	favorite4.skill.clan = "莫羅"

	return []myFavorite{favorite1, favorite2, favorite3, favorite4}
}

func main() {
	favorites := getMyFavorites()
	for i := 0; i < len(favorites); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, favorites[i])
	}
}
