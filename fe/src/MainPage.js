import React, {useState, useEffect} from "react";

function MainPage() {
    const [data, setData] = useState(null);

    //make http GET call to localhost:8080/fe

    useEffect(() => {
        fetch("http://localhost:8080/fe")
            .then((res) => res.json())
            .then((data) => setData(data.message))
            .then((data) => console.log("Hello there" +  data));
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