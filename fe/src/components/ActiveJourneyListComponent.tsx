import JourneyCard from "./JourneyCardComponent";
import {useContext} from "react";
import {JourneyListContext} from "../util/JourneyListContext";
import "../assets/css/HomePage.css"
import "../assets/css/SideBar.css"

function ActiveJourneyList() {
    const{journeys} = useContext(JourneyListContext);
    return (
        <div className="sidebar">
            { journeys.size > 0 ? (
                Array.from(journeys.values()).map(journey => (
                <JourneyCard key={journey.id} journey={journey} />
                ))
            ) : (
                <div className="empty-nav"> No Journeys In Progress.</div>
            )}
        </div>
    );
}

export default ActiveJourneyList;