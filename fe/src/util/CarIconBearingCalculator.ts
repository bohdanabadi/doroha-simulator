export function calculateBearing(startLat: number, startLng: number, destLat: number, destLng: number): number {
    const startLatRad = degreesToRadians(startLat);
    const startLngRad = degreesToRadians(startLng);
    const destLatRad = degreesToRadians(destLat);
    const destLngRad = degreesToRadians(destLng);

    const y = Math.sin(destLngRad - startLngRad) * Math.cos(destLatRad);
    const x = Math.cos(startLatRad) * Math.sin(destLatRad) - Math.sin(startLatRad) * Math.cos(destLatRad) * Math.cos(destLngRad - startLngRad);
    const bearing = (radiansToDegrees(Math.atan2(y, x)) + 360) % 360;

    return bearing;
}

function degreesToRadians(degrees: number): number {
    return degrees * Math.PI / 180;
}

function radiansToDegrees(radians: number): number {
    return radians * 180 / Math.PI;
}