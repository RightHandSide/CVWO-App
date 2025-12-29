import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Alert, Box, Button, Paper, TextField, Typography } from "@mui/material";

function CreateForm() {
    const [title, setTitle] = useState("");
    const [desc, setDesc] = useState("");
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    async function HandleSubmit(e: React.FormEvent) {
        e.preventDefault();
        setError(null);

        if (!title.trim() || !desc.trim()) {
            setError("Value Required");
            return;
        }

        try {
            const res = await fetch("http://localhost:8000/createthreads", {
                method: "POST",
                headers: {"Content-Type" : "application/json"},
                body: JSON.stringify({
                    title: title,
                    body: desc
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
            navigate("http://localhost:5173/threads")
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
                            Create Thread
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
                                    label="Title"
                                    value={title}
                                    onChange={e => setTitle(e.target.value)}
                                    fullWidth/>
                                <TextField
                                    label="Description"
                                    value={desc}
                                    onChange={e => setDesc(e.target.value)}
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

export default CreateForm;