import {MetricResponse} from "../types/Metric";
import {useEffect, useState} from "react";

export const API_METRIC_NAMES: string[] = ["avg_response_latency_seconds", "avg_database_latency_seconds", "total_response_error_counter"]

export const SIMULATOR_METRIC_NAMES: string[] = ["simulations_per_minute", "avg_simulation_duration", "total_successful_simulations", "total_failed_simulations"]

enum Metrics {
    AVG_RESPONSE_LATENCY = "avg_response_latency_seconds",
    AVG_DB_RESPONSE_LATENCY = "avg_database_latency_seconds",
    TOTAL_RESPONSE_ERROR = "total_response_error_counter",
    SIMULATION_PER_MINUTE = "simulations_per_minute",
    AVG_SIMULATION_DURATION = "avg_simulation_duration",
    // AVG_SIMULATION_DISTANCE = "avg_simulation_distance",
    TOTAL_SUCCESSFUL_SIMULATIONS = "total_successful_simulations",
    TOTAL_FAILED_SIMULATIONS = "total_failed_simulations"
}

export const METRIC_PRETTY_NAMES: { [key in Metrics]: string } = {
    [Metrics.AVG_RESPONSE_LATENCY]: "Average Response Latency",
    [Metrics.AVG_DB_RESPONSE_LATENCY]: "Average Database Query Latency",
    [Metrics.TOTAL_RESPONSE_ERROR]: "Total Error Response",
    [Metrics.SIMULATION_PER_MINUTE]: "Simulations Per Minute",
    [Metrics.AVG_SIMULATION_DURATION]: "Average Simulation Duration",
    // [Metrics.AVG_SIMULATION_DISTANCE]: "Average Simulation Distance",
    [Metrics.TOTAL_SUCCESSFUL_SIMULATIONS]: "Total Completed Simulations",
    [Metrics.TOTAL_FAILED_SIMULATIONS]: "Total Failed Simulations"
}


export function useFetchMetrics(metricType: string, duration: string): [MetricResponse | null, boolean, any] {
    const [data, setData] = useState<MetricResponse | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<any>(null);

    useEffect(() => {
        const fetchData = async () => {
            setIsLoading(true);
            try {
                const response = await fetch(process.env.REACT_APP_API_METRICS + `?metricType=${metricType}&duration=${duration}`);
                if(!response.ok){
                    throw new Error('Net response was not ok')
                }
                const result = await response.json();
                setData(result);
                setIsLoading(false);
            } catch (error) {
                setError(error);
                setIsLoading(false);
            }
        };
        fetchData(); // Initial fetch
        // Set up polling
        const intervalId = setInterval(fetchData, 45000);
        // Clean up on unmount or when dependencies change
        return () => clearInterval(intervalId);
    }, [metricType, duration]); // Dependencies

    return [data, isLoading, error];
}
