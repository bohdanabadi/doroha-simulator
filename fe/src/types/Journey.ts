export interface Journey {
    id: number; // Assuming int32 maps to TypeScript's number
    startingPoint: Point; // float64 in Go maps to number in TypeScript
    endingPoint: Point;
    currentPoint: Point;
    prevPoint: Point;
    bearing: number;
    progress: number
    distance: number;
    dateCreate: string;
    status: string;
    attempts: number;
}
interface Point {
    X : number
    Y: number
}