import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function Register() {
    const [name, setName] = useState("");
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    async function HandleSubmit(e: React.FormEvent) {
        e.preventDefault();
        setError(null);

        if (!name.trim()) {
            setError("Name Required");
            return;
        }

        try {
            const res = await fetch("http://localhost:8000/register", {
                method: "POST",
                headers: {"Content-Type" : "application/json"},
                body: JSON.stringify({ name })
            });

            if (!res.ok) {
                setError(`Server Returned Status ${res.status}`);
                return;
            }

            const data = await res.json();
            const code = data.errorCode;
            if (code == 1) {
                setError("Repeated Name. Please Input a Different Name");
                return;
            }
            navigate("/register/completed", {state: { name }});
        } catch (err) {
            setError("Network or Server Error");
        }
    }

    return (
        <div className="center-box">
            <h1>Register</h1>
            <form onSubmit={HandleSubmit}>
                <label>
                    Name: {""}
                    <input
                        type="text"
                        value={name}
                        onChange={e => setName(e.target.value)}/>
                </label>
                <hr/>
                <button type="submit">
                    Submit
                </button>
            </form>
            {error && <p>{error}</p>}
        </div>
    )
}

export default Register;