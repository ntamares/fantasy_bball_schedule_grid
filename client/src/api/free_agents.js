import { config } from '../config/environment.js';


export async function fetchFreeAgents() {
    const url = `${config.apiBaseUrl}/freeAgents`;

    try {
        const response = await fetch(url);

        if (!response.ok) {
            const errorData = await response.json().catch(() => ({ message: 'Unknown error' }));
            console.error(`API Error (${response.status}):`, errorData);
            throw new Error(`Failed to fetch free agents: ${response.status} ${response.statusText}`);
        }

        const result = await response.json();

        return result;
    } catch (error) {
        console.error('Error fetching free agents: ', error);

        if (error instanceof TypeError && error.message.includes('fetch')) {
            throw new Error('Network error: Unable to connect to the server');
        } else if (error.message.includes('Failed to fetch free agents')) {
            throw error;
        } else {
            throw new Error('Failed to process free agent data');
        }
    }
}