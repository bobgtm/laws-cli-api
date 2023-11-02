package main

import "fmt"

func callBackLegSession(cfg *config, args ...string) error {
	sessions, err := cfg.lawsClient.ListSessions()
	if err != nil {
		return err
	}

	if len(sessions.Sessions) >= 5 {
		firstFive := sessions.Sessions[:5]
		for _, sess := range firstFive {
			fmt.Printf("(id: %v) %v\n", sess.SessionID, sess.SessionName)
		}
	}
	return nil
}
