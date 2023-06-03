# Traffic-Simulation
This is a personal project that would simulate traffic on a map.

This project was inspired Juraj Majerik Uber Clone project, I want to build something similar that exposes me to the full SDLC of a distributed system.
This will have a similar approach of simulating rides, but focusing on traffic and picking up the best router a trip make to get from point A to point B.

There will    
## Concept 
![image](https://github.com/bohdanabadi/Traffic-Simulation/assets/24784589/bce73f75-b85f-46c1-8385-537808c60d3d)

The basic idea that we will simulate someone driving from point A to point B and we need to navigate them using the fastest route that will take few things into our consideration. This wil take into consideration traffic levels and routes will keep updating accordingly.

The main components will be be the Simulation engine, APIs, Database, and frontend.


<img width="960" alt="img-6" src="https://github.com/bohdanabadi/Traffic-Simulation/assets/24784589/dbdbeb94-c7ca-4827-b533-2256865822a1">


## Simulation Engine
This component will basically be creating rides from point A to point B and interfacing API component. Also we could introduce functionality based on the time, like simulated peak hours.

## API
This component be reasponsible for the logic of the app picking up routes, calculating ETA, and more.

## Database 
This will store user related information, and the city of the map in order to have up to date view on the route state.

## FrontEnd 
We will have a frontend web app, that display a simple map and cars driving to their destination and on the side list of trips and each trip ETA. We can also perhaps display a number by how much this trip being delayed versus if not obstacles occuring.
In addition I want to include monitoring metrics that will give some info about the state of the app.

![image](https://github.com/bohdanabadi/Traffic-Simulation/assets/24784589/e25c6c93-a977-4bd8-9163-ad02c7ff1804)

The outcome of this project is to go through all the hoops of developing software and mimic as much as possible production like practices in order to learn as much as possible.
