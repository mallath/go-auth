I had created the authentication and authorization using oauth2
initially I had urn with go run main.go
In  the web browser http://localhost:8080/google_login it directs to Google login page
Later run above command 
in the web browser http://localhost:8080/google_callback it directs to Google login page  and check user credintils and then request to
google oauth server ,when it gives permission it return json format data as data in that it contains user details like their emailid,name,picture details
