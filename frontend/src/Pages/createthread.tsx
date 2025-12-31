import CreateForm from "../Components/CreateForm";

function CreateThread() {
    return (
        <CreateForm
            fetch="http://localhost:8000/createthreads"
            nav="/threads"
            header="Create Thread"
            title="Title"
            body="Description"/>
    );
}

export default CreateThread;