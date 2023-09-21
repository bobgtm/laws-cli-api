package main

import "os"

var key = os.Getenv("LAW_API_KEY")
var url = "http://api.legiscan.com/?key=" + key

// Per the LegiSCan API User Manual
// Most use cases for the api can be satsified with daily updates
// although recommended frequency varies for different operations
func commandShowMaster() error {
	// &op=getMasterList&id=(session ID to be retrieved from Show Session)"
	return nil
}

func commandShowSession() error {
	//op=getSessionList&state=FL

	//Return sessions.name, sessions.session_id, sessions.year_start,
	//sessions. year_end
	// sessions.session_id can be used for ShowMaster
	return nil
}
