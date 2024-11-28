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
     ```sql
     CREATE DATABASE api_assignment;
     USE api_assignment;
     ```

3. **Create the Table**:
   - Create a table named `time_log` to store API request timestamps:
     ```sql
     CREATE TABLE time_log (
         id INT AUTO_INCREMENT PRIMARY KEY,
         timestamp DATETIME NOT NULL
     );
     ```

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

---

## Time Zone Conversion

The application uses Go’s `time` package to handle time zone conversion to Toronto's local time.

### Implementation:
```go
const torontoTimeZone = "America/Toronto"

func getTorontoTime() (time.Time, error) {
    // Load Toronto timezone
    location, err := time.LoadLocation(torontoTimeZone)
    if err != nil {
        return time.Time{}, fmt.Errorf("failed to load timezone: %w", err)
    }

    // Get the current time in Toronto
    return time.Now().In(location), nil
}
