import JourneyCard from "./JourneyCardComponent";
import {useContext} from "react";
import {JourneyListContext} from "../util/JourneyListContext";

function ActiveJourneyList() {
    const{journeys} = useContext(JourneyListContext);
    return (
        <div >
            { journeys.size > 0 ? (
                Array.from(journeys.values()).map(journey => (
                <JourneyCard key={journey.id} journey={journey} />
                ))
            ) : (
                <div> No Journeys In Progress.</div>
            )}
        </div>
    );
}

export default ActiveJourneyList;