export function formatDateHeaders(dates) {
    return dates.map(dateStr => {
        const dateOnly = dateStr.split('T')[0];
        const date = new Date(dateOnly + 'T00:00:00');

        const dayName = date.toLocaleDateString('en-US', { weekday: 'short' });
        const monthNumber = date.getMonth() + 1;
        const dayNumber = date.getDate();

        return `${dayName} ${monthNumber}/${dayNumber}`;
    });
}