export function buildScheduleGrid(teams, dates, scheduleData) {
    const grid = {};

    teams.forEach(team => {
        grid[team] = {};
        // TODO remove?
        dates.forEach(date => {
            grid[team][date] = null;
        });
    });

    scheduleData.forEach(({ date, home, away }) => {
        if (grid[home] && grid[home][date] !== undefined) {
            grid[home][date] = `vs ${away}`;
        }
        if (grid[away] && grid[away][date] !== undefined) {
            grid[away][date] = `vs ${home}`;
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
    //         grid[awayTeam][gameDate] = `@ ${homeTeam}`;
    //     }
    // });


    return grid;
}