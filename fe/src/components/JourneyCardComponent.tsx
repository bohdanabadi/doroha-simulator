import {Journey} from "../types/Journey";
import Typography from '@mui/material/Typography';
import CardContent from '@mui/material/CardContent';
import Card from '@mui/material/Card';
import {Box, LinearProgress, LinearProgressProps} from "@mui/material";
import {grey} from "@mui/material/colors";


interface JourneyCardProps {
    journey: Journey;
}

function LinearProgressWithLabel(props: LinearProgressProps & { value: number }) {
    const lightGrey = '#d3d3d3'; // Lighter shade
    const darkGrey = '#9e9e9e'; // Darker shade
    return (
        <Box sx={{ display: 'flex', alignItems: 'center' }}>
            <Box sx={{ width: '100%', mr: 1 }}>
                <LinearProgress variant="determinate" {...props}
                    sx={{
                        backgroundColor: lightGrey, // Style for the background of the bar
                        '& .MuiLinearProgress-bar': {
                            backgroundColor: darkGrey, // Style for the determinate bar
                        }
                    }}
                 />
            </Box>
            <Box sx={{ minWidth: 35 }}>
                <Typography variant="body2" color={grey[500]}>{`${Math.round(props.value,)}%`}
                </Typography>
            </Box>
        </Box>
    );
}
function JourneyCard( {journey}: Readonly<JourneyCardProps>) {
    const readableDistance = `${Math.round(journey.distance)} meters`;
    return (
    <Card sx={{ border: `2px solid grey`, backgroundColor: grey[200], margin: `3px` }}>
        <CardContent>
            <Typography gutterBottom variant="h5" component="div" align={"center"} color={grey[500]}>
                {journey.id}
            </Typography>
            <Typography sx={{ mb: 1.5, fontSize: 14 , color:grey[500]}}>
                {readableDistance}
            </Typography>
            <LinearProgressWithLabel value={journey.progress * 100} />
        </CardContent>
    </Card>
);

}

export default JourneyCard;