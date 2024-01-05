import {useState, useRef, useEffect} from 'react';
import {LatLngTuple, LatLngBounds} from 'leaflet';
import {mapSettings} from "../config/mapSettings";

// Define types for clarity
type MapSetup = {
    center: LatLngTuple;
    zoom: number;
    leafletBounds: React.MutableRefObject<LatLngBounds | null>;
};

export const useMapSetup = (): MapSetup => {
    // Define state for center and zoom
    const [center, setCenter] = useState<LatLngTuple>([0, 0]); // Default values
    const [zoom, setZoom] = useState<number>(13); // Default zoom level

    // Define ref for leafletBounds
    const leafletBounds = useRef<LatLngBounds | null>(null);

    useEffect(() => {
        const [sw, ne] = mapSettings.defaultBounds;

        leafletBounds.current = new LatLngBounds(sw, ne);

        const centerLatitude = (ne[0] + sw[0]) / 2;
        const centerLongitude = (ne[1] + sw[1]) / 2;
        setCenter([centerLatitude, centerLongitude]);

        const earthCircumferenceInMeters = mapSettings.earthCircumferenceInMeters;
        const viewportHeight = window.innerHeight; // Height of the viewport in pixels

        // Assuming your CSS values
        const mapHeightVh = 95; // 95vh for height
        const mapWidthVh = 130; // 130vh for width

        // Convert vh to pixels
        const mapHeight = (mapHeightVh / 100) * viewportHeight;
        const mapWidth = (mapWidthVh / 100) * viewportHeight; //


        const longitudeDiff = Math.abs(ne[1] - sw[1]);
        const latitudeDiff = Math.abs(ne[0] - sw[0]);

        const meterPerPxLongitude = earthCircumferenceInMeters * (longitudeDiff / 360);
        const meterPerPxLatitude = earthCircumferenceInMeters * (latitudeDiff / 360) / Math.cos(((ne[0] + sw[0]) / 2) * (Math.PI / 180));

        let zoomLongitude = Math.log(mapWidth * (1 / meterPerPxLongitude)) / Math.log(2);
        let zoomLatitude = Math.log(mapHeight * (1 / meterPerPxLatitude)) / Math.log(2);
        // or any other value that fits your need
        setZoom(Math.min(zoomLongitude, zoomLatitude) + mapSettings.zoomOffset);


        // If these values depend on external data, make sure to handle loading states
    }, []); // Empty dependency array to run only once

    return {center, zoom, leafletBounds};
};