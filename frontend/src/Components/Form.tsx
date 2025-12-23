import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import * as PropTypes from "prop-types";
import { Alert, Box, Button, Paper, TextField, Typography } from "@mui/material";

type FormProps = {
    type: string;
    fetch: string;
    to: string;
}

function Form(props: FormProps) {
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
        <Box
            sx={{
                minHeight: "100vh",
                display: "flex",
                justifyContent: "center",
                alignItems: "center"
            }}>
                <Paper
                    sx={{
                        p: 4,
                        width: 320
                    }}>
                        <Typography variant="h4" align="center" gutterBottom>
                            {props.type}
                        </Typography>
                        <Box
                            component="form"
                            onSubmit={HandleSubmit}
                            sx={{
                                display: "flex",
                                flexDirection: "column",
                                gap: 2
                            }}>
                                <TextField
                                    label="Name"
                                    value={name}
                                    onChange={e => setName(e.target.value)}
                                    fullWidth
                                />
                                <Button type="submit" variant="contained" fullWidth>
                                    {props.type}
                                </Button>

                                {error && <Alert severity="error">{error}</Alert>}
                            </Box>
                    </Paper>
            </Box>
    );
}

Form.propTypes = {
    type: PropTypes.string.isRequired,
    fetch: PropTypes.string.isRequired,
    to: PropTypes.string.isRequired
}

export default Form;