# Time_logger
# Go API Assignment: Current Time in Toronto

## Overview
This project demonstrates how to create a Go application with the following functionality:
- Connect to a MySQL database.
- Provide an API endpoint `/current-time` that returns the current time in Toronto's timezone as JSON.
- Log each API request's timestamp into a MySQL database.
- Handle errors for database operations and time zone conversions.

---

## Set Up MySQL Database

1. **Install MySQL**:
   - Download MySQL from [MySQL Downloads](https://dev.mysql.com/downloads/mysql/).
   - Set a root password during installation.

2. **Start MySQL Server**:
   - Create a database using the following commands:
3. **Create the Table**:
   - Create a table named `time_log` to store API request timestamps:
     ```sql
     CREATE TABLE time_log (
         id INT AUTO_INCREMENT PRIMARY KEY,
         timestamp DATETIME NOT NULL
     );
     ```
![image](https://github.com/user-attachments/assets/b95fe35f-5746-46b0-87eb-14e7554b6f7d)
![image](https://github.com/user-attachments/assets/1e3a1b85-9a6f-4c09-b2a2-174d404237ee)

---

## API Development

This Go application provides an API endpoint `/current-time` that returns the current time in Toronto.

### Steps:
1. Use Go’s `net/http` package to create a web server.
2. Use Go’s `time` package to handle and format time.
3. Set the timezone to Toronto using `time.LoadLocation`.
4. Format the current time using `time.Format` with the `2006-01-02 15:04:05` layout.
5. Encode the response as JSON and send it to the client.
6. The server runs on port `8080`.
![image](https://github.com/user-attachments/assets/47e10727-6df3-4355-84e9-c718beee0e99)
![image](https://github.com/user-attachments/assets/6c2d9722-681f-4c88-8aac-62ece79557c4)

---

## Time Zone Conversion

The application uses Go’s `time` package to handle time zone conversion to Toronto's local time.

func getTorontoTime() (time.Time, error) {
    // Load Toronto timezone
    location, err := time.LoadLocation(torontoTimeZone)
    if err != nil {
        return time.Time{}, fmt.Errorf("failed to load timezone: %w", err)
    }

    // Get the current time in Toronto
    return time.Now().In(location), nil
}

![image](https://github.com/user-attachments/assets/f6f780c7-dcc7-460c-b04f-af031bcb41c8)

## Database connection

Connect to your MySQL database from your Go application.
On each API call, insert the current time into the time_log table
![image](https://github.com/user-attachments/assets/0ff1481c-7796-4fc1-8c51-d2b1a338dc22)

![image](https://github.com/user-attachments/assets/a30a868a-f866-4a28-ba2e-93ebad744adc)

When we  user makes a GET request to http://localhost:8080/current-time.
The server:Converts the current system time to Toronto time.Logs this time into the database (time_log table).Sends the current Toronto time as a JSON response.
The database stores each request’s timestamp for tracking or auditing

## Return Time in JSON
![image](https://github.com/user-attachments/assets/1cd048a9-9770-465e-b9a3-b7358e4d12b6)
## Error Handling
![image](https://github.com/user-attachments/assets/2aecae65-e168-4a7b-ae78-07d31e24dcee)

'''func logTimeToDatabase(timestamp time.Time) error {
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err := db.Exec(query, timestamp)
	if err != nil {
		log.Printf("Error logging time to database: %v", err) // Log the error
		return fmt.Errorf("failed to insert time into database: %w", err) // Wrap and return the error
	}

	log.Printf("Time logged to database: %s", timestamp.Format("2006-01-02 15:04:05"))
	return nil
} '''
Errors are logged using log.Printf for debugging and monitoring purposes.
Errors encountered during API request processing are returned as HTTP 500 responses using http.Error.



