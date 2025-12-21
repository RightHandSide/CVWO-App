import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import * as PropTypes from 'prop-types';

function Form(props) {
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
            const res = await fetch(props.fetch, {
                method: "POST",
                headers: {"Content-Type" : "application/json"},
                body: JSON.stringify({ name }),
                credentials: "include",
            });

            if (!res.ok) {
                setError(`Server Returned Status ${res.status}`);
                return;
            }

            const data = await res.json();
            const code = data.errorCode;
            if (code != 0) {
                setError(data.messages[0]);
                return;
            }
            navigate(props.to);
        } catch (err) {
            setError("Network or Server Error");
        }
    }

    return (
        <div className="center-box">
            <h1>{props.type}</h1>
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
                    {props.type}
                </button>
            </form>
            {error && <p>{error}</p>}
        </div>
    )
}
Form.propTypes = {
    type: PropTypes.string.isRequired,
    fetch: PropTypes.string.isRequired,
    to: PropTypes.string.isRequired,
};

export default Form;