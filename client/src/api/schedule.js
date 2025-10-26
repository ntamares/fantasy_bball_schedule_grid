import { config } from '../config/environment.js';

export async function fetchWeeklySchedule() {
    const url = `${config.apiBaseUrl}/schedule`;

    try {
        const response = await fetch(url);

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: 'Unknown error' }));
            console.error(`API Error (${response.status}):`, errorData);
            throw new Error(`Failed to fetch weekly schedule: ${response.status} ${response.statusText}`);
        }

        const result = await response.json();

        return result;
    } catch (error) {
        console.error('Error fetching weekly schedule: ', error);

        if (error instanceof TypeError && error.message.includes('fetch')) {
            throw new Error('Network error: Unable to connect to the server');
        } else if (error.message.includes('Failed to fetch weekly schedule')) {
            throw error;
        } else {
            throw new Error('Failed to process weekly schedule data');
        }
    }
}