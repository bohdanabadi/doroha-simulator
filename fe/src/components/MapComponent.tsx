import  {MapContainer, TileLayer, Marker} from "react-leaflet";
import 'leaflet-rotatedmarker'; // Import the plugin

import {useMapSetup} from "../util/MapSetup";
import {JourneyListContext} from "../util/JourneyListContext";
import {useContext, useEffect, useMemo, useRef} from "react";
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
    const lastPositionsRef = useRef(new Map());

    const carIcon = L.divIcon({
        className: 'car-icon',
        html: `<img src="${carIconUrl}" width="32" height="32">`,
        iconSize: [32, 32],
        iconAnchor: [16, 16],
        popupAnchor: [0, -19]
    });

    const markers = useMemo(() => Array.from(journeys.entries()).map(([id, journey]) => ({
        id,
        position: new LatLng(journey.currentPoint.Y, journey.currentPoint.X),
        bearing: journey.bearing

    })), [journeys]);

    useEffect(() => {
        markers.forEach(({id, position, bearing}) => {
            const marker = markerRefs.current.get(id);
            const lastPosition = lastPositionsRef.current.get(id);
            const journey = journeys.get(id);
            if (marker) {
                if(journey !== undefined && journey.status === "FINISHED"){
                    // Remove journey from state
                    removeJourney(id);
                } else {
                    const hasPositionChanged = !lastPosition ||
                        lastPosition.lat !== position.lat ||
                        lastPosition.lng !== position.lng;
                    if (hasPositionChanged) {
                        // Animate and update marker
                        animateMarker(marker, lastPosition || position, position, 400);
                        marker.setRotationAngle(bearing);
                        // Update the last known position
                        lastPositionsRef.current.set(id, position);
                    }
                }
            }
        });
    }, [markers, journeys, removeJourney]);

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
                {markers.map(({id, position, bearing}) =>  {
                    return <Marker key={id}
                                   position={position}
                                   icon={carIcon}
                                   rotationAngle={bearing}
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
