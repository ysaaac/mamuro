export enum FetchMethods {
    GET = "GET",
    POST = "POST"
}

export async function fetchData(endpoint: string, method: FetchMethods = FetchMethods.GET, baseUrl = 'http://localhost:8080') {
    const uri = baseUrl + endpoint;
    try {
        const response = await fetch(baseUrl + endpoint, {
            method,
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return await response.json();
    } catch (error: unknown) { // Specify the type of the caught error as 'unknown'
        if (error instanceof Error) {
            // This error is an instance of Error, handle accordingly
            throw new Error(`Error fetching data from ${uri}: ${error.message}`);
        } else {
            // Handle other error types
            throw new Error(`Unknown error fetching data from ${uri}`);
        }
    }
}
