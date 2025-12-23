import { useNavigate } from "react-router-dom";
import { Box, Button, Stack, Typography } from "@mui/material";

function RegisterComplete() {
    const navigate = useNavigate();

    return (
        <Box
            sx={{
                minHeight: "100vh",
                display: "flex",
                justifyContent: "center",
                alignItems: "center"
            }}>
                <Box textAlign="center">
                    <Stack direction="column">
                        <Typography variant="h4">
                            Registered Successfully!
                        </Typography>
                        <Typography sx={{mb: 4}}>
                            Please Login with your Username
                        </Typography>
                        <Button variant="contained" onClick={() => navigate("/login")}>
                            Login
                        </Button>
                    </Stack>
                </Box>
            </Box>
    );
}

export default RegisterComplete;