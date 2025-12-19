import { useLocation } from "react-router-dom";

function RegisterComplete() {
    const location = useLocation();
    const state = location.state as { name?: string } | null;
    const name = state?.name ?? "";

    return (
        <div className="center-box">
            {name && <h1>{name} Registered Successfully!</h1>}
            <p>Please Login with your Username</p>
        </div>
    )
}

export default RegisterComplete;