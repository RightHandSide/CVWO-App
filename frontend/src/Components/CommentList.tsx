import * as PropTypes from 'prop-types';
import { Card, CardContent, Typography } from "@mui/material";

type FormProps = {
    type: string;
    id: number;
    text: string;
};

function Ending(props: FormProps) {
    return (
        <Card
            sx={{
                maxWidth: 800,
                mx: "auto",
                mb: 2
            }}>
                <CardContent>
                    <Typography variant="body2" color="textSecondary">
                        {props.text}
                    </Typography>
                </CardContent>
            </Card>
    );
}

Ending.propTypes = {
    type: PropTypes.string.isRequired,
    id: PropTypes.number,
    title: PropTypes.string,
    text: PropTypes.string
};

export default Ending;