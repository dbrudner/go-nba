# Golang wrapper around NBA stats API

## Info on NBA stats API

### [today.json.links](http://data.nba.net/10s/prod/v1/today.json)

### Params

**{{gameDate}}** - `string` - formated yyyymmdd
*   *example: 20171005* 

**{{playerId}}** - `string` - Can be retrieved from leagueRosterPlayers. Also called `personId` in some endpoints
*   *example: 203500* 

**{{gameId}}** - `string`
*   *example: 0011700027* 

**{{teamUrlCode}}** - `string` - This value is called "teamId" in teams.json
*   *example: 1610612738* 

### Endpoints
- "calendar": "/prod/v1/calendar.json",
- "todayScoreboard": "/prod/v1/20190121/scoreboard.json",
- "currentScoreboard": "/prod/v1/20190121/scoreboard.json",
- "scoreboard": "/prod/v2/{{gameDate}}/scoreboard.json",
- "teams": "/prod/v2/2018/teams.json",
- "leagueRosterPlayers": "/prod/v1/2018/players.json",
- "allstarRoster": "/prod/v1/allstar/2017/AS_roster.json",
- "leagueRosterCoaches": "/prod/v1/2018/coaches.json",
- "leagueSchedule": "/prod/v1/2018/schedule.json",
- "leagueConfStandings": "/prod/v1/current/standings_conference.json",
- "leagueDivStandings": "/prod/v1/current/standings_division.json",
- "leagueUngroupedStandings": "/prod/v1/current/standings_all.json",
- "leagueMiniStandings": "/prod/v1/current/standings_all_no_sort_keys.json",
- "leagueTeamStatsLeaders": "/prod/v1/2018/team_stats_rankings.json",
- "leagueLastFiveGameTeamStats": "/prod/v1/2018/team_stats_last_five_games.json",
- "previewArticle": "/prod/v1/{{gameDate}}/{{gameId}}_preview_article.json",
- "recapArticle": "/prod/v1/{{gameDate}}/{{gameId}}_recap_article.json",
     * Seems like this endpoint doesnt work
- "gameBookPdf": "/prod/v1/{{gameDate}}/{{gameId}}_Book.pdf",
- "boxscore": "/prod/v1/{{gameDate}}/{{gameId}}_boxscore.json",
- "miniBoxscore": "/prod/v1/{{gameDate}}/{{gameId}}_mini_boxscore.json",
- "pbp": "/prod/v1/{{gameDate}}/{{gameId}}_pbp_{{periodNum}}.json",
- "leadTracker": "/prod/v1/{{gameDate}}/{{gameId}}_lead_tracker_{{periodNum}}.json",
- "playerGameLog": "/prod/v1/2018/players/{{personId}}_gamelog.json",
- "playerProfile": "/prod/v1/2018/players/{{personId}}_profile.json",
- "playerUberStats": "/prod/v1/2018/players/{{personId}}_uber_stats.json",
     * Returns empty strings for all values
- "teamSchedule": "/prod/v1/2018/teams/{{teamUrlCode}}/schedule.json",
- "teamsConfig": "/prod/2018/teams_config.json",
- "teamRoster": "/prod/v1/2018/teams/{{teamUrlCode}}/roster.json",
- "teamsConfigYear": "/prod/{{seasonScheduleYear}}/teams_config.json",
- "teamScheduleYear": "/prod/v1/{{seasonScheduleYear}}/teams/{{teamUrlCode}}/schedule.json",
- "teamLeaders": "/prod/v1/2018/teams/{{teamUrlCode}}/leaders.json",
- "teamScheduleYear2": "/prod/v1/{{seasonScheduleYear}}/teams/{{teamId}}/schedule.json",
- "teamLeaders2": "/prod/v1/2018/teams/{{teamId}}/leaders.json",
- "playoffsBracket": "/prod/v1/2017/playoffsBracket.json",
- "playoffSeriesLeaders": "/prod/v1/2018/playoffs_{{seriesId}}_leaders.json"

### Current features

- Single player info
- Standings by conference
- Team roster
- Team schedule