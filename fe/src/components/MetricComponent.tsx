import {SetStateAction, useState} from "react";
import {METRIC_PRETTY_NAMES, useFetchMetrics} from "../util/MetricClient";
import {convertTimestampToLocalDateTime, HealthColorMapping, HealthKey} from "../util/MetricHelper";
function MetricComponent(metricName: string) {

    const [duration, setDuration] = useState('1h'); // Default duration
    const [data, isLoading, error] = useFetchMetrics(metricName, duration);

    const handleDurationChange = (newDuration: SetStateAction<string>) => {
        setDuration(newDuration);
    };

    let metricHealth: HealthKey = 'u'
    let { bgColor, textColor, circleColor } = HealthColorMapping[metricHealth];
    let metricValue: string
    let lastMetricUpdate : string
    if (isLoading) {
        metricValue = 'Loading'
        lastMetricUpdate  = 'Loading'
    } else {
        metricHealth = data?.metricHealth ?  (data.metricHealth as HealthKey) : 'u'
        const colors = HealthColorMapping[metricHealth];
        bgColor = colors.bgColor;
        textColor = colors.textColor;
        circleColor = colors.circleColor;
        metricValue = data?.metricValue ? data.metricValue : 'N/A';
        lastMetricUpdate = data?.metricTime ? convertTimestampToLocalDateTime(data.metricTime) : 'N/A';
    }

    // Get pretty name for the metric
    let MetricPrettyName = METRIC_PRETTY_NAMES[metricName as keyof typeof METRIC_PRETTY_NAMES] || metricName;
    return (
        <div className={`p-4 ${bgColor} rounded-lg shadow-md`}>
            <div className="flex items-center justify-between">
                <h3 className={`font-semibold ${textColor}`}>{MetricPrettyName}</h3>
                <div className="flex items-center">
                    <span className={`text-lg font-bold ${textColor}`}>{metricValue}</span>
                    <div className={`w-4 h-4 ${circleColor} rounded-full ml-2`}></div>
                </div>
            </div>
            <div className="flex justify-between items-center mt-2">
                <div className="flex">
                    {/* Toggle buttons */}
                    {['1h', '6h', '24h', '7d'].map((timeFrame) => (
                        <button
                            key={timeFrame}
                            onClick={() => handleDurationChange(timeFrame)}
                            className={`px-3 py-1 mr-2 text-sm rounded ${
                                duration === timeFrame ? 'bg-blue-500 text-white' : 'bg-gray-200'
                            }`}
                        >
                            {timeFrame}
                        </button>
                    ))}
                </div>
                <div className="text-right">
                    <span className="text-sm text-gray-500">Last update: {lastMetricUpdate}</span>
                </div>
            </div>
        </div>
    );
}

export default MetricComponent