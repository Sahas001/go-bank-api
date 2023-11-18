package main


func main() {
  store, err := NewPostgresStore()
  if err != nil {
    log.Fatal(err)
  }

	server := NewAPIServer(":8080")
	server.Run()
}
