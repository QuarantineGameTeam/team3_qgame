package database

/*
	This is where the client’s database connection code is implemented.
	Typically, the server side is implemented such that a new thread processes requests for the connection.
	A connection pool is implemented here to minimize outlet opening.
	Usually, you are not indifferent through the connection (if everyone is connected as the same user)
	that you received the database result set. You do not want to consume resources, so you want to be pleasant,
	and when you are done, you close the connection. I believe that every server ends the connection today if there
	is no activity for some time (timeout), that is, working with the database.
*/
