export function getTeamGameCount(scheduleData) {
    const teamGameCount = {};

    scheduleData.forEach(game => {
        const homeTeamName = game.home.name;
        const awayTeamName = game.away.name;

        teamGameCount[homeTeamName] = (teamGameCount[homeTeamName] || 0) + 1;
        teamGameCount[awayTeamName] = (teamGameCount[awayTeamName] || 0) + 1;
    });

    const sortedEntries = Object.entries(teamGameCount)
        .sort((a, b) => a[0].localeCompare(b[0]));

    return Object.fromEntries(sortedEntries);
}