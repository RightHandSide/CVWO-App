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

            navigate("/register/completed", {state: { name }});
        } catch (err) {
            setError("Network or Server Error");
        }
    }

    return (
        <>
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
        </>
    )
}

export default Register;