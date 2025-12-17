import { useLocation } from "react-router-dom";

function RegisterComplete() {
    const location = useLocation();
    const state = location.state as { name?: string } | null;
    const name = state?.name ?? "";

    return (
        <>
        {name && <h1>{name} Registered Successfully!</h1>}
        </>
    )
}

export default RegisterComplete;