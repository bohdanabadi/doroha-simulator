import {createContext, Dispatch, SetStateAction} from 'react';
import {Journey} from "../types/Journey";


interface JourneyListContextType {
    journeys: Map<number, Journey>;
    setJourneys: Dispatch<SetStateAction<Map<number, Journey>>>;
    removeJourney: (id: number) => void;
}

export const JourneyListContext = createContext<JourneyListContextType>({
    journeys: new Map(),
    setJourneys: () => {},
    removeJourney: () => {}
});