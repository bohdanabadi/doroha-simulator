import MetricComponent from "./MetricComponent";

function MetricSectionComponent(apiMetrics: string[], simulatorMetrics: string[]) {
    return (
        <div>
            <div className="container mx-auto p-4">
                <h2 className="text-xl font-bold uppercase mb-4">API Metrics</h2>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {apiMetrics.map(item => (
                        MetricComponent(item)
                    ))}
                </div>
            </div>
            <div className="container mx-auto p-4 mt-8">
                <h2 className="text-xl font-bold uppercase mb-4">Simulator Metrics</h2>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {simulatorMetrics.map(item => (
                        MetricComponent(item)
                    ))}
                </div>
            </div>
        </div>
    )
        ;
    // <div className="flex">
    // {/* Section 1 */}
    // <section className="w-1/2 p-4">
    // <h1 className="text-xl font-bold mb-4">API Metrics</h1>
    // <div className="flex flex-col">
    // {/* Render API metric cards */}
    // {apiMetrics.map(item => (
    //                 MetricComponent(item)
    //             ))}
    //         </div>
    //     </section>
    //     <section className="w-1/2 p-4">
    //         <h1 className="text-xl font-bold mb-4">Simulator Metrics</h1>
    //         <div className="flex flex-col">
    //             {/* Render API metric cards */}
    //             {simulatorMetrics.map(item => (
    //                 MetricComponent(item)
    //             ))}
    //         </div>
    //     </section>
    // </div>
}

export default MetricSectionComponent