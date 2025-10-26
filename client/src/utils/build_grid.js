export function buildScheduleGrid(teams, dates, scheduleData) {
    const grid = {};

    // Initialize grid with team names
    teams.forEach(team => {
        grid[team.name] = {};
        dates.forEach(date => {
            grid[team.name][date] = null;
        });
    });

    // Process games from your actual JSON structure
    scheduleData.forEach(game => {
        const gameDate = game.scheduled;           // Use "scheduled" field
        const homeTeamName = game.home.name;       // Extract nested home team name
        const awayTeamName = game.away.name;       // Extract nested away team name

        // Add home team entry
        if (grid[homeTeamName] && grid[homeTeamName][gameDate] !== undefined) {
            grid[homeTeamName][gameDate] = `vs ${game.away.alias}`;
        }

        // Add away team entry  
        if (grid[awayTeamName] && grid[awayTeamName][gameDate] !== undefined) {
            grid[awayTeamName][gameDate] = `@ ${game.home.alias}`;
        }
    });

    // TODO use instead?
    // scheduleData.forEach(game => {
    //     const gameDate = game.date;
    //     const homeTeam = game.homeTeam;
    //     const awayTeam = game.awayTeam;

    //     // Add home team entry
    //     if (grid[homeTeam]) {
    //         grid[homeTeam][gameDate] = `vs ${awayTeam}`;
    //     }

    //     // Add away team entry
    //     if (grid[awayTeam]) {
    //         grid[awayTeam][gameDate] = `@ ${homeTeam}`;date
    //     }
    // });


    return grid;
}