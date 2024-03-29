import React, {useEffect, useMemo, useState} from "react";
import ActiveJourneyList from "../components/ActiveJourneySideBarComponent";
import {Journey} from "../types/Journey";
import MapComponent from "../components/MapComponent";
import {JourneyListContext} from "../util/JourneyListContext";
import {calculateBearing} from "../util/CarIconBearingCalculator";
import WebSocketClient from "../util/WebSocketClient";
import NavBarComponent from "../components/NavBarComponent";


function Homepage() {

    const [journeys, setJourneys] = useState<Map<number, Journey>>(new Map());

    useEffect(() => {
        const websocketEndpoint: string | undefined = process.env.REACT_APP_API_WEBSOCKET;
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
        <div className="min-h-screen items-center justify-center bg-gray-100">
            <div className="mx-auto w-4/5 bg-white p-1 rounded shadow">
                <NavBarComponent/>
                <div className= "flex">
                <JourneyListContext.Provider value={journeyContextMemo}>
                    <div className="p-1 w-5/6">
                        <MapComponent/>
                    </div>
                    <div className="p-1 w-1/6 overflow-y-auto"
                         style={{maxHeight: 'calc(100vh - 4rem)'}}>
                        <ActiveJourneyList/>
                    </div>
                </JourneyListContext.Provider>
                </div>
            </div>
        </div>
    )
}

export default Homepage