import NavBarComponent from "../components/NavBarComponent";
import React from "react";
import MetricSectionComponent from "../components/MetricSectionComponent";
import {API_METRIC_NAMES, SIMULATOR_METRIC_NAMES} from "../util/MetricClient";

function MetricPage () {
MetricSectionComponent(API_METRIC_NAMES, SIMULATOR_METRIC_NAMES)
    return (
    <div className="min-h-screen items-center justify-center bg-gray-100">
        <div className="mx-auto w-4/5 bg-white p-1 rounded shadow">
            <NavBarComponent/>
            {MetricSectionComponent(API_METRIC_NAMES, SIMULATOR_METRIC_NAMES)}
        </div>
    </div>
)
}

export default MetricPage