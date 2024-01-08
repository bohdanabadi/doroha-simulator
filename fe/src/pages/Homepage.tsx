import {useEffect, useMemo, useState} from "react";
import ActiveJourneyList from "../components/ActiveJourneyListComponent";
import {Journey} from "../types/Journey";
import MapComponent from "../components/MapComponent";
import "../assets/css/HomePage.css"
import {JourneyListContext} from "../util/JourneyListContext";
import {calculateBearing} from "../util/CarIconBearingCalculator";
import WebSocketClient from "../util/WebSocketClient";


function Homepage () {

    const [journeys, setJourneys] = useState<Map<number, Journey>>(new Map());

    useEffect(() => {
        const websocketEndpoint:string | undefined   = process.env.REACT_APP_API_WEBSOCKET;
        if (typeof websocketEndpoint === 'string' && websocketEndpoint.trim() !== '') {
            // Valid endpoint, proceed with the connection
            const webSocketClient = new WebSocketClient(websocketEndpoint, 5);
            const handleMessage = (message: string) => {
                const journeyData: Journey = JSON.parse(message);
                setJourneys((prevCarPosition) => {
                    const updatedCarPositions = new Map(prevCarPosition)
                    const bearing = calculateBearing(journeyData.prevPoint.Y, journeyData.prevPoint.X, journeyData.currentPoint.Y, journeyData.currentPoint.X)
                    journeyData.bearing = (bearing - 90 + 360) % 360;
                    updatedCarPositions.set(journeyData.id, journeyData)
                    return updatedCarPositions;
                })
            };

            webSocketClient.addListener(handleMessage);
            webSocketClient.connect();

            return () => {
                webSocketClient.removeListener(handleMessage);
                webSocketClient.close();
            }
        } else {
            console.log(websocketEndpoint);
            // Invalid or missing endpoint, handle the error
            console.error("WebSocket endpoint is undefined, null, or empty. Cannot establish WebSocket connection.");
        }
    }, []);

    const removeJourney = (id: number) => {
        setJourneys(prevJourneys => {
            const updatedJourneys = new Map(prevJourneys);
            updatedJourneys.delete(id);
            return updatedJourneys;
        });
    };

    // useMemo to optimize performance
    const journeyContextMemo = useMemo(() => ({
        journeys,
        setJourneys,
        removeJourney
    }), [journeys, setJourneys]);

    return (
        <div className="home-container">
            <JourneyListContext.Provider value={journeyContextMemo}>
            <div className="map-container">
                    <MapComponent/>
                </div>
                <div className="journey-list">
                    <ActiveJourneyList/>
                </div>
            </JourneyListContext.Provider>
        </div>
    )
}

export default Homepage