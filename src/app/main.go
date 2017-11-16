package main

func main() {
	route := mainRoutes()

	route.Run(":3001")
}
