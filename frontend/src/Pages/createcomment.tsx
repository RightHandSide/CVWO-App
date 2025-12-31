import React, { useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Alert, Box, Button, Paper, TextField, Typography } from "@mui/material";

function CreateComment() {
    const {id} = useParams<{ id: string }>();
    const [body, setBody] = useState("");
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    async function HandleSubmit(e: React.FormEvent) {
        e.preventDefault();
        setError(null);

        if (!body.trim()) {
            setError("Value Required");
            return;
        }

        try {
            const res = await fetch(`http://localhost:8000/createcomments/${id}`, {
                method: "POST",
                headers: {"Content-Type" : "application/json"},
                body: JSON.stringify({
                    title: "",
                    body: body
                }),
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
            navigate(`/post/${id}`)
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
                            Create Comment
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
                                    label="Body"
                                    value={body}
                                    onChange={e => setBody(e.target.value)}
                                    fullWidth/>
                                <Button type="submit" variant="contained" fullWidth>
                                    Create
                                </Button>

                                {error && <Alert severity="error">{error}</Alert>}
                            </Box>
                </Paper>
        </Box>
    );
}

export default CreateComment;