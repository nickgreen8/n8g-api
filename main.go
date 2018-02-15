package main

import (
    "log"

    "github/nickgreen8/n8g-api/utils"
)

// our main function
func main() {
    // Add dummy data
    fixtures = append(fixtures, Fixture{ID: 1, Opposition: "Preston Grasshoppers", Date: "02/12/2017", Time: "15.00", Competition: "Northern Premier", Venue: "H", Location: "Scatchard Lane", Result: "L", Score: "10-38", Report: 1})
    fixtures = append(fixtures, Fixture{ID: 2, Opposition: "Ilkley", Date: "16/12/2017", Time: "14.15", Competition: "Northern Premier", Venue: "A", Location: "Ilkley Rugby Club", Result: "L", Score: "25-23", Report: 2})
    fixtures = append(fixtures, Fixture{ID: 3, Opposition: "Billingham", Date: "06/01/2018", Time: "14.15", Competition: "Northern Premier", Venue: "H", Location: "Scatchard Lane", Result: "L", Score: "13-14", Report: 3})
    fixtures = append(fixtures, Fixture{ID: 4, Opposition: "Sandal", Date: "13/01/2018", Time: "14.15", Competition: "Northern Premier", Venue: "A", Location: "Sandal Rugby Club", Result: "W", Score: "17-24", Report: 4})
    fixtures = append(fixtures, Fixture{ID: 5, Opposition: "Wirral", Date: "20/01/2018", Time: "14.15", Competition: "Northern Premier", Venue: "H", Location: "Scatchard Lane", Result: "W", Score: "22-19", Report: 5})
    fixtures = append(fixtures, Fixture{ID: 6, Opposition: "Hull", Date: "27/01/2018", Time: "14.15", Competition: "Northern Premier", Venue: "A", Location: "Hull Rugby Club"})
    fixtures = append(fixtures, Fixture{ID: 7, Opposition: "Birkenhead Park", Date: "03/02/2018", Time: "14.15", Competition: "Northern Premier", Venue: "H", Location: "Scatchard Lane"})
    fixtures = append(fixtures, Fixture{ID: 8, Opposition: "Harrogate", Date: "10/02/2018", Time: "14.15", Competition: "Northern Premier", Venue: "A", Location: "Harrogate Rugby Club"})
    fixtures = append(fixtures, Fixture{ID: 9, Opposition: "Kendal", Date: "17/02/2018", Time: "14.15", Competition: "Northern Premier", Venue: "A", Location: "Kendal Rugby Club"})
    fixtures = append(fixtures, Fixture{ID: 10, Opposition: "Kirkby Lonsdale", Date: "03/03/2018", Time: "14.15", Competition: "Northern Premier", Venue: "H", Location: "Scatchard Lane"})
    squad = append(squad, Player{ID: 1, Name: "Nick Green", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 2, Name: "Dan Ripley", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 3, Name: "Steve Redhead", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 4, Name: "James Glover", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 5, Name: "Harry Baylis", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 6, Name: "Robbie McFaul", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 7, Name: "Chris Norton", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 8, Name: "Joe Firth", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 9, Name: "Jamie Denison", Picture: "CameraShy.jpg"})
    squad = append(squad, Player{ID: 10, Name: "Mikey Dixon", Picture: "CameraShy.jpg"})

    // Start the server
    log.Fatal(http.ListenAndServe(":8000", NewRouter()))
}
