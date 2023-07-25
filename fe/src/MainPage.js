import React, {useState, useEffect} from "react";

function MainPage() {
    const [data, setData] = useState(null);

    //make http GET call to localhost:8080/fe

    useEffect(() => {
        fetch(process.env.REACT_APP_FE_URL+"/fe")
            .then((res) => res.json())
            .then((data) => setData(data.message))
    },[]);

    if (data === null) {
        return <div>Loading...</div>;
    } else {
        return (
            <div>
                <h1>Data from the Backend & Testing Github Actions:</h1>
                <pre>{data}</pre>
            </div>
        );
    }
}
export default MainPage;