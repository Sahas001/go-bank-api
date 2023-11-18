package main


func main() {
  store, err := NewPostgresStore()
  if err != nil {
    
  }

	server := NewAPIServer(":8080")
	server.Run()
}
