export function buildScheduleGrid(teams, dates, scheduleData) {
    const grid = {};

    teams.forEach(team => {
        grid[team.name] = {};
        dates.forEach(date => {
            grid[team.name][date] = null;
        });
    });

    scheduleData.forEach(game => {
        const gameDate = game.scheduled;
        const homeTeamName = game.home.name;
        const awayTeamName = game.away.name;

        if (grid[homeTeamName] && grid[homeTeamName][gameDate] !== undefined) {
            grid[homeTeamName][gameDate] = `${game.away.alias}`;
        }

        if (grid[awayTeamName] && grid[awayTeamName][gameDate] !== undefined) {
            grid[awayTeamName][gameDate] = `@ ${game.home.alias}`;
        }
    });

    return grid;
}