export function convertTimestampToLocalDateTime(timestamp?: number): string {
    const effectiveTimestamp = (timestamp !== undefined) ? timestamp * 1000 : Date.now();
    const date = new Date(effectiveTimestamp);
    return date.toLocaleTimeString();
}
export type HealthKey = 'g' | 'o' | 'r' | 'u';
export const HealthColorMapping = {
    g: {bgColor: 'bg-green-100', textColor: 'text-green-700', circleColor: 'bg-green-500'},
    o: {bgColor: 'bg-orange-100', textColor: 'text-orange-700', circleColor: 'bg-orange-500'},
    r: {bgColor: 'bg-red-100', textColor: 'text-red-700', circleColor: 'bg-red-500'},
    u: {bgColor: 'bg-gray-100', textColor: 'text-gray-700', circleColor: 'bg-gray-500'}, // Assuming 'U' stands for 'Undefined'
};