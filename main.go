package main

import "Zura-chanZura/app/infrastructure"

func main() {
	r := infrastructure.NewRouting()
	r.Run()
}