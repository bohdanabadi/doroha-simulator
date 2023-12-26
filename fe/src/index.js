import React from 'react';
import ReactDOM from 'react-dom';
import reportWebVitals from './reportWebVitals';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';
import Homepage from "./pages/Homepage.tsx";

// Correctly set the paths for Leaflet marker images
delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
    iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
    iconUrl: require('leaflet/dist/images/marker-icon.png'),
    shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
});


ReactDOM.render(
    <React.StrictMode>
        <Homepage/>
    </React.StrictMode>,
    document.getElementById('root')
);

reportWebVitals();
