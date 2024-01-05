import {MapContainer, TileLayer, Marker} from "react-leaflet";
import 'leaflet-rotatedmarker'; // Import the plugin

import {useMapSetup} from "../util/MapSetup";
import {JourneyListContext} from "../util/JourneyListContext";
import {useContext, useEffect, useRef} from "react";
import carIconUrl from "../assets/car.png"
import L, {LatLng} from "leaflet";
import "../assets/css/MapComponent.css"


function animateMarker(marker : L.Marker, startPosition : LatLng, endPosition : LatLng, duration : number) {
    let startTime: number;
    function animate(time: number) {
        if (!startTime) {
            startTime = time;
        }

        const timeElapsed = time - startTime;
        let progress = timeElapsed / duration;

        if (progress > 1) progress = 1; // Ensure progress doesn't exceed 1
        const lat = startPosition.lat + (endPosition.lat - startPosition.lat) * progress;
        const lng = startPosition.lng + (endPosition.lng - startPosition.lng) * progress;
        marker.setLatLng([lat, lng]);

        if (progress < 1) {
            requestAnimationFrame(animate);
        }
    }

    requestAnimationFrame(animate);
}

function MapComponent() {
    const {center, zoom, leafletBounds} = useMapSetup();
    const {journeys, removeJourney} = useContext(JourneyListContext);
    const markerRefs = useRef(new Map());
    useEffect(() => {
        journeys.forEach((journey, id) => {
            if(journey.status === "FINISHED"){
                removeJourney(id);
            }
        });
    })
    const carIcon = L.icon({
        iconUrl: carIconUrl, // URL to the car icon image
        iconSize: [32, 32], // Size of the icon
        iconAnchor: [16, 16], // Point of the icon which will correspond to marker's location
        popupAnchor: [0, -19] // Point from which the popup should open relative to the iconAnchor
    });

    useEffect(() => {
        journeys.forEach((journey, id) => {
            const marker = markerRefs.current.get(id);
            if (marker) {

                const startPosition = new LatLng(journey.prevPoint.Y, journey.prevPoint.X);
                const endPosition = marker.getLatLng();
                animateMarker(marker, startPosition, endPosition, 350);
            }
        });
    }, [journeys]);


    return (
        <div className={"leaflet-map-div"}>
            <MapContainer className={"leaflet-map"}
                key={zoom}
                center={center}
                zoom={zoom}
                minZoom={zoom}
                scrollWheelZoom={false} // disable zoom on scroll
                dragging={false} // disable dragging
                zoomControl={false}
                maxBounds={leafletBounds.current ? leafletBounds.current : undefined}
                maxBoundsViscosity={50.0} // Determines the elasticity of the bounds.
            >
                <TileLayer
                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                    attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
                />
                {Array.from(journeys.entries()).map(([id, journey]) => {
                    return <Marker key={`${id}-${journey.bearing}`}
                                   position={[journey.currentPoint.Y, journey.currentPoint.X]}
                                   icon={carIcon}
                                   rotationAngle={journey.bearing}
                                   rotationOrigin="center"
                                   ref={(ref) => {
                                       if (ref) {
                                           markerRefs.current.set(id, ref);
                                       }
                                   }}

                    />;
                })}
            </MapContainer>

        </div>
    );
}

export default MapComponent;
