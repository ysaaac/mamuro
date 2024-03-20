export function formatDate(inputDate: string): string {
    const months: string[] = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
    const date: Date = new Date(inputDate);
    const today: Date = new Date();

    const addZero = (num: number): string => (num < 10 ? "0" + num : num.toString());

    if (date.getFullYear() < today.getFullYear()) {
        // Past years: day/month/year
        return `${date.getDate()} ${months[date.getMonth()]} ${date.getFullYear()}`;
    } else if (date.getMonth() < today.getMonth()) {
        // Current year, past months: day/month
        return `${date.getDate()} ${months[date.getMonth()]}`;
    } else if (date.getDate() < today.getDate()) {
        // Current month, past days: day month
        return `${date.getDate()} ${months[date.getMonth()]}`;
    } else {
        // Today: hours:minutes
        return `${addZero(date.getHours())}:${addZero(date.getMinutes())}`;
    }
}

// Test cases
// console.log(formatDate("Mon, 29 Oct 2001 14:23:26 -0800 (PST)")); // Output: 29 Oct 2001
// console.log(formatDate("Tue, 5 Feb 2002 17:42:58 -0800 (PST)")); // Output: 5 Feb
// console.log(formatDate("Tue, 25 Dec 2001 23:00:02 -0800 (PST)")); // Output: 25 Dec 2001
// console.log(formatDate("Mon, 4 Feb 2002 15:34:55 -0800 (PST)")); // Output: 15:34
