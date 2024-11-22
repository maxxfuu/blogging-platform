// Connect to database 

var db *sql.DB 

func main() {
	// Capture connection properties. 
	type cfg struct = mysql.Config{
		User: os.Getenv("DBUSER"), 
		Password, os.Getenv("DBUSER"), 
		Net: "tcp" 
		Addr:
		
	}
}