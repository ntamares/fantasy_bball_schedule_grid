const API_BASE_URL = 'http://localhost:8080/api';

export async function fetchGameDates() {
    const url = `${API_BASE_URL}/dates`;

    try {
        const response = await fetch(url);

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: 'Unknown error' }));
            console.error(`API Error (${response.status}):`, errorData);
            throw new Error(`Failed to fetch game dates: ${response.status} ${response.statusText}`);
        }

        const result = await response.json();

        return result;
    } catch (error) {
        console.error('Error fetching game dates: ', error);

        if (error instanceof TypeError && error.message.includes('fetch')) {
            throw new Error('Network error: Unable to connect to the server');
        } else if (error.message.includes('Failed to fetch game dates')) {
            throw error;
        } else {
            throw new Error('Failed to process game dates data');
        }
    }
}