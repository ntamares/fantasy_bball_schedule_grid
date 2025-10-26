export function formatDateHeaders(dates) {
    return dates.map(dateStr => {
        // Remove the time and timezone part, treat as local date
        const dateOnly = dateStr.split('T')[0];  // "2025-10-20T00:00:00Z" → "2025-10-20"
        const date = new Date(dateOnly + 'T00:00:00');  // Add local midnight time

        const dayName = date.toLocaleDateString('en-US', { weekday: 'short' });
        const monthNumber = date.getMonth() + 1;
        const dayNumber = date.getDate();

        return `${dayName} ${monthNumber}/${dayNumber}`;
    });
}