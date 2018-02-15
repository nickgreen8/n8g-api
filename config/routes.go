package config

import (
    "net/http"

    "github/nickgreen8/n8g-api/modals"
)

// Types

type Route struct {
    Name    string
    Method  string
    Pattern string
    Handler http.HandlerFunc
}
type Routes []Route

// Variables

var routes = Routes{
    // Fixtures
    Route{"GetAllFixtures", "GET", "/fixtures", GetFixtures},
    Route{"GetFixturesByTeam", "GET", "/fixtures/{team}", GetFixtures},
    Route{"GetFixtureByID", "GET", "/fixture/{id}", GetFixture},
    Route{"AddFixture", "POST", "/fixture", AddFixture},
    Route{"DeleteFixture", "DELETE", "/fixture/{id}", DeleteFixture},
    Route{"EditFixture", "PATCH", "/fixture/{id}", EditFixture},
    // Players
    Route{"GetAllPlayers", "GET", "/players", GetSquad},
    Route{"GetPlayersByTeam", "GET", "/players/{team}", GetSquad},
    Route{"GetPlayerByID", "GET", "/player/{id}", GetPlayer},
    // Events
    Route{"GetAllEvents", "GET", "/events", GetEvents},
    Route{"GetEventByTimePeriod", "GET", "/events/{period}", GetEvent}, // period = (today|happened|to-come)
    Route{"GetEventByID", "GET", "/event/{id}", GetEvent},
    // League
    Route{"GetLeagueTableForTeam", "GET", "/league/{team}", GetLeague},
    // News
    Route{"GetNewsItems", "GET", "/news", GetNews},
    Route{"GetNewsStory", "GET", "/news/{id}", GetNewsStory},
    // Pages
    Route{"GetPage", "GET", "/page/{id}", GetPage},
    // Team
    Route{"GetTeamData", "GET", "/team/{id}", GetTeam},
}

// Functions
func GetRoutes() *Routes {
    return routes
}
