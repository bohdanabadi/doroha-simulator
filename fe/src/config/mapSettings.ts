// Define the types for clarity and type safety (optional)
type Point = [number, number];  // A point represented as [latitude, longitude]
type Bounds = [Point, Point];   // Bounds represented as [southWest, northEast]

export const mapSettings = {
    defaultBounds: [
        [50.4190004, 30.4990038], // Southwest point
        [50.4489975, 30.5749117]  // Northeast point
    ] as Bounds,
    defaultCenter: [50.4339990, 30.5369577] as Point,
    earthCircumferenceInMeters: 40075016.686,
    mapDimensions: {
        width: 1920,
        height: 1080
    },
    zoomOffset: 17
};