package main

import (
	"fmt"
)

type WeaponType int

// const (
// 	axe       WeaponType = 1
// 	sword     WeaponType = 2
// 	knife     WeaponType = 3
// 	woodstick WeaponType = 4
// )

func (weapon WeaponType) String() string {
	switch weapon {
	case Axe:
		return "axe"
	case Sword:
		return "woodstick"

	case Knife:
		return "knife"

	case Woodstick:
		return "woodstick"

	default:
		return "unknown"
	}
}

const (
	Axe WeaponType = iota
	Sword
	Knife
	Woodstick
)

func getWeaponEnum(weaponType WeaponType) string {

	switch weaponType {
	case 1:
		return "Axe"
	case 2:
		return "Sword"
	case 3:
		return "Knife"
	case 4:
		return "Woodstick"

	default:
		return "unknown"

	}
}

func getWeapons(weaponType WeaponType) int {

	switch weaponType {

	case Axe:
		return 30
	case Sword:
		return 10

	case Woodstick:
		return 1
	case Knife:
		return 20

	default:
		panic("unsupported weapon type")

	}
}

// func getWeapons(weaponType string) int {

// 	switch weaponType {

// 	case "axe":
// 		return 30
// 	case "sword":
// 		return 10

// 	case "woodstrick":
// 		return 1
// 	case "knife":
// 		return 20

// 	default:
// 		panic("unsupported weapon type")

// 	}
// }

func main() {

	fmt.Printf("Starting with Weapon (%s),(%d)\n", Axe, getWeapons(Axe))
	fmt.Printf("Starting with Weapon (%s),(%d)\n", Sword, getWeapons(Sword))
	fmt.Println("Starting with weapon", getWeaponEnum(1))

}
