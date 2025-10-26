export function formatDateHeaders(dates) {
    return dates.map(dateStr => {
        const date = new Date(dateStr);
        const dayName = date.toLocaleDateString('en-US', { weekday: 'short' });
        const monthDay = date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
        return {
            date: dateStr,
            display: `${dayName} ${monthDay}`
        };
    });
}