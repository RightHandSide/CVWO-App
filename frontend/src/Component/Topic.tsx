import * as PropTypes from 'prop-types';
import { Link } from "react-router-dom";

function Topic(props) {
    return (
        <Link
            key={props.id}
            to={"/" + props.type + "/" + props.id}
            style={{textDecoration: "none"}} >
            <div key={props.id} className="thread" >
                <h2 className="thread-title">{props.title}</h2>
                <p className="thread-text">{props.text}</p>
            </div>
        </Link>
    );
}
Topic.propTypes = {
    type: PropTypes.string.isRequired,
    id: PropTypes.number,
    title: PropTypes.string,
    text: PropTypes.string
};

export default Topic;