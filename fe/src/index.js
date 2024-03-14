import React from 'react';
import ReactDOM from 'react-dom';
import reportWebVitals from './reportWebVitals';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Homepage from "./pages/Homepage.tsx";
import MetricPage from "./pages/MetricPage";
import "./assets/css/index.css"
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import ErrorPage from "./pages/ErrorPage";


// Correctly set the paths for Leaflet marker images
delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
    iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
    iconUrl: require('leaflet/dist/images/marker-icon.png'),
    shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
});

const router = createBrowserRouter([
    {
        path: "/",
        element:<Homepage/>,
        errorElement: <ErrorPage />,
    },
    {
        path: "/metrics",
        element: <MetricPage/>,
        errorElement: <ErrorPage/>
    }
]);

ReactDOM.createRoot(document.getElementById("root")).render(
    <React.StrictMode>
            <RouterProvider router={router} />
    </React.StrictMode>
);

reportWebVitals();
