import { useParams } from "react-router-dom";
import CreateForm from "../Components/CreateForm";

function CreatePost() {
    const {id} = useParams<{ id: string }>();

    return (
        <CreateForm
            fetch={`http://localhost:8000/createposts/${id}`}
            nav={`/thread/${id}`}
            header="Create Post"
            title="Title"
            body="Body"/>
    );
}

export default CreatePost;