import { Link } from "react-router-dom";
import * as PropTypes from 'prop-types';
import { Card, CardActionArea, CardContent, Typography } from "@mui/material";

type FormProps = {
    type: string;
    id: number;
    title: string;
    text: string;
};

function Topic(props: FormProps) {
    return (
        <Card
            sx={{
                maxWidth: 800,
                mx: "auto",
                mb: 2
            }}>
                <CardActionArea component={Link} to={`/${props.type}/${props.id}`}>
                    <CardContent>
                        <Typography variant="h6">
                            {props.title}
                        </Typography>
                        <Typography variant="body2" color="textSecondary">
                            {props.text}
                        </Typography>
                    </CardContent>
                </CardActionArea>
            </Card>
    );
}

Topic.propTypes = {
    type: PropTypes.string.isRequired,
    id: PropTypes.number,
    title: PropTypes.string,
    text: PropTypes.string
};

export default Topic;